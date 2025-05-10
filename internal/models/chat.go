package models

import "time"

// Message represents a chat message
type Message struct {
	ID         uint   `gorm:"primaryKey"`
	SenderID   uint   // References the User who sent the message
	ReceiverID uint   // References the User or Organization receiving the message
	Content    string `gorm:"not null"`
	CreatedAt  time.Time
}
