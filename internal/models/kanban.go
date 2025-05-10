package models

import "time"

// Board represents a Kanban board
type Board struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	OrganizationID uint   // References the Organization
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Column represents a column on a Kanban board
type Column struct {
	ID        uint   `gorm:"primaryKey"`
	BoardID   uint   // References the Board
	Name      string `gorm:"not null"`
	Position  int    // Order of the column
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Card represents a task card on a Kanban board
type Card struct {
	ID          uint   `gorm:"primaryKey"`
	ColumnID    uint   // References the Column
	Title       string `gorm:"not null"`
	Description string
	AssigneeID  uint // References the User
	Position    int  // Order within the column
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
