package utils

import "fmt"

// Stole this from github

type MaxHeap struct {
    heapArray []int
    size      int
    maxsize   int
}

func NewMaxHeap(maxsize int) *MaxHeap {
    MaxHeap := &MaxHeap{
        heapArray: []int{},
        size:      0,
        maxsize:   maxsize,
    }
    return MaxHeap
}

func (m *MaxHeap) leaf(index int) bool {
    if index >= (m.size/2) && index <= m.size {
        return true
    }
    return false
}

func (m *MaxHeap) parent(index int) int {
    return (index - 1) / 2
}

func (m *MaxHeap) leftchild(index int) int {
    return 2*index + 1
}

func (m *MaxHeap) rightchild(index int) int {
    return 2*index + 2
}

func (m *MaxHeap) Insert(item int) error {
    if m.size >= m.maxsize {
        return fmt.Errorf("Heap is full")
    }
    m.heapArray = append(m.heapArray, item)
    m.size++
    m.upHeapify(m.size - 1)
    return nil
}

func (m *MaxHeap) swap(first, second int) {
    temp := m.heapArray[first]
    m.heapArray[first] = m.heapArray[second]
    m.heapArray[second] = temp
}

func (m *MaxHeap) upHeapify(index int) {
    for m.heapArray[index] > m.heapArray[m.parent(index)] {
        m.swap(index, m.parent(index))
        index = m.parent(index)
    }
}

func (m *MaxHeap) downHeapify(current int) {
    if m.leaf(current) {
        return
    }
    largest := current
    leftChildIndex := m.leftchild(current)
    rightRightIndex := m.rightchild(current)
    //If current is smallest then return
    if leftChildIndex < m.size && m.heapArray[leftChildIndex] > m.heapArray[largest] {
        largest = leftChildIndex
    }
    if rightRightIndex < m.size && m.heapArray[rightRightIndex] > m.heapArray[largest] {
        largest = rightRightIndex
    }
    if largest != current {
        m.swap(current, largest)
        m.downHeapify(largest)
    }
    return
}

func (m *MaxHeap) buildMaxHeap() {
    for index := ((m.size / 2) - 1); index >= 0; index-- {
        m.downHeapify(index)
    }
}

func (m *MaxHeap) Remove() int {
    top := m.heapArray[0]
    m.heapArray[0] = m.heapArray[m.size-1]
    m.heapArray = m.heapArray[:(m.size)-1]
    m.size--
    m.downHeapify(0)
    return top
}
