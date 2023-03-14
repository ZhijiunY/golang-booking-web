package dbrepo

import (
	"github.com/ZhijiunY/booking-web/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
		end_date, room_id, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	m.DB.Exec(stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone)

	return nil
}
