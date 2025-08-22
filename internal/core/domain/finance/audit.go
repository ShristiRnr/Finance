package finance

import (
	"errors"
	"time"
)

type AuditEvent struct {
	ID           string
	UserID       string
	Action       string
	Timestamp    time.Time
	Details      string
	ResourceType string
	ResourceID   string
}

// AuditEventRepository is the port (interface) for persistence.
type AuditEventRepository interface {
	Save(event AuditEvent) (AuditEvent, error)
	FindAll() ([]AuditEvent, error)  
	FindByResource(resourceType, resourceID string) ([]AuditEvent, error)
	FindByUser(userID string) ([]AuditEvent, error)
}

func (e AuditEvent) Validate() error {
	if e.UserID == "" {
		return errors.New("userID cannot be empty")
	}
	if e.Action == "" {
		return errors.New("action cannot be empty")
	}
	if e.Timestamp.IsZero() {
		return errors.New("timestamp cannot be zero")
	}
	if e.ResourceType == "" || e.ResourceID == "" {
	 	return errors.New("resource type and id must be set")
	}
	return nil
}
