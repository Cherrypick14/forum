package auth

import (
	"fmt"
	"net/http"
	"net/url"

	utils "forum/backend/util"
)

func GoogleSignUp(w http.ResponseWriter, r *http.Request) {
	// generate state with cookie?

	state := generateStateCookie(w, "signup")

	// set the Google OAuth 2.0 authorization URL
	redirectUrl := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=openid email profile&state=%s&prompt=select_account&access_type=offline",
		GoogleAuthUrl, utils.GoogleClientID, url.QueryEscape(BaseUrl+"/auth/google/signin/callback"), state)

	//  set CORS specifying domain
	w.Header().Set("Access-Control-Allow-Origin", BaseUrl)

	//  redirect user
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}
