package event_dispatcher

var FunListeningEvent ListeningEvent = "event.fun"

func NewFunEvent() *FunEvent {
	return &FunEvent{
		BaseEvent: &BaseEvent{},
		name:      FunListeningEvent,
	}
}

type FunEvent struct {
	*BaseEvent
	name ListeningEvent
}

func (fe *FunEvent) GetName() ListeningEvent {
	return fe.name
}

func (fe *FunEvent) GetData() EventData {
	return fe.BaseEvent.GetData()
}
