package event

import "github.com/xyy0411/bleachVSnaruto/core/audio"

// Bus 事件总线，用于发布和订阅事件
type Bus struct {
	subscribers map[Type][]Handler
}

// NewBus 创建新的事件总线
func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[Type][]Handler),
	}
}

// Subscribe 订阅指定类型的事件
func (b *Bus) Subscribe(eventType Type, handler Handler) {
	b.subscribers[eventType] = append(b.subscribers[eventType], handler)
}

// Publish 发布事件
func (b *Bus) Publish(event Event) {
	if handlers, exists := b.subscribers[event.Type]; exists {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

// PublishAudio 发布音频事件
func (b *Bus) PublishAudio(characterID string, audioEvent audio.Event) {
	b.Publish(Event{
		Type:        Audio,
		CharacterID: characterID,
		AudioEvent:  audioEvent,
	})
}
