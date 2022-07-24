package models

import (
	"encoding/json"
	"time"
)

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

func (u DBUser) MarshalBinary() ([]byte, error) {
	bytes, err := json.Marshal(u)
	return bytes, err
}
