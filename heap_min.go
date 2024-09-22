package main

type Item struct {
	Name  string
	Value int32
}

type HeapMin struct {
	tree []*Item
}

func (h *HeapMin) getParentIndex(nodeIndex int) int {
	return int(nodeIndex / 2)
}

func (h *HeapMin) Delete(nodeIndex int) {
	h.swap(nodeIndex, len(h.tree)-1)
	h.tree = h.tree[0 : len(h.tree)-1]
	h.moveToBottom(nodeIndex)
}

func (h *HeapMin) Add(name string, value int32) {
	h.tree = append(h.tree, &Item{Name: name, Value: value})
	childIndex := len(h.tree) - 1
	if len(h.tree) == 1 {
		return
	}

	for {
		parentIndex := h.getParentIndex(childIndex)
		if parentIndex == childIndex {
			return
		}

		if h.tree[childIndex].Value >= h.tree[parentIndex].Value {
			return
		}
		h.swap(parentIndex, childIndex)
		childIndex = parentIndex
	}
}

func (h *HeapMin) swap(i, j int) {
	if i == j {
		return
	}

	h.tree[i], h.tree[j] = h.tree[j], h.tree[i]
}

func (h *HeapMin) moveToBottom(nodeIndex int) {
	heapSize := len(h.tree)
	parentIndex := nodeIndex
	for {
		leftChildIndex := 2*parentIndex + 1
		rightChildIndex := 2*parentIndex + 2

		// swap with the smallest child
		smallestChildIndex := parentIndex
		if heapSize > rightChildIndex && h.tree[rightChildIndex].Value < h.tree[parentIndex].Value { // swap with right child
			smallestChildIndex = rightChildIndex
		}
		if heapSize > leftChildIndex && h.tree[leftChildIndex].Value < h.tree[parentIndex].Value {
			smallestChildIndex = leftChildIndex
		}

		if smallestChildIndex == parentIndex {
			break
		}

		h.swap(parentIndex, smallestChildIndex)
		parentIndex = smallestChildIndex
	}
}

func (h *HeapMin) GetRoot() string {
	if len(h.tree) == 0 {
		return ""
	}

	root := h.tree[0]
	h.Delete(0)

	return root.Name
}
