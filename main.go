package main

import (
	"fmt"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
)

func main() {

	// Connect to the receiver at 192.168.1.180
	onkyo := eiscp.Onkyo("192.168.1.180")

	fmt.Println(onkyo.GetAudioInfo())

	defer onkyo.Disconnect()

	//fmt.Println(x.SetVolume(10))
	//fmt.Println(x.SetMute(false))
	//fmt.Println(x.GetAudioInfo())
}
