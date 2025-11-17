package shared

type EventPublisher interface {
	Publish(event Event) error
}
