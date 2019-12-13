package main

import (
	"context"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/envelope"
	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/event"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/uw-labs/substrate"
)

type actionEventHandler struct {
	todoMgr todoManager
}

func newActionEventHandler(todoMgr todoManager) actionEventHandler {
	return actionEventHandler{todoMgr: todoMgr}
}

func (h actionEventHandler) handle(ctx context.Context, msg substrate.Message) error {
	var env envelope.Event
	if err := proto.Unmarshal(msg.Data(), &env); err != nil {
		return errors.Wrap(err, "failed to unmarshal message")
	}

	if types.Is(env.Payload, &event.CreateTodoActionEvent{}) {
		var ev event.CreateTodoActionEvent
		if err := types.UnmarshalAny(env.Payload, &ev); err != nil {
			return errors.Wrap(err, "failed to unmarshal payload")
		}

		if err := h.todoMgr.projectTodo(todo{
			id:          ev.Id,
			title:       ev.Title,
			description: ev.Description,
		}); err != nil {
			return errors.Wrap(err, "failed to project a todo")
		}
	}

	return nil
}
