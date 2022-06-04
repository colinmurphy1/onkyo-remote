package main

import "fmt"

func main() {

	// Connect to the receiver at 192.168.1.180
	onkyo := Onkyo("192.168.1.180")

	fmt.Println(onkyo.GetAudioInfo())

	defer onkyo.Disconnect()

	//fmt.Println(x.SetVolume(10))
	//fmt.Println(x.SetMute(false))
	//fmt.Println(x.GetAudioInfo())
}
