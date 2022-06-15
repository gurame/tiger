package api

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gurame/tiger/app/api/routes"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Run() {

	db, err := sql.Open("postgres", "dbname=TigerDb user=admin password=admin sslmode=disable")
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	boil.SetDB(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Welcome tiger api!"))
	})

	routes.ConfigureSeller(app)

	log.Fatal(app.Listen(":3000"))
}
