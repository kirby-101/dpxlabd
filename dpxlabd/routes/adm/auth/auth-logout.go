package auth

import "github.com/kirby-101/dpxlab/dpxlabd"

func (r *AuthRoute) HandleGetLogout(ctx *dpxlabd.RouteHandlerContext) {}

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
