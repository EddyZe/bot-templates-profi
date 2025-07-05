package entity

import "time"

type User struct {
	Id         int64     `db:"id" json:"id" csv:"id"`
	Username   string    `db:"username" json:"username" csv:"username"`
	TelegramId int64     `db:"telegram_id" json:"telegram_id" csv:"telegram_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at" csv:"created_at"`
}
