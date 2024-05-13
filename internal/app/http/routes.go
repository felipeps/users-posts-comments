package http

import (
	"net/http"

	"github.com/felipeps/user-post-comments/internal/app/handlers"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.GetUsersPosts)
}
