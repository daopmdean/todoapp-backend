package entity

type Todo struct {
	Content string
	IsDone  bool
}

func (t *Todo) Check() {
	t.IsDone = !t.IsDone
}
