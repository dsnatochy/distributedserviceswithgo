package server

import (
	"fmt"
	"testing"
)

func TestLog_Read(t *testing.T) {
	var a uint
	a = 6
	var b uint
	b = 10
	var i uint
	i = a-b
	i = i+7
	fmt.Println(i)

}