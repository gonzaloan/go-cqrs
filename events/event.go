package events

import (
	"context"
	"cqrs/models"
)

type EventStore interface {
	Close()
	PublishCreatedFeed(ctx context.Context, feed *models.Feed) error
	SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error)
	OnCreateFeed(f func(message CreatedFeedMessage)) error //Callback which is going to react when new message was created
}

var eventStore EventStore

func Close() {
	eventStore.Close()
}
func PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishCreatedFeed(ctx, feed)
}

func SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error) {
	return eventStore.SubscribeCreatedFeed(ctx)
}

func OnCreateFeed(f func(message CreatedFeedMessage)) error {
	return eventStore.OnCreateFeed(f)
}
