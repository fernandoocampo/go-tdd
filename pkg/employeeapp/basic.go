package employeeapp

import (
	"context"
	"log"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/google/uuid"
)

// basicService implements employee service
type basicService struct {
	repository domain.EmployeeRepository
	notifier   domain.Notifier
}

// NewBasicService creates a employee service using basic employee service impl.
func NewBasicService(newRepository domain.EmployeeRepository) Service {
	return &basicService{
		repository: newRepository,
	}
}

// NewBasicServiceWithNotifier creates a employee service using basic employee service impl and
// contains a notifier to notify to an employee when there is a new event in their account.
func NewBasicServiceWithNotifier(newRepository domain.EmployeeRepository, newNotifier domain.Notifier) Service {
	return &basicService{
		notifier:   newNotifier,
		repository: newRepository,
	}
}

// Add validates and save a new employee.
func (b *basicService) Add(ctx context.Context, newemployee domain.Employee) (string, error) {
	newemployee.ID = uuid.New().String()
	err := b.repository.Save(ctx, newemployee)
	if b.notifier != nil && err == nil {
		newMessage := domain.Message{
			Subject: "Hey",
			To:      newemployee.Email,
			From:    "anybody@somewhere.com",
			Body:    "Good morning there",
		}
		notifierErr := b.notifier.Notify(ctx, newMessage)
		if notifierErr != nil {
			log.Printf("unexpected error: %s sending message %+v to %q ", notifierErr, newMessage, newemployee.Email)
		}
	}
	return newemployee.ID, err
}
