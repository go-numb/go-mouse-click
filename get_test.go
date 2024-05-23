package mouse_test

import (
	"fmt"
	"os"
	"os/signal"
	"testing"

	hook "github.com/go-numb/go-mouse-click"
	"github.com/stretchr/testify/assert"
)

func TestGetMousePostions(t *testing.T) {
	// プログラムの停止、Ctrl+Cで停止 を監視する
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go hook.Setting()

	<-ch
	os.Exit(0)
}

func TestGetXYs(t *testing.T) {
	// プログラムの停止、Ctrl+Cで停止 を監視する

	n := 10
	timelimit := 30

	xs, ys := hook.GetXYs(n, timelimit)
	assert.Equal(t, len(xs), n)
	assert.Equal(t, len(ys), n)

	for i := 0; i < len(xs); i++ {
		fmt.Printf("x: %f\n", xs[i])
	}

	for i := 0; i < len(ys); i++ {
		fmt.Printf("x: %f\n", ys[i])
	}

}
