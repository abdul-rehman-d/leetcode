package min_stack

type MinStack struct {
	arr  []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		arr:  []int{},
		mins: []int{},
	}
}

func (this *MinStack) Push(val int) {
	m := val
	if len(this.arr) != 0 {
		if lastMin := this.GetMin(); lastMin < m {
			m = lastMin
		}
	}
	this.arr = append(this.arr, val)
	this.mins = append(this.mins, m)
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
