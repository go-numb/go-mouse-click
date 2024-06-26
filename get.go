package mouse

import (
	"context"
	"fmt"
	"log"
	"time"

	hook "github.com/robotn/gohook"
	"gonum.org/v1/gonum/stat"
)

const (
	WAITWORD = "5秒後に、指定する単一ボタンの座標を取得し始めます。開始後、おおよそ10秒間、指定するボタンの中心部を何度かクリックしてください。"

	STARTWORD = "開始します。"
	FALSE     = "座標が取得できません、オブジェクト内でクリックしてください"
)

// Setting 指定秒数の間にマウスが滞在する座標を取得し、平均を返す
func Setting() (x, y int) {
	var (
		xs, ys []float64
	)

	evChan := hook.Start()
	hook.Register(hook.MouseDown, []string{"w"}, func(e hook.Event) {
		log.Println("click")
	})
	defer hook.End()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	log.Println(WAITWORD)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i+1)
		time.Sleep(time.Second)
	}

	log.Println(STARTWORD)

L:
	for ev := range evChan {
		select {
		case <-ctx.Done():
			break L
		case <-ticker.C:
			x = int(stat.Mean(xs, nil))
			y = int(stat.Mean(ys, nil))
			if x < 0 || y < 0 {
				log.Println(FALSE)
				continue
			}
			log.Println(x, y)

		default:
			if ev.Kind == hook.MouseUp || ev.Kind == hook.MouseDown {
				xs = append(xs, float64(ev.X))
				ys = append(ys, float64(ev.Y))
			}
		}
	}

	x = int(stat.Mean(xs, nil))
	y = int(stat.Mean(ys, nil))

	return
}

// GetXYs クリック座標を指定回数分取得し返す
func GetXYs(n, timelimit int) (xs, ys []float64) {
	evChan := hook.Start()
	hook.Register(hook.MouseDown, []string{"w"}, func(e hook.Event) {
		log.Println("click")
	})
	defer hook.End()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timelimit)*time.Second)
	defer cancel()

	time.Sleep(time.Second)
	log.Println(STARTWORD)

L:
	for ev := range evChan {
		select {
		case <-ctx.Done():
			break L

		default:
			if !(ev.Kind == hook.MouseUp) {
				continue
			}

			log.Printf("link: %d, %d, %d", ev.Kind, ev.X, ev.Y)
			log.Printf("len: %d, %d", len(xs), len(ys))

			if len(xs) >= n || len(ys) >= n {
				break L
			} else if len(xs) != len(ys) {
				break L
			}

			xs = append(xs, float64(ev.X))
			ys = append(ys, float64(ev.Y))
		}
	}

	// check length
	if len(xs) > len(ys) {
		xs = xs[:len(ys)]
	} else if len(ys) > len(xs) {
		ys = ys[:len(xs)]
	}

	return
}

type Corners struct {
	X1 int
	Y1 int

	X2 int
	Y2 int
}

// GetFourCorners クリック座標を指定回数分取得し返す
func GetFourCorners() *Corners {
	evChan := hook.Start()
	hook.Register(hook.MouseDown, []string{"w"}, func(e hook.Event) {
		log.Println("click")
	})
	defer hook.End()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println(STARTWORD)

	c := new(Corners)

L:
	for ev := range evChan {
		select {
		case <-ctx.Done():
			break L

		default:
			if ev.Kind == hook.MouseHold {
				c.X1, c.Y1 = int(ev.X), int(ev.Y)
			} else if ev.Kind == hook.MouseDown {
				c.X2, c.Y2 = int(ev.X), int(ev.Y)
			}

			if c._done() {
				break L
			}
		}
	}

	return c
}

func (c *Corners) _done() bool {
	if c.X1 != 0 && c.Y1 != 0 && c.X2 != 0 && c.Y2 != 0 {
		return true
	}

	return false
}

func (c *Corners) MinX() int {
	if c.X1 > c.X2 {
		return c.X2
	} else if c.X1 == c.X2 {
		return c.X1
	}
	return c.X1
}

func (c *Corners) MinY() int {
	if c.Y1 > c.Y2 {
		return c.Y2
	} else if c.Y1 == c.Y2 {
		return c.Y1
	}
	return c.Y1
}

func (c *Corners) MaxX() int {
	if c.X1 > c.X2 {
		return c.X1
	} else if c.X1 == c.X2 {
		return c.X1
	}
	return c.X2
}

func (c *Corners) MaxY() int {
	if c.Y1 > c.Y2 {
		return c.Y1
	} else if c.Y1 == c.Y2 {
		return c.Y1
	}
	return c.Y2
}
