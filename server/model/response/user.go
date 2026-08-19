package response

import (
	"server/model/database"

	"github.com/gofrs/uuid"
)

type Login struct {
	User                 database.User `json:"user"`
	AccessToken          string        `json:"access_token"`
	AccessTokenExpiresAt int64         `json:"access_token_expires"` // 过期时间
}

type UserCard struct {
	UUID      uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	Address   string    `json:"address"`
	Signature string    `json:"signature"`
}
