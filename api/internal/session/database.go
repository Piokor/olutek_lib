package session

import (
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/google/uuid"
)

func insert(session *Session, dbService *database.DbService) error {
	upsertUpdate := `
		ON CONFLICT (user_id) 
		DO UPDATE
		SET 
			id = EXCLUDED.id,
			expires_at = EXCLUDED.expires_at,
			data = EXCLUDED.data		
	`
	return dbService.InsertWithOptions("session", session, upsertUpdate, "id", "user_id", "expires_at")
}

func selectUserId(sessionId uuid.UUID, dbService *database.DbService) (int64, error) {
	queryStr := `
		SELECT user_id FROM session 
		WHERE id=$1
	`
	var userId int64
	err := dbService.Client.QueryRowx(queryStr, sessionId).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
