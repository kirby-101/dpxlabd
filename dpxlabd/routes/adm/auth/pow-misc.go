package auth

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

// returns nonce, hash
func clientSolveChallenge(challenge *PoWChallenge) (uint64, []byte) {
	for nonce := uint64(0); ; nonce++ {
		client_hash := argon2.IDKey(
			append(challenge.Salt, []byte(fmt.Sprintf("%d", nonce))...),
			challenge.Salt,
			pow_time, pow_mem, pow_threads, pow_keylenght,
		)

		if hasLeadingZeroBits(client_hash, pow_difficulty) {
			return nonce, client_hash
		}

		if nonce > 1000000 {
			panic("nonce > 1000000")
		}
	}
}

func buildClientSolution(id uuid.UUID, nonce uint64, hash []byte) []byte {
	var solution = &PoWSolution{
		ChallengeID: id,
		Nonce:       nonce,
		Hash:        hash,
	}

	solution_json, err := json.Marshal(solution)
	if err != nil {
		panic(err)
	}

	return solution_json
}
