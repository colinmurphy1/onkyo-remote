# Onkyo Remote

By Colin Murphy, 2022

Onkyo-remote exposes the eISCP protocol Onkyo/Integra/newer Pioneer home audio
equipment uses into a REST API. It utilizes a self-built library that sends and
receives eISCP commands from a receiver using documentation from Onkyo. 

In the future it will also provide a JavaScript-based web remote that can be
used to control the receiver instead of interfacing with an API.

## Usage

Make a copy of the `config.example.yaml` file and change the IP address to
the IPv4 address of your Onkyo/Integra receiver. You can also make other
setting changes, for example to disable the eISCP logging. 

Once done, you may start the service like so:

    go run main.go -config ./config.yaml

Optionally, you may also compile the software and run it:

    go build
    ./onkyo-remote -config ./config.yaml

Open a web browser and navigate to http://localhost:8080/api/status. You should
get a response returning the general status of the receiver.

For more information, please see the [Wiki][0].

## Debugging

I use Wireshark for debugging the network activity between this application
and my stereo receiver. To only inspect eISCP traffic, use this filter:

    tcp.port == 60128 && ip.addr == receiver_ip_addr

[0]: https://github.com/colinmurphy1/onkyo-remote/wiki
