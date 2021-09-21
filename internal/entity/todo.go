package entity

type Todo struct {
	ID       string `json:"id" bson:"id,omitempty" gorm:"primaryKey"`
	Username string `json:"username" bson:"username,omitempty"`
	Content  string `json:"content" bson:"content,omitempty"`
	IsDone   bool   `json:"isDone" bson:"isDone,omitempty"`
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
