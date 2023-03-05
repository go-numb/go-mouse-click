package main

import (
	"fmt"
	"time"

	mouse "github.com/go-numb/go-mouse-click"
)

func main() {
	x, y := mouse.Setting()

	fmt.Println("取得完了")
	time.Sleep(time.Second)
	fmt.Printf("取得座標 x: %d, y: %d\n", x, y)

	time.Sleep(3 * time.Second)
	fmt.Println("3秒後にクリックしてみます")
	time.Sleep(3 * time.Second)

	mouse.Click(x, y, 1)
}
