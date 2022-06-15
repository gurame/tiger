package api

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gurame/tiger/app/api/routes"
	"github.com/gurame/tiger/app/core/services"
	"github.com/gurame/tiger/app/infrastructure/db"
	"github.com/gurame/tiger/app/infrastructure/db/repositories"
	_ "github.com/lib/pq"
)

func Run() {

	//db
	sqlDb := db.Connection()

	sellerRepository := repositories.NewSellerRepository(sqlDb)
	sellerService := services.NewSellerService(sellerRepository)

	//web
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Welcome tiger api!"))
	})
	routes.ConfigureSeller(app, sellerService)

	app.Get("/docs/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}
