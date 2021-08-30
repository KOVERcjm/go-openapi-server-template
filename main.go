package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"kovercheng/driver"
	"kovercheng/handler"
	. "kovercheng/middleware"
	"os"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		ReadTimeout: 60 * time.Second,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{"msg": err.Error()})
		}})

	app.Use(cors.New(), logger.New(logger.Config{
		/**
		 *  Colourful HTTP logger for console use
		 */
		Format: "${time} ${magenta}${method}${reset} ${url} ${body} | ${cyan}${status}${reset} | ${error}\n",
		/**
		 *  JSON format HTTP logger for log collecting use
		 */
		//Format: "{\"time\":\"${time}\",\"method\":\"${method}\",\"url\":\"${url}\",\"body\":\"${body}\",\"response_status\":\"${status}\",\"error\":\"${error}\"}\n",
	}))
	handler.Route(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "cannot find endpoint",
		})
	})

	/**
	 *  The following code will start an HTTP server.
	 */
	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		Logger.Fatalf("Server Terminated: %v", err)
	}
	/**
	 *  Use the following code for starting HTTPS server.
	 */
	//if err := app.ListenTLS(os.Getenv("SERVER_URL"), "server/cert/certificate.crt", "server/cert/private.key"); err != nil {
	//	Logger.Fatalf("Server Terminated: %v", err)
	//}

	_ = driver.CloseConnection()
	_ = app.Shutdown()
}
