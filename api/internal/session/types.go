package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id        uuid.UUID `db:"id"`
	UserId    int64     `db:"user_id"`
	ExpiresAt time.Time `db:"expires_at"`
	Data      string    `db:"data"`
}

func New(userId int64) *Session {
	sessionId := uuid.New()
	sessionExpireDate := time.Now().AddDate(0, 1, 0) // 1 month
	return &(Session{
		sessionId,
		userId,
		sessionExpireDate,
		"",
	})
}
