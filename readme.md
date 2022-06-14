# Onkyo Remote

By Colin Murphy, 2022

This application exposes the eISCP protocol Onkyo home audio equipment uses into
a REST API. It utilizes a self-built library that communicates with a receiver.

For a work in progress API reference, see `api.md`.

## Usage

Set the receiver's IP address, and optionally the port the web server listens
on.

    export RECEIVER_IP="192.168.1.180"
    export HTTP_PORT="8080"

Start the service...

    go run main.go


Open a web browser and navigate to http://localhost:8080/api/power. You should
get a response returning the power status of the receiver.
