package heap

type MinIntHeap []int

func NewFrom(x []int) *MinIntHeap {
	dst := make([]int, len(x))
	copy(dst, x)

	h := MinIntHeap(dst)
	n := len(h) - 1

	for i := n/2 - 1; i >= 0; i-- {
		h.heapifyDown(i, n)
	}
	return &h
}

func (h MinIntHeap) Len() int {
	return len(h)
}

func (h *MinIntHeap) Push(item int) {
	*h = append(*h, item)
	h.heapifyUp()
}

func (h *MinIntHeap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	h.heapifyDown(0, n)

	old := *h
	x := old[n]
	*h = old[0:n]

	return x
}

func (h MinIntHeap) leftChildIndex(i int) int {
	return 2*i + 1
}

func (h MinIntHeap) rightChildIndex(i int) int {
	return 2*i + 2
}

func (h MinIntHeap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h MinIntHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinIntHeap) heapifyUp() {
	i := len(h) - 1
	for {
		p := h.parentIndex(i)
		if p == i || h[p] < h[i] {
			break
		}

		h.swap(i, p)
		i = p
	}
}

func (h MinIntHeap) heapifyDown(i, n int) {
	for {
		leftChild := h.leftChildIndex(i)
		if leftChild >= n {
			break
		}

		j := leftChild

		rightChild := h.rightChildIndex(i)
		if rightChild < n && h[rightChild] < h[leftChild] {
			j = rightChild
		}

		if h[j] > h[i] {
			break
		}

		h.swap(i, j)
		i = j
	}
}
