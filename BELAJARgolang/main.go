// Modifikasi 1: Menambahkan komentar untuk commit point
// Modifikasi 2: Komentar kedua untuk point commit
// Modifikasi 3: Komentar ketiga untuk point commit
package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Modifikasi 4: Variabel dummy untuk point commit
// Modifikasi 5: Komentar kelima untuk point commit
// Modifikasi 6: Komentar keenam untuk point commit
// Modifikasi 7: Komentar ketujuh untuk point commit
// Modifikasi 8: Komentar kedelapan untuk point commit
	dummyVar := 0
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hallo World!")
	})

	log.Fatal(app.Listen(":3000"))
}
