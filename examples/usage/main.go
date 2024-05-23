package main

import (
	"log"
	"time"

	mouse "github.com/go-numb/go-mouse-click"
)

func main() {
	x, y := mouse.Setting()

	log.Default().Println("取得完了")
	time.Sleep(time.Second)
	log.Default().Printf("取得座標 x: %d, y: %d\n", x, y)

	time.Sleep(3 * time.Second)
	log.Default().Println("3秒後にクリックしてみます")
	time.Sleep(3 * time.Second)

	mouse.Click(x, y, 1)
}

func __main() {
	x, y := mouse.GetXYs(10, 20)

	log.Default().Println("取得完了")
	time.Sleep(time.Second)

	log.Default().Println(len(x), len(y), "gets")
	log.Default().Printf("%v\n", x)
	log.Default().Printf("%v\n", y)
}

func _main() {
	c := mouse.GetFourCorners()
	log.Default().Println("start: ", c.X1, c.Y1)
	log.Default().Println("end: ", c.X2, c.Y2)
}
