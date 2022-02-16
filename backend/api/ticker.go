package api

import "github.com/labstack/echo"

type addTickerRequest struct {
	CompanyName string `json:"companyname"`
	Ticker      string `json:"ticker"`
}

func (server *Server) addTicker(ctx *echo.Context) echo.HandlerFunc {

}
