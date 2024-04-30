package event_dispatcher

type ListeningEvent string
type EventData map[string]any

type BaseEvent struct {
	data EventData
}

func (be *BaseEvent) GetData() EventData {
	return be.data
}

func (be *BaseEvent) SetData(data EventData) {
	be.data = data
}

type Event interface {
	GetName() ListeningEvent
	GetData() EventData
}
