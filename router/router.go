package router

import (
	"workspace_booking/config"
	"workspace_booking/controller"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// No Auth routes
	api.Post("/sign-up", controller.Register)
	api.Post("/sign-in", controller.Login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.GetJWTSecret()),
	}))

	app.Use(func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		cookie := c.Cookies(email)
		var err error
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSecret()), nil
		})
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Invalid Access",
			})
		}

		if token.Raw == cookie {
			c.Locals("verify", "true")
		} else {
			c.Locals("verify", "false")
		}
		return c.Next()
	})

	// Authorization routes

	api.Get("/roles", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllRoles(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Post("/roles", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateRole(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Post("/users", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateUser(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/users", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.GetUsers(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/users/:id", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.GetUser(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Delete("/users/:id", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.DeleteUser(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/logout", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.Logout(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Post("/book_workspace", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateBooking(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/workspace_details", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.WorkSpacesDetails(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* Building API's */
	api.Post("/buildings", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateBuilding(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/buildings", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllBuildings(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* City API's */
	api.Post("/cities", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateCity(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/cities", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllCities(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* Location API's */
	api.Post("/locations", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateLocation(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/locations", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllLocations(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* Floor API's */
	api.Post("/floors", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateFloor(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/floors", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllFloors(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})
}
