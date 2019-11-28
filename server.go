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

func (s *server) CreateTodo(ctx context.Context, req *service.CreateTodoRequest) (*service.CreateTodoResponse, error) {
	id, err := s.todoMgr.projectTodo(todo{
		title:       req.Todo.Title,
		description: req.Todo.Description,
	})
	if err != nil {
		return nil, err
	}

	return &service.CreateTodoResponse{
		Success: true,
		Id:      id,
	}, nil
}

func (*server) ListTodos(context.Context, *service.ListTodosRequest) (*service.ListTodosResponse, error) {
	return &service.ListTodosResponse{}, nil
}
