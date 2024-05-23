package mouse_test

import (
	"testing"

	"github.com/go-numb/go-mouse-click"
)

func TestMouse(t *testing.T) {
	mouse.Click(100, 100, 10)
}
