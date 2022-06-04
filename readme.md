# Go Onkyo Library

## Examples

Create a new receiver at IP 192.168.1.180, and disconnect when the program is done with it

    dev := Onkyo("192.168.1.180")
    defer dev.Disconnect()


Power the receiver on or off

    dev.SetPower(true)
    dev.SetPower(false)

Set the volume to 25, or mute it

    dev.SetVolume(25)
    dev.SetMute(true)

Get information about the current source in a hashmap

    info, _ := x.GetAudioInfo())
    fmt.Println(info)

This will return:  
`map[InputFormat: InputPort:NETWORK InputSignalChannel:Stereo ListenMode:2.0 ch OutputSignalChannel: SamplingFrequency:]`

## Useful information

I use Wireshark for debugging this library. To only inspect EISCP traffic, use this filter:

    tcp.port == 60128 && ip.dst == 192.168.1.180
