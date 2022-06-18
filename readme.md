# Onkyo Remote

By Colin Murphy, 2022

Onkyo-remote exposes the eISCP protocol Onkyo home audio equipment uses into
a REST API. It utilizes a self-built library that sends and receives eISCP
commands from a receiver using documentation from Onkyo. 

In the future it will also provide a JavaScript-based web remote that can be
used to control the receiver instead of interfacing with an API.

## Usage

Set the receiver's IP address, and optionally the port the web server listens
on.

    export RECEIVER_IP="192.168.1.180"
    export HTTP_PORT="8080"

Start the service...

    go run main.go


Open a web browser and navigate to http://localhost:8080/api/status. You should
get a response returning the general status of the receiver.

For more information, please see the [Wiki][0].

## Debugging

I use Wireshark for debugging the network activity between this application
and my stereo receiver. To only inspect eISCP traffic, use this filter:

    tcp.port == 60128 && ip.addr == receiver_ip_addr

[0]: https://github.com/colinmurphy1/onkyo-remote/wiki
