package main

//credit teivah for this and not having to implement basic data structures go should just have
type PriorityQueue[T comparable] struct {
  items      []T
  comparator func(a, b T) int
}

func NewPriorityQueue[T comparable](comparator func(a, b T) int) PriorityQueue[T] {
  return PriorityQueue[T]{comparator: comparator}
}

func (pq *PriorityQueue[T]) Push(item T) {
  pq.items = append(pq.items, item)
  pq.heapUp(len(pq.items) - 1)
}

func (pq *PriorityQueue[T]) Pop() T {
  top := pq.items[0]
  lastIndex := len(pq.items) - 1
  pq.items[0], pq.items[lastIndex] = pq.items[lastIndex], pq.items[0]
  pq.items = pq.items[:lastIndex]
  pq.heapDown(0)
  return top
}

func (pq *PriorityQueue[T]) Peek() T {
  return pq.items[0]
}

func (pq *PriorityQueue[T]) Len() int {
  return len(pq.items)
}

func (pq *PriorityQueue[T]) isEmpty() bool {
  return len(pq.items) == 0
}

func (pq *PriorityQueue[T]) heapUp(index int) {
  for index > 0 {
    parentIndex := (index - 1) / 2
    if pq.comparator(pq.items[index], pq.items[parentIndex]) < 0 {
      pq.items[index], pq.items[parentIndex] = pq.items[parentIndex], pq.items[index]
      index = parentIndex
    } else {
      break
    }
  }
}
func (pq *PriorityQueue[T]) heapDown(index int) {
  for {
    left := 2*index + 1
    right := 2*index + 2
    smallest := index
    if left < len(pq.items) && pq.comparator(pq.items[left], pq.items[smallest]) < 0 {
      smallest = left
    }
    
    if right < len(pq.items) && pq.comparator(pq.items[right], pq.items[smallest]) < 0 {
      smallest = right
    }
    
    if smallest != index {
      pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
      index = smallest
    } else {
      break
    }
  }
}
