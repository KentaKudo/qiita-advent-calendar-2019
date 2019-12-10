package main

import (
	"context"
	"testing"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uw-labs/substrate"
)

var _ substrate.SynchronousMessageSink = mockMessageSink{}

type mockMessageSink struct{}

func (s mockMessageSink) PublishMessage(_ context.Context, _ substrate.Message) error {
	return nil
}

func (s mockMessageSink) Close() error {
	return nil
}

func (s mockMessageSink) Status() (*substrate.Status, error) {
	return &substrate.Status{}, nil
}

type serverTestSuite struct {
	sut  *server
	ctrl *gomock.Controller
}

func newServerTestSuite(t *testing.T) serverTestSuite {
	ctrl := gomock.NewController(t)
	return serverTestSuite{
		sut:  newServer(mockMessageSink{}),
		ctrl: ctrl,
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

		got, err := s.sut.CreateTodo(context.Background(), input)
		require.NoError(t, err)
		assert.NotEmpty(t, got.Id)
	})
}
