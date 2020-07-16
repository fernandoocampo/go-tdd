package domain

import "context"

// Notifier notifies to an external application about any event.
type Notifier interface {
	Notify(ctx context.Context, message Message) error
}
