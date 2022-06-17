# Onkyo Remote

By Colin Murphy, 2022

This application exposes the eISCP protocol Onkyo home audio equipment uses into
a REST API. It utilizes a self-built library that communicates with a receiver.

## Usage

Set the receiver's IP address, and optionally the port the web server listens
on.

    export RECEIVER_IP="192.168.1.180"
    export HTTP_PORT="8080"

Start the service...

    go run main.go


Open a web browser and navigate to http://localhost:8080/api/status. You should
get a response returning the general status of the receiver.


## Debugging

I use Wireshark for debugging the network activity between this application
and my stereo receiver. To only inspect eISCP traffic, use this filter:

    tcp.port == 60128 && ip.addr == receiver_ip_addr
