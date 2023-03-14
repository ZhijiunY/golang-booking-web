package models

import "time"

// Reservation holds reservation data
type Reservation struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// Users is the user model
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// Rooms is the room model
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

// Restriction is the restirction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdateAt        time.Time
}

// Reservation is the reservation model
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdateAt  time.Time
}
