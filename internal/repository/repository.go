package repository

import "github.com/ZhijiunY/booking-web/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
