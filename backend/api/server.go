package api

import (
	"github.com/labstack/echo/v4"
	db "github.com/proggeguden/stock-app/backend/db/sqlc"
)

// Server serves HTP requests for our Stock App
type Server struct {
	store  *db.Store
	router *echo.Echo
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := echo.New()

	router.POST("/tickers", server.addTicker)
}
