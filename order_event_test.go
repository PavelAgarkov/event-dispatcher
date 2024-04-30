package event_dispatcher

var OrderListeningEvent ListeningEvent = "event.order"

func NewOrderEvent() *OrderEvent {
	return &OrderEvent{
		BaseEvent: &BaseEvent{},
		name:      OrderListeningEvent,
	}
}

type OrderEvent struct {
	*BaseEvent
	name ListeningEvent
}

func (oe *OrderEvent) GetName() ListeningEvent {
	return oe.name
}

func (oe *OrderEvent) GetData() EventData {
	return oe.BaseEvent.GetData()
}
