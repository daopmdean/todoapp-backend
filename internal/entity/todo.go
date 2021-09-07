package entity

type Todo struct {
	Username string
	Content  string
	IsDone   bool
}

func (t *Todo) Toggle() {
	t.IsDone = !t.IsDone
}
