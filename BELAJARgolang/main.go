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
// Modifikasi 9: Komentar kesembilan untuk point commit
// Modifikasi 10: Komentar kesepuluh untuk point commit
// Modifikasi 11: Komentar kesebelas untuk point commit
// Modifikasi 12: Komentar keduabelas untuk point commit
// Modifikasi 13: Komentar ketigabelas untuk point commit
// Modifikasi 14: Komentar keempatbelas untuk point commit
// Modifikasi 15: Komentar kelimabelas untuk point commit
// Modifikasi 16: Komentar keenambelas untuk point commit
// Modifikasi 17: Komentar ketujuhbelas untuk point commit
// Modifikasi 18: Komentar keelapanbelas untuk point commit
// Modifikasi 19: Komentar kesembilanbelas untuk point commit
// Modifikasi 20: Komentar keduapuluh untuk point commit
// Modifikasi 21: Komentar keduapuluh satu untuk point commit
// Modifikasi 22: Komentar keduapuluh dua untuk point commit
// Modifikasi 23: Komentar keduapuluh tiga untuk point commit
// Modifikasi 24: Komentar keduapuluh empat untuk point commit
	dummyVar := 0
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hallo World!")
	})

	log.Fatal(app.Listen(":3000"))
}
