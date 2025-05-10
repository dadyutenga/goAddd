package models

import "time"

// User represents a user in the system (can be superadmin, admin, or regular user)
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
	Role      string `gorm:"default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Organization represents an organization managed by an admin
type Organization struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	AdminID   uint   // References the User who is the admin
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Member represents a member of an organization with a specific role
type Member struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   // References the User
	OrganizationID uint   // References the Organization
	Role           string `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
