package event_dispatcher

import (
	"context"
	"testing"
	"time"
)

func TestDispatching(t *testing.T) {
	ctx := context.Background()
	ed := NewEventDispatcher()
	ed.RegisterSubscriber(
		NewOrderSubscriber(),
		[]ListeningEvent{
			OrderListeningEvent,
		},
		7,
	)

	ed.RegisterSubscriber(
		NewFunSubscriber(),
		[]ListeningEvent{
			OrderListeningEvent,
			FunListeningEvent,
		},
		2,
	)

	//ed.Dispatch(ctx, NewOrderEvent())
	//ed.AsyncDispatchWithWait(ctx, NewOrderEvent())
	ed.AsyncDispatch(ctx, NewOrderEvent())
	//ed.Dispatch(ctx, NewFunEvent())
	//ed.AsyncDispatchWithWait(ctx, NewFunEvent())
	ed.AsyncDispatch(ctx, NewFunEvent())
	time.Sleep(2 * time.Second)
}
