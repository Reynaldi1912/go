package routes

import (
	"database/sql"
	"net/http"

	"github.com/reynaldi1912/go/controllers"
)

func MapRoute(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controllers.SilentControllers())
	server.HandleFunc("/index", controllers.IndexPage(db))
}
