package implement_queue_using_stacks

// https://leetcode.com/problems/implement-queue-using-stacks/

type MyQueue struct {
	stack1 []int
	stack2 []int
	size   int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: []int{},
		stack2: []int{},
		size:   0,
	}
}

func (this *MyQueue) Push(x int) {
	for i := 0; i < this.size; i++ {
		this.stack2[this.size-1-i] = this.stack1[i]
	}

	if this.size+1 > len(this.stack1) {
		this.stack1 = append(this.stack1, 0)
	}

	this.stack1[0] = x

	for i := 0; i < this.size; i++ {
		this.stack1[this.size-i] = this.stack2[i]
	}

	if this.size+1 > len(this.stack2) {
		this.stack2 = append(this.stack2, 0)
	}

	this.size++
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	this.size--

	return this.stack1[this.size]
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}

	return this.stack1[this.size-1]
}

func (this *MyQueue) Empty() bool {
	return this.size < 1
}
