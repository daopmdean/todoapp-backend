package api_rest_test

import (
	"fmt"
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
	"todoapp/test/api_rest_test/internal"
)

func TestGetTodoList(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	todoRepo := serverRouter.GetTodoRepo()
	seedDataForTodoListTest(todoRepo)
	tokenStr := internal.GenAccessToken("daopham")
	request := internal.Request{
		Method: "GET",
		Path:   "/api/todos",
		Headers: map[string]string{
			"Authorization": "Bearer " + tokenStr,
		},
	}

	response := serverRouter.HandleRequest(request)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}

	todos := todoRepo.GetTodoListOf("daopham")
	for _, todo := range todos {
		fmt.Println(todo)
	}
	if len(todos) != 2 {
		t.Errorf("expected todo list with len 2, got len %d", len(todos))
	}

	expectedTodo1 := entity.Todo{
		Username: "daopham",
		Content:  "Walking",
		IsDone:   false,
	}
	expectedTodo2 := entity.Todo{
		Username: "daopham",
		Content:  "Running",
		IsDone:   true,
	}
	if !internal.CheckIfTodoListContainsTodo(todos, expectedTodo1) {
		t.Errorf("todos dont contain expectedTodo1: %v", expectedTodo1)
	}
	if !internal.CheckIfTodoListContainsTodo(todos, expectedTodo2) {
		t.Errorf("todos dont contain expectedTodo2: %v", expectedTodo2)
	}
}

func seedDataForTodoListTest(todoRepo repo.TodoRepo) {
	fakeData1 := entity.Todo{
		Username: "daopham",
		Content:  "Walking",
		IsDone:   false,
	}
	fakeData2 := entity.Todo{
		Username: "daopham",
		Content:  "Running",
		IsDone:   true,
	}
	fakeData3 := entity.Todo{
		Username: "hungpham",
		Content:  "Learning",
		IsDone:   false,
	}
	todoRepo.CreateTodo(&fakeData1)
	todoRepo.CreateTodo(&fakeData2)
	todoRepo.CreateTodo(&fakeData3)
}
