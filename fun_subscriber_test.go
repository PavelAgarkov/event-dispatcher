package event_dispatcher

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type FunSubscriber struct {
	*BaseSubscriber
	*OrderEvent
}

func NewFunSubscriber() *FunSubscriber {
	return &FunSubscriber{
		BaseSubscriber: &BaseSubscriber{},
		OrderEvent:     NewOrderEvent(),
	}
}

func (fs *FunSubscriber) Handle(ctx context.Context, event Event) {
	fmt.Println(event.GetData(), "_fun")
	fmt.Println("------->", fs.OrderEvent.GetName())
	return
}

func (fs *FunSubscriber) GetBaseSubscriber() *BaseSubscriber {
	return fs.BaseSubscriber
}

func (fs *FunSubscriber) HandleWithWait(ctx context.Context, event Event, wg *sync.WaitGroup) {
	fmt.Println(event.GetData(), "_order")
	time.Sleep(2 * time.Second)
	wg.Done()
	return
}
