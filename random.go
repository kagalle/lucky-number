package lucky_number_package

import (
	"math/rand"
)

func RandomNumber() int {
	return rand.Intn(100)
}
