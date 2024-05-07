package event_dispatcher

import (
	"context"
	"sync"
)

type BaseSubscriber struct {
	handler      func(ctx context.Context, event Event) error
	listenEvents []ListeningEvent
	priority     int64
}

func (bs *BaseSubscriber) SetListenEvents(events []ListeningEvent) {
	bs.listenEvents = events
}

func (bs *BaseSubscriber) GetListenEvents() []ListeningEvent {
	return bs.listenEvents
}

func (bs *BaseSubscriber) GetPriority() int64 {
	return bs.priority
}

func (bs *BaseSubscriber) SetPriority(priority int64) {
	bs.priority = priority
}

type Subscriber interface {
	Handle(ctx context.Context, event Event)
	HandleWithWait(ctx context.Context, event Event, wg *sync.WaitGroup)
	GetBaseSubscriber() *BaseSubscriber
}
