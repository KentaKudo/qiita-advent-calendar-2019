package main

import (
	"context"
	"errors"
	"testing"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type serverTestSuite struct {
	sut     *server
	ctrl    *gomock.Controller
	todoMgr *MockTodoManager
}

func newServerTestSuite(t *testing.T) serverTestSuite {
	ctrl := gomock.NewController(t)
	todoMgr := NewMockTodoManager(ctrl)
	return serverTestSuite{
		sut:     &server{todoMgr: todoMgr},
		ctrl:    ctrl,
		todoMgr: todoMgr,
	}
}

func TestServer_CreateTodo(t *testing.T) {
	t.Run("project a new todo", func(t *testing.T) {
		s := newServerTestSuite(t)
		defer s.ctrl.Finish()

		input := &service.CreateTodoRequest{
			Todo: &service.Todo{
				Title:       "foo todo",
				Description: "foo description",
			},
		}

		todoID := uuid.New().String()
		want := &service.CreateTodoResponse{
			Success: true,
			Id:      todoID,
		}

		s.todoMgr.EXPECT().
			projectTodo(todo{
				title:       input.Todo.Title,
				description: input.Todo.Description,
			}).
			Return(todoID, nil)

		got, err := s.sut.CreateTodo(context.Background(), input)
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("error in projection", func(t *testing.T) {
		s := newServerTestSuite(t)
		defer s.ctrl.Finish()

		input := &service.CreateTodoRequest{
			Todo: &service.Todo{
				Title:       "foo todo",
				Description: "foo description",
			},
		}

		s.todoMgr.EXPECT().
			projectTodo(gomock.Any()).
			Return("", errors.New("foo error"))

		_, err := s.sut.CreateTodo(context.Background(), input)
		require.Error(t, err)
	})
}
