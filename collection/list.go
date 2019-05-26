package collection

type List struct {
	l []int
}

func (l *List) Push(i int) {
	l.l = append(l.l, i)
}

func (l *List) Pop() int {
	v := l.l[len(l.l)-1]
	l.l = l.l[:len(l.l)-1]
	return v
}

func (l *List) Shift() int {
	v := l.l[0]
	l.l = l.l[1:]
	return v
}

func (l *List) Unshift(i int) {
	l.l = append([]int{i}, l.l...)
}

func (l *List) Len() int {
	return len(l.l)
}
