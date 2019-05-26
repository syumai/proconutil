package collection

type Queue struct {
	q []int
}

func (q *Queue) Push(i int) {
	q.q = append(q.q, i)
}

func (q *Queue) Pop() int {
	v := q.q[len(q.q)-1]
	q.q = q.q[:len(q.q)-1]
	return v
}

func (q *Queue) Shift() int {
	v := q.q[0]
	q.q = q.q[1:]
	return v
}

func (q *Queue) Unshift(i int) {
	q.q = append([]int{i}, q.q...)
}

func (q *Queue) Len() int {
	return len(q.q)
}
