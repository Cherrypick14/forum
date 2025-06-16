package route

import (
	"net/http"

	"forum/backend/auth"
	authmiddleware "forum/backend/auth-middleware"
	"forum/backend/handlers"
)

func InitRoutes() *http.ServeMux {
	r := http.NewServeMux()

	fs := http.FileServer(http.Dir("./frontend"))
	r.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	uploadsFs := http.FileServer(http.Dir("./uploads"))
	r.Handle("/uploads/", http.StripPrefix("/uploads/", uploadsFs))

	//  routes

	r.HandleFunc("/home", authmiddleware.Authenticate(handlers.IndexHandler))
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/sign-in", handlers.LoginHandler)
	r.HandleFunc("/sign-up", handlers.SignupHandler)
	r.HandleFunc("/logout",authmiddleware.Authenticate(handlers.LogoutHandler))

	r.HandleFunc("/auth/google/signin", auth.GoogleSignIn)
	r.HandleFunc("/auth/google/signup", auth.GoogleSignUp)
	r.HandleFunc("/auth/google/signin/callback", auth.GoogleSignInCallback)

	r.HandleFunc("/auth/github/signup", auth.GithubSignUp)
	r.HandleFunc("/auth/github/signin", auth.GitHubSignIn)
	r.HandleFunc("/auth/github/callback", auth.GithubCallback)

	r.HandleFunc("/upload", authmiddleware.Authenticate(handlers.CreatePost))
	r.HandleFunc("/filter", handlers.FilterPosts)
	r.HandleFunc("/comments", authmiddleware.Authenticate(handlers.CommentHandler))
	r.HandleFunc("/reaction", authmiddleware.Authenticate(handlers.ReactionHandler))
	r.HandleFunc("/likes", authmiddleware.Authenticate(handlers.ReactionHandler))
	r.HandleFunc("/dilikes", authmiddleware.Authenticate(handlers.ReactionHandler))
	r.HandleFunc("/validate", handlers.ValidateInputHandler)

	return r
}
