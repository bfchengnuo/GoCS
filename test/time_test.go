package test

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	fmt.Println(time.Now().Format("15:04:05\n"))
}
