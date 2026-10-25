package auth

import (
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

//TODO: js PoW solver im web-worker

// reqs: client hash,
type PoWSolution struct {
	ChallengeID uuid.UUID
	Nonce       uint64
	Hash        []byte
}

func PoWUnmarshalSolution(solution_str string) (*PoWSolution, error) {
	solution_json, err := base64.RawURLEncoding.DecodeString(solution_str)
	if err != nil {
		return nil, err
	}

	var solution = &PoWSolution{}
	if err := json.Unmarshal(solution_json, &solution); err != nil {
		return nil, err
	}

	return solution, nil
}

func PoWVerifySolution(solution *PoWSolution, challenge *PoWChallenge) error {
	if solution.ChallengeID != challenge.ID {
		return errors.New("PoW: challenge-id mismatch")
	}

	// ExpireDate?
	if time.Now().Unix() > challenge.ExpireDate {
		return errors.New("PoW: challenge expired")

	}

	// rehash
	server_hash := argon2.IDKey(
		append(
			challenge.Salt,
			[]byte(fmt.Sprintf("%d", solution.Nonce))...,
		),
		challenge.Salt,
		pow_time, pow_mem, pow_threads, pow_keylenght,
	)

	// verify difficulty

	if !hasLeadingZeroBits(server_hash, pow_difficulty) {
		return errors.New("PoW: difficulty not met")
	}

	// verify hashes
	if subtle.ConstantTimeCompare(server_hash, solution.Hash) != 1 {
		return errors.New("PoW: hash mismatch")
	}

	return nil
}

func hasLeadingZeroBits(h []byte, n int) bool {
	full, rem := n/8, n%8
	if len(h) < full+btoi(rem > 0) {
		return false
	}
	for i := 0; i < full; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return rem == 0 || h[full]>>(8-rem) == 0
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
