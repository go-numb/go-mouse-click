package mouse

import (
	"log"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

// Click around x and y, var around is determined by random.
func Click(x, y, around int) {
	start := time.Now()
	defer func() {
		log.Default().Println("Click() elapsed time: ", time.Since(start))
	}()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := rand.Int()
	if n&2 == 0 {
		x += r.Intn(around)
		x -= r.Intn(around)
	} else {
		x -= r.Intn(around)
		y += r.Intn(around)
	}

	log.Printf("ランダム加算後座標 x: %d, y: %d\n", x, y)

	robotgo.Move(x, y)
	robotgo.Click()
}
