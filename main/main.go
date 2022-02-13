package main

import (
	"fmt"
	"loc-system-order/controller"
)

func main() {

	r := controller.InitRounter()
	fmt.Println("runing")
	r.Run()
}
