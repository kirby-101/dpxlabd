package auth

import (
	"encoding/base64"
	"testing"
	"time"
)

func BenchmarkProofOfWorkDifficulty(b *testing.B) {
	// New Challenge
	challenge, err := PoWGenerateChallenge()
	if err != nil {
		b.Fatal(err)
	}

	// Solve Challenge
	var (
		nonce, hash   = clientSolveChallenge(challenge)
		solution_json = []byte(base64.RawURLEncoding.EncodeToString(
			buildClientSolution(
				challenge.ID,
				nonce,
				hash,
			),
		))
	)

	solution, err := PoWUnmarshalSolution(string(solution_json))
	if err != nil {
		b.Fatal(err)
	}

	/*
	* Benchmark
	 */

	for b.Loop() {
		// Verify Solution
		_ = PoWVerifySolution(solution, challenge)
	}
}

func TestProofOfWork(t *testing.T) {
	// New Challenge
	challenge, err := PoWGenerateChallenge()
	if err != nil {
		t.Error(err)
	}

	// Solve Challenge
	t.Log("trying to solve")
	var (
		nonce, hash   = clientSolveChallenge(challenge)
		solution_json = []byte(base64.RawURLEncoding.EncodeToString(
			buildClientSolution(
				challenge.ID,
				nonce,
				hash,
			),
		))
	)

	solution, err := PoWUnmarshalSolution(string(solution_json))
	if err != nil {
		t.Error(err)
	}

	ts := time.Now()

	// Verify Solution
	if err := PoWVerifySolution(solution, challenge); err != nil {
		t.Error(err)
	}

	t.Log("solution verified in:", ts.Sub(time.Now()))
}
