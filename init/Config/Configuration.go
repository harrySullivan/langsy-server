package Config

import (
	"github.com/joho/godotenv"
)

func Load() {
	godotenv.Overload()
}
