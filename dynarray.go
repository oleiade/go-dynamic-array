package dynarray

import "errors"

type DynamicArray struct {
	logical_size int
	capacity     int
	container    []interface{}
}

// set_size method resizes the Dynamic array to the `length`
// size and copies elements in place
func (s *DynamicArray) set_size(length int) {
	new_container := make([]interface{}, length)
	copy(new_container, s.container)

	s.container = new_container
	s.capacity = length
}

// expand method resizes the dynamic array to twice
// it's current length. Fills the blanks with nil values
func (s *DynamicArray) expand() {
	if s.capacity == 0 {
		s.set_size(1)
	} else {
		s.set_size(s.capacity * 2)
	}
}

// shrink method resizes the dynamic array to half
// it's size. Leaves the blanks filled with nil values
func (s *DynamicArray) shrink() {
	if s.capacity >= 2 {
		s.set_size(s.capacity / 2)
	} else if s.capacity == 1 {
		s.set_size(1)
	}
}

// resize automatically expands the dynamic array if an
// insertion is required and it is already full. Or shrinks
// it whenever it is only a quarter full
func (s *DynamicArray) resize(new_logical_size int) {
	if new_logical_size >= s.capacity {
		s.expand()
	} else if new_logical_size > 0 && new_logical_size == s.capacity/4 {
		s.shrink()
	}
}

// Get retrieves the value, stored at the given index
// in dynamic array. Throws an error if index out of range.
func (s *DynamicArray) Get(index int) (interface{}, error) {
	if index >= 0 && index < s.logical_size {
		return s.container[index], nil
	}
	return nil, errors.New("Index out of range")
}

// Range retrieves a dynamic store slice with the provided
// edges. Throws an error if index out of range.
func (s *DynamicArray) Range(start int, stop int) ([]interface{}, error) {
	if start >= 0 && stop < s.logical_size {
		return s.container[start:stop], nil
	}
	return nil, errors.New("Index out of range")
}

// prepend_val adds a value at the beginning of the dynamic
// array, resizing it when needed
func (s *DynamicArray) PrependVal(elem interface{}) {
	new_logical_size := s.logical_size + 1
	s.resize(new_logical_size)

	container_copy := make([]interface{}, s.logical_size)
	copy(container_copy, s.container)
	s.container[0] = elem

	for i := 1; i < new_logical_size; i++ {
		s.container[i] = container_copy[i-1]
	}

	s.logical_size = new_logical_size
}

// prepend_val adds a slice of interface{} values
// at the beginning of the dynamic array,
// Resizes the array when needed.
func (s *DynamicArray) PrependVals(elems []interface{}) {
	new_logical_size := s.logical_size + len(elems)
	s.resize(new_logical_size)

	container_copy := make([]interface{}, s.logical_size)
	copy(container_copy, s.container)
	copy(s.container, elems)

	index := 0
	for i := len(elems); i < new_logical_size; i++ {
		s.container[i] = container_copy[index]
		index ++
	}

	s.logical_size = new_logical_size
}

// append_val adds a value at the end of the logical
// content of the array (non-nil values). Resizes
// the array when needed.
func (s *DynamicArray) AppendVal(elem interface{}) {
	new_logical_size := s.logical_size + 1
	s.resize(new_logical_size)

	s.container[s.logical_size] = elem
	s.logical_size += 1
}

// append_vals adds a slice of interface{} values
// at the end of the logical content of the array (non-nil values).
// Resizes the array when needed.
func (s *DynamicArray) AppendVals(elems []interface{}) {
	new_logical_size := s.logical_size + len(elems)
	s.resize(new_logical_size)

	index := 0
	for i := s.logical_size; i < new_logical_size; i++ {
		s.container[i] = elems[index]
		index ++
	}

	s.logical_size = new_logical_size
}

// insert_val inserts a value in the array at the index pos
// (0 indexed). Resizes the array when needed.
func (s *DynamicArray) InsertVal(elem interface{}, index int) {
	new_logical_size := s.logical_size + 1
	s.resize(new_logical_size)

	container_copy := make([]interface{}, s.logical_size)
	copy(container_copy, s.container)

	elem_inserted := false
	for i := 0; i < new_logical_size; i++ {
		if i == index {
			s.container[i] = elem
			elem_inserted = true
		} else {
			if elem_inserted == false {
				s.container[i] = container_copy[i]
			} else {
				s.container[i] = container_copy[i-1]
			}
		}
	}

	s.logical_size = new_logical_size

}

// remove_left pops the dynamic array 0 indexed value and realign
// values. Resizes the array when needed.
func (s *DynamicArray) RemoveLeft() {
	new_logical_size := s.logical_size - 1
	s.resize(new_logical_size)

	new_container := make([]interface{}, s.capacity)
	copy(new_container, s.container[1:s.logical_size])

	s.container = new_container
	s.logical_size = new_logical_size
}

// remove_rights pops the dynamic array last logical (non nil)
// indexed value and realign values. Resizes the array when needed.
func (s *DynamicArray) RemoveRight() {
	new_logical_size := s.logical_size - 1
	s.resize(new_logical_size)

	new_container := make([]interface{}, s.capacity)
	copy(new_container, s.container[0:new_logical_size])

	s.container = new_container
	s.logical_size = new_logical_size
}
