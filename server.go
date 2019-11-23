package main

import (
	"context"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
)

var _ service.TodoAPIServer = (*server)(nil)

type server struct{}

func (*server) GetTodo(context.Context, *service.GetTodoRequest) (*service.GetTodoResponse, error) {
	return &service.GetTodoResponse{
		Todo: &service.Todo{
			Title:       "clean your desk!",
			Description: "clean up the piles of documents from the desk to make more space.",
		},
	}, nil
}

func (*server) CreateTodo(context.Context, *service.CreateTodoRequest) (*service.CreateTodoResponse, error) {
	return &service.CreateTodoResponse{}, nil
}

func (*server) ListTodos(context.Context, *service.ListTodosRequest) (*service.ListTodosResponse, error) {
	return &service.ListTodosResponse{}, nil
}
