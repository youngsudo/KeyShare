package models

import "time"

type DBUser struct {
	ID         int       `db:"id"`
	UserID     int64     `db:"user_id"`
	Address    string    `db:"address"`
	Account    string    `db:"account"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	PrivateKey string    `db:"privatekey"`
	Time       time.Time `db:"time"`
}
