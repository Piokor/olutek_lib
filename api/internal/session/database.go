package session

import "github.com/Piokor/olutek_lib/internal/database"

func insert(session *Session, service *database.DbService) error {
	upsertUpdate := `
		ON CONFLICT (user_id) 
		DO UPDATE
		SET 
			id = EXCLUDED.id,
			expires_at = EXCLUDED.expires_at,
			data = EXCLUDED.data		
	`
	return service.InsertWithOptions("session", session, upsertUpdate, "id", "user_id", "expires_at")
}
