package finance

import "time"

type AuditEvent struct {
	ID           string
	UserID       string
	Action       string
	Timestamp    time.Time
	Details      string
	ResourceType string
	ResourceID   string
}
