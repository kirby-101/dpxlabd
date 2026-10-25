package auth

import (
	"time"

	"github.com/google/uuid"
	"github.com/kirby-101/dpxlab/dpxlabd"
)

const (
	tokenCookieName = "dpxlab_sid"
	tokenExpireTime = 2 * time.Hour
)

type Role uint

const (
	RoleAnyone Role = iota
	RoleAdmin
	RoleUser
	RoleCustomer
)

type Token struct {
	ID         uuid.UUID
	Username   string
	Role       Role // admin / user
	ExpireDate time.Time
}

func NewToken(username string, role Role) *Token {
	return &Token{
		ID:         uuid.New(),
		Username:   username,
		Role:       role,
		ExpireDate: time.Now().Add(tokenExpireTime),
	}
}

func (t Token) Sign() {}

// via cookie
func SetClientToken(ctx *dpxlabd.RouteHandlerContext, t *Token) {
	// sign token

	// set token
}

// via cookie
func GetClientToken(ctx *dpxlabd.RouteHandlerContext) *Token {
	cookie, err := ctx.R.Cookie(tokenCookieName)
	if err != nil {
		return nil
	}

	_ = cookie

	// expired?

	// server signatur intakt?

	// return token

	return nil
}
