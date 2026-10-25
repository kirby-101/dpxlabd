package auth

import (
	"encoding/json"

	"github.com/kirby-101/dpxlab/dpxlabd"
)

/*
 * GET /adm/auth => login page !cookie banner bzw akzeptieren!
 * POST /adm/auth pow_challenge=XXX & pow_counter=XXX & username=str &  pw=(argon2id) => token
 */

func (r *AuthRoute) HandlePostLogin(ctx *dpxlabd.RouteHandlerContext) {
	var resp = struct {
		Username     string
		PasswordHash string
		PoWSolution
	}{}

	//
	if err := json.NewDecoder(ctx.R.Body).Decode(&resp); err != nil {
		//!
	}

	// get challenge
	challenge, oke := r.activeChallenges[resp.ChallengeID]
	if !oke {
		//!
	}

	// Verify Solution
	if err := PoWVerifySolution(&resp.PoWSolution, challenge); err != nil {
		//!
	}

	// PoW solved, Username? ; PasswordHash?

	// SetLogin(): set signed token cookie

}

/*
// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
*/
