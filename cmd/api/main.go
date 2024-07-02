package main

import (
	t_h "backend-hex-arch/cmd/api/handlers/ticket"
	t_r "backend-hex-arch/pkg/repository/mongo/ticket"
	t_s "backend-hex-arch/pkg/service/ticket"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ticketH := t_h.NewTicketHandler(t_s.NewPlayerService(t_r.NewRepository()))
	e := echo.New()
	e.GET("/ticket/:id", ticketH.Get)
	e.PUT("/ticket/:id", ticketH.Edit)
	e.DELETE("/ticket/:id", ticketH.Delete)
	e.GET("/ticket", ticketH.GetAll)
	e.POST("/ticket", ticketH.Create)
	e.Logger.Fatal(e.Start(":8080"))
}
