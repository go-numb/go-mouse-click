package main

import (
	"fmt"
	"time"

	mouse "github.com/go-numb/go-mouse-click"
)

func _main() {
	x, y := mouse.Setting()

	fmt.Println("取得完了")
	time.Sleep(time.Second)
	fmt.Printf("取得座標 x: %d, y: %d\n", x, y)

	time.Sleep(3 * time.Second)
	fmt.Println("3秒後にクリックしてみます")
	time.Sleep(3 * time.Second)

	mouse.Click(x, y, 1)
}

func main() {
	x, y := mouse.GetXYs(10, 20)

	fmt.Println("取得完了")
	time.Sleep(time.Second)

	fmt.Println(len(x), len(y), "gets")
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
}
