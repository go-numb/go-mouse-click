package mouse

import (
	"context"
	"fmt"
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
		fmt.Println("click")
	})
	defer hook.End()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	fmt.Println(WAITWORD)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i+1)
		time.Sleep(time.Second)
	}

	fmt.Println(STARTWORD)

L:
	for ev := range evChan {
		select {
		case <-ctx.Done():
			break L
		case <-ticker.C:
			x = int(stat.Mean(xs, nil))
			y = int(stat.Mean(ys, nil))
			if x < 0 || y < 0 {
				fmt.Println(FALSE)
				continue
			}
			fmt.Println(x, y)

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
