package memdb

import "todoapp/internal/usecase/repo"

func CreateRepos() (repo.UserRepo, repo.TodoRepo) {
	return NewUserRepo(), NewTodoRepo()
}
