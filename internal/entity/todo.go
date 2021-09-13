package entity

type Todo struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Content  string `json:"content"`
	IsDone   bool   `json:"isDone"`
}

func (t *Todo) Toggle() {
	t.IsDone = !t.IsDone
}

func (t *Todo) Equals(otherTodo *Todo) bool {
	if t.Username != otherTodo.Username {
		return false
	}
	if t.Content != otherTodo.Content {
		return false
	}
	if t.IsDone != otherTodo.IsDone {
		return false
	}
	return true
}
