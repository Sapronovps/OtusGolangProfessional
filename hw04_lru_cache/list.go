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
	Key   Key
}

type list struct {
	List
	length int // zero value == 0
	head   *ListItem
	tail   *ListItem
}

func NewList() List {
	return new(list)
}

// Добавление элемента в начало списка.
func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: nil, Next: nil}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		newItem.Next = l.head
		l.head.Prev = newItem
		l.head = newItem
	}
	l.length++

	return newItem
}

// Добавление элемента в конец списка.
func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: nil, Next: nil}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		newItem.Prev = l.tail
		l.tail.Next = newItem
		l.tail = newItem
	}
	l.length++

	return newItem
}

// Удаление элемента из списка по значению.
func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	// Если узел - это голова списка.
	if i.Prev == nil {
		l.head = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	// Если узел - это хвост списка
	if i.Next == nil {
		l.tail = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

// Длина двусвязного списка.
func (l *list) Len() int {
	return l.length
}

// Возвращает первый элемент двусвязного списка.
func (l *list) Front() *ListItem {
	return l.head
}

// Возвращает последний элемент двусвязного списка.
func (l *list) Back() *ListItem {
	return l.tail
}
