package adm

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Role uint

const (
	RoleAdmin Role = iota
	RoleUser
)

type Token struct {
	Username   string
	Role       Role
	SessionID  uuid.UUID // uuid.NewString()
	ExpireDate time.Time // time.Now().Add(12 *time.Hours)
}

func (t Token) Export() {}

func (t Token) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Token) Unmarshal(token []byte) error {
	return json.Unmarshal(token, t)
}

func (t Token) IsExpired() bool {
	return t.ExpireDate.Before(time.Now())
}

func (t Token) sign()    {}
func (t Token) encrypt() {}
