package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) insert(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if nil == l.front {
		l.front = item
		l.front.Prev = item
		l.front.Next = item
	} else {
		item.Prev = l.front.Prev
		item.Next = l.front
		l.front.Prev = item
		item.Prev.Next = item
	}

	l.len++

	return item
}
func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if nil == l.front {
		l.front = item
		l.back = item
	} else {
		item.Next = l.front
		l.front.Prev = item
		l.front = item
	}

	l.len++

	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if nil == l.front {
		l.front = item
		l.back = item
	} else {
		item.Prev = l.back
		l.back.Next = item
		l.back = item
	}

	l.len++

	return item
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == l.front {
		l.front = l.front.Next
	}
	if i == l.back {
		l.back = l.back.Prev
	}
	if nil != i.Prev {
		i.Prev.Next = i.Next
	}
	if nil != i.Next {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
