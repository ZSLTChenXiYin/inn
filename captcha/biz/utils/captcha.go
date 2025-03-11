package utils

import (
	"fmt"
	"time"

	"math/rand"
)

func GenerateCaptcha() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", random.Int31n(1000000))
}
