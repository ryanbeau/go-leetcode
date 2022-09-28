package solution

import "container/heap"

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

type MedianFinder struct {
	left  IntHeap
	right IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  IntHeap{},
		right: IntHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() < this.right.Len() {
		if this.right.Len() > 0 && this.right[0] < num {
			heap.Push(&this.left, -heap.Pop(&this.right).(int))
			heap.Push(&this.right, num)
		} else {
			heap.Push(&this.left, -num)
		}
	} else {
		if this.left.Len() > 0 && -this.left[0] > num {
			heap.Push(&this.right, -heap.Pop(&this.left).(int))
			heap.Push(&this.left, -num)
		} else {
			heap.Push(&this.right, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64(-this.left[0]+this.right[0]) / 2
	}
	return float64(this.right[0])
}
