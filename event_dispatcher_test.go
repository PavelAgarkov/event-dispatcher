package event_dispatcher

import (
	"context"
	"fmt"
	"testing"
)

func TestDispatching(t *testing.T) {
	ctx := context.Background()
	ed := NewEventDispatcher(true)
	ed.RegisterSubscriber(
		NewOrderSubscriber(),
		[]ListeningEvent{
			OrderListeningEvent,
		},
		func(ctx context.Context, event Event) error {
			fmt.Println(event.GetData(), "_fun")
			fmt.Println("------->", event.GetName())
			return nil
		},
		7,
	)

	ed.RegisterSubscriber(
		NewFunSubscriber(),
		[]ListeningEvent{
			OrderListeningEvent,
			FunListeningEvent,
		},
		func(ctx context.Context, event Event) error {
			fmt.Println(event.GetData(), "_order")
			return nil
		},
		2,
	)

	ed.Dispatch(ctx, NewOrderEvent())
	//ed.AsyncDispatchWithWait(ctx, NewOrderEvent())
	//ed.AsyncDispatch(ctx, NewOrderEvent())
	ed.Dispatch(ctx, NewFunEvent())
	//ed.AsyncDispatchWithWait(ctx, NewFunEvent())
	//ed.AsyncDispatch(ctx, NewFunEvent())
	//time.Sleep(2 * time.Second)
}
