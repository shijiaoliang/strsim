package main

import (
	"fmt"
	"github.com/shijiaoliang/strsim"
)

func main() {
	fmt.Println(strsim.IsMatch("大江东去", []string{"呼啦哗啦", "大江东去浪"}, 0.6))
}
