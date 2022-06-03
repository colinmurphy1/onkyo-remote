package main

import "fmt"

func main() {

	x := &Connection{}
	x.ip = "192.168.1.180"

	// Connect to the receiver at 192.168.1.180
	x.Connect()
	defer x.Disconnect()

	//fmt.Println(x.GetPower())

	fmt.Println(x.SetPower(true))

	//fmt.Println(x.SetVolume(10))

}
