package dynarray

import (
	"testing"
)

func Test_Get_valid_index(t *testing.T) {
	arr := DynamicArray{}
	arr.PrependVal("1")
	val, err := arr.Get(0)

	if err != nil || val != "1" {
		t.Error("Get valid index test failed")
	} else {
		t.Log("Get valid index test passed")
	}
}

func Test_Get_non_existing_index(t *testing.T) {
	arr := DynamicArray{}
	_, err := arr.Get(0)

	if err == nil {
		t.Error("Get invalid index test failed")
	} else {
		t.Log("Get invalid index test passed")
	}
}
