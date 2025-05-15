// Modifikasi 1: Menambahkan komentar untuk commit point
// Modifikasi 2: Komentar kedua untuk point commit
// Modifikasi 3: Komentar ketiga untuk point commit
package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hallo World!")
	})

	log.Fatal(app.Listen(":3000"))
}
