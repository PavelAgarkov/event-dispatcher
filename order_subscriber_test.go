package event_dispatcher

import (
	"context"
	"fmt"
	"sync"
)

type OrderSubscriber struct {
	*BaseSubscriber
}

func NewOrderSubscriber() *OrderSubscriber {
	return &OrderSubscriber{&BaseSubscriber{}}
}

func (os *OrderSubscriber) Handle(ctx context.Context, event Event) {
	_ = os.handler(ctx, event)
}

func (os *OrderSubscriber) GetBaseSubscriber() *BaseSubscriber {
	return os.BaseSubscriber
}

func (os *OrderSubscriber) HandleWithWait(ctx context.Context, event Event, wg *sync.WaitGroup) {
	fmt.Println(event.GetData(), "_order")
	wg.Done()
	return
}
