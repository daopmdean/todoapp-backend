package usecase

import "todoapp/internal/usecase/repo"

func NewAuthorizationUsecase(todoRepo repo.TodoRepo) *AuthorizationUsecase {
	return &AuthorizationUsecase{todoRepo}
}

type AuthorizationUsecase struct {
	todoRepo repo.TodoRepo
}

func (au *AuthorizationUsecase) CheckIfUserIsAllowedToModifyTodo(username, todoID string) bool {
	todoOwner := au.todoRepo.GetTodoByID(todoID)

	if todoOwner == nil {
		return false
	}

	if todoOwner.Username == username {
		return true
	}
	return false
}
