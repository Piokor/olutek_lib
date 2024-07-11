package session

import "github.com/Piokor/olutek_lib/internal/database"

func (session *Session) Insert(service *database.DbService) error {
	return service.Insert("session", session, "id", "user_id", "expires_at")
}
