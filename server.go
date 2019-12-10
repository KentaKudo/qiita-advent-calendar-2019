package main

import (
	"context"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/envelope"
	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/event"
	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/uw-labs/substrate"
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
		sink substrate.SynchronousMessageSink
	}
)

func newServer(
	sink substrate.SynchronousMessageSink,
) *server {
	return &server{
		sink: sink,
	}
}

func (s *server) GetTodo(context.Context, *service.GetTodoRequest) (*service.GetTodoResponse, error) {
	return &service.GetTodoResponse{
		Todo: &service.Todo{
			Title:       "clean your desk!",
			Description: "clean up the piles of documents from the desk to make more space.",
		},
	}, nil
}

func (s *server) CreateTodo(ctx context.Context, req *service.CreateTodoRequest) (*service.CreateTodoResponse, error) {
	todoID := uuid.New().String()
	ev := &event.CreateTodoActionEvent{
		Id:          todoID,
		Title:       req.Todo.Title,
		Description: req.Todo.Description,
	}

	any, err := types.MarshalAny(ev)
	if err != nil {
		return nil, err
	}

	env := envelope.Event{
		Id:        uuid.New().String(),
		Timestamp: types.TimestampNow(),
		Payload:   any,
	}

	b, err := proto.Marshal(&env)
	if err != nil {
		return nil, err
	}

	if err := s.sink.PublishMessage(ctx, message(b)); err != nil {
		return nil, err
	}

	return &service.CreateTodoResponse{
		Success: true,
		Id:      todoID,
	}, nil
}

func (*server) ListTodos(context.Context, *service.ListTodosRequest) (*service.ListTodosResponse, error) {
	return &service.ListTodosResponse{}, nil
}
