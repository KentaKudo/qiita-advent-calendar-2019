package main

import (
	"context"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
)

var _ service.TodoAPIServer = (*server)(nil)

type (
	todo struct {
		title       string
		description string
	}
	todoManager interface {
		projectTodo(todo) (string, error)
	}

	server struct {
		todoMgr todoManager
	}
)

func (s *server) GetTodo(context.Context, *service.GetTodoRequest) (*service.GetTodoResponse, error) {
	return &service.GetTodoResponse{
		Todo: &service.Todo{
			Title:       "clean your desk!",
			Description: "clean up the piles of documents from the desk to make more space.",
		},
	}, nil
}

func (s *server) CreateTodo(context.Context, *service.CreateTodoRequest) (*service.CreateTodoResponse, error) {
	return &service.CreateTodoResponse{}, nil
}

func (*server) ListTodos(context.Context, *service.ListTodosRequest) (*service.ListTodosResponse, error) {
	return &service.ListTodosResponse{}, nil
}
