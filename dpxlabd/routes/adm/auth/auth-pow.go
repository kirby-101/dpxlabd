package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

const (
	pow_time          uint32 = 1
	pow_mem           uint32 = 16 * 1024
	pow_threads       uint8  = 2
	pow_keylenght     uint32 = 32
	pow_difficulty           = 3 // len(leading zeros)
	pow_expire_period        = 5 * time.Second
)

func (r *AuthRoute) HandleGetPoWChallenge(ctx *dpxlabd.RouteHandlerContext) {
	// serve pow
	challenge, err := PoWGenerateChallenge()
	if err != nil {
		ctx.Orch.Logger.Error("PoW: generate challenge", err.Error())
		routes.Error(ctx, http.StatusInternalServerError)

		return
	}

	challenge_str, err := PoWMarshalChallenge(challenge)
	if err != nil {
		ctx.Orch.Logger.Error("PoW: marshal challenge", err.Error())
		routes.Error(ctx, http.StatusInternalServerError)

		return
	}

	// track active Challenges+
	r.cleanActiveChallenges()
	r.activeChallenges[challenge.ID] = challenge

	ctx.W.Header().Set("Content-Type", "application/json")
	ctx.W.Write([]byte(
		fmt.Sprintf(`{ "challenge": "%s" }`, challenge_str),
	))

	return
}

func (r *AuthRoute) cleanActiveChallenges() {
	for _, challenge := range r.activeChallenges {
		// is expired?
		if time.Now().Unix() > challenge.ExpireDate {
			delete(
				r.activeChallenges,
				challenge.ID,
			)
		}
	}
}
