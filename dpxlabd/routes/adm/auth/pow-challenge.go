package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Challenge an User-ID / Session-ID / IP + Timestamp binden
// Ablauf Datum
// nach verify, muss die challenge invalidiert werden bzw muss eine neue gelten
// was ist subtle.ConstantTimeCompare()
type PoWChallenge struct {
	ID         uuid.UUID
	Salt       []byte
	ExpireDate int64

	// userID? hash(ip) zum monitoring oder binden?
}

func PoWGenerateChallenge() (*PoWChallenge, error) {
	var (
		salt = make([]byte, 16)

		challenge = &PoWChallenge{
			ID:         uuid.New(),
			Salt:       salt,
			ExpireDate: time.Now().Add(pow_expire_period).Unix(),
		}
	)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return challenge, nil
}

func PoWMarshalChallenge(c *PoWChallenge) (string, error) {
	if c == nil {
		return "", errors.New("challenge is nil")
	}

	if len(c.Salt) == 0 {
		// NOPE
		return "", errors.New("len(Salt) == 0!")
	}

	if c.ExpireDate == 0 {
		return "", errors.New("ExpireDate == 0!")
	}

	challenge_json, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	var challenge_str = base64.RawURLEncoding.EncodeToString(challenge_json)

	return challenge_str, nil
}
