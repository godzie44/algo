package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertLinkedList(t *testing.T) {
	testCases := []struct {
		inputList    []int
		expectedList []int
	}{
		{inputList: []int{1}, expectedList: []int{1}},
		{inputList: []int{1, 2}, expectedList: []int{1, 2}},
	}

	for _, tk := range testCases {
		list := NewLinkedList()
		fillList(list, tk.inputList)
		assert.Equal(t, tk.expectedList, list.IntoArray())
	}
}

func TestSearchLinkedList(t *testing.T) {
	testCases := []struct {
		inputList []int
		searchKey int
		keyExists bool
	}{
		{inputList: []int{1}, searchKey: 1, keyExists: true},
		{inputList: []int{1}, searchKey: 0, keyExists: false},
		{inputList: []int{1, 2, 3}, searchKey: 2, keyExists: true},
	}

	for _, tk := range testCases {
		list := NewLinkedList()
		fillList(list, tk.inputList)

		node, err := list.Search(tk.searchKey)
		if err != nil {
			assert.False(t, tk.keyExists)
		} else {
			assert.True(t, tk.keyExists)
			assert.Equal(t, tk.searchKey, node.key)
		}
	}
}

func TestDeleteLinkedList(t *testing.T) {
	testCases := []struct {
		inputList    []int
		deleteKey    int
		expectedList []int
	}{
		{inputList: []int{1}, deleteKey: 1, expectedList: []int(nil)},
		{inputList: []int{1, 2, 3}, deleteKey: 1, expectedList: []int{2, 3}},
		{inputList: []int{1, 2, 3}, deleteKey: 2, expectedList: []int{1, 3}},
	}

	for _, tk := range testCases {
		list := NewLinkedList()
		fillList(list, tk.inputList)
		node, _ := list.Search(tk.deleteKey)
		list.Delete(node)
		assert.Equal(t, tk.expectedList, list.IntoArray())
	}
}

func TestRevertLinkedList(t *testing.T) {
	testCases := []struct {
		inputList    []int
		expectedList []int
	}{
		{inputList: []int{1, 2}, expectedList: []int{2, 1}},
		{inputList: []int{1}, expectedList: []int{1}},
		{inputList: []int{1, 2, 3, 4}, expectedList: []int{4, 3, 2, 1}},
	}

	for _, tk := range testCases {
		list := NewLinkedList()
		fillList(list, tk.inputList)

		list.Revert()
		assert.Equal(t, tk.expectedList, list.IntoArray())
	}
}

func fillList(list *LinkedList, values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		list.Insert(&llNode{key: values[i]})
	}
}
