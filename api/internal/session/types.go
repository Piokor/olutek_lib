package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id        uuid.UUID `database:"id"`
	UserId    int64     `database:"user_id"`
	ExpiresAt time.Time `database:"expires_at"`
	Data      string    `database:"data"`
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
