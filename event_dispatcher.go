package event_dispatcher

import (
	"context"
	"slices"
	"sort"
	"sync"
)

type EventDispatcher struct {
	subscribers     []Subscriber
	withPriority    bool
	prioritisedList bool
}

func NewEventDispatcher(withPriority bool) *EventDispatcher {
	return &EventDispatcher{
		withPriority: withPriority,
	}
}

func (ed *EventDispatcher) RegisterSubscriber(subscriber Subscriber, events []ListeningEvent, subscriberPriority int64) {
	bs := subscriber.GetBaseSubscriber()
	bs.SetListenEvents(events)
	bs.SetPriority(subscriberPriority)
	ed.subscribers = append(ed.subscribers, subscriber)
}

func (ed *EventDispatcher) Dispatch(ctx context.Context, event Event) {
	ed.sortSubscribersByPriority()
	for _, sub := range ed.subscribers {
		if len(sub.GetBaseSubscriber().GetListenEvents()) > 0 {
			if slices.Contains(sub.GetBaseSubscriber().GetListenEvents(), event.GetName()) {
				sub.Handle(ctx, event)
			}
		}
	}
}

func (ed *EventDispatcher) CustomDispatch(ctx context.Context, event Event, customDispatchFunction func(ctx context.Context, event Event) error) error {
	ed.sortSubscribersByPriority()
	err := customDispatchFunction(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (ed *EventDispatcher) AsyncDispatch(ctx context.Context, event Event) {
	ed.sortSubscribersByPriority()
	for _, sub := range ed.subscribers {
		if len(sub.GetBaseSubscriber().GetListenEvents()) > 0 {
			if slices.Contains(sub.GetBaseSubscriber().GetListenEvents(), event.GetName()) {
				go sub.Handle(ctx, event)
			}
		}
	}
}

func (ed *EventDispatcher) AsyncDispatchWithWait(ctx context.Context, event Event) {
	ed.sortSubscribersByPriority()
	wg := sync.WaitGroup{}
	for _, sub := range ed.subscribers {
		if len(sub.GetBaseSubscriber().GetListenEvents()) > 0 {
			if slices.Contains(sub.GetBaseSubscriber().GetListenEvents(), event.GetName()) {
				wg.Add(1)
				go sub.HandleWithWait(ctx, event, &wg)
			}
		}
	}
	wg.Wait()
}

func (ed *EventDispatcher) sortSubscribersByPriority() {
	if ed.withPriority && !ed.prioritisedList {
		sort.Slice(ed.subscribers, func(i, j int) bool {
			return ed.subscribers[i].GetBaseSubscriber().GetPriority() > ed.subscribers[j].GetBaseSubscriber().GetPriority()
		})
		ed.prioritisedList = true
	}
}
