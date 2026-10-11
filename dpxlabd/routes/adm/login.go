package adm

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

/*
/adm/login
   - POST: username, hash(pw)
   => encrypted,signed JWT token
*/

// ASSERT METHOD?
func (r *AdmRoute) HandleLogin(ctx *dpxlabd.RouteHandlerContext) {
	routes.Error(ctx, http.StatusNotImplemented)
}

// ctx.W.Header().Set("Content-Type", routes.HTMLContentType)

/*
http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})


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

func (r *AdmRoute) HandleLogout(ctx *dpxlabd.RouteHandlerContext) {

	/*
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   "",
			Expires: time.Now(),
		})
	*/
}
