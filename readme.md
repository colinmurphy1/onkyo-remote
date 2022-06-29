# Onkyo Remote

By Colin Murphy, 2022

Onkyo-remote exposes the eISCP protocol Onkyo/Integra/newer Pioneer home audio
equipment uses into a web-based remote and REST API. It utilizes a self-built
library that sends and receives eISCP commands from a receiver using
documentation from Onkyo. 

The frontend is written using Svelte.js, and the backend code is written in Go.

## Requirements

Please ensure you have the following software installed on your computer to
compile onkyo-remote:

* Go ^1.18
* Node ^17.x
* GNU Make

If you do not want to have all these dependencies installed, you may download
a precompiled release once I release a stable version of the software.

## Usage

Make a copy of the `config.example.yaml` file and change the IP address to
the IPv4 address of your Onkyo/Integra receiver. You can also make other
setting changes, for example to disable the eISCP logging and web interface. 

Once done, compile the software:

    make build

This will compile the Svelte.js frontend and then build three executables. When
it finishes compiling everything, you may launch onkyo-remote by running
this command:

    # Linux
    ./onkyo-remote -config ./config.yaml

    # macOS
    ./onkyo-remote-darwin -config ./config.yaml

    # Windows
    ./onkyo-remote-windows.exe -config ./config.yaml

Open a web browser and navigate to http://localhost:8080/api/status. You should
get a response returning the general status of the receiver.

For more information, please see the [Wiki][0].

### Running multiple instances

Onkyo-remote only supports controlling a single receiver at a time by design.
If you need to control multiple Onkyo/Integra receivers, you will need to have
multiple instances of the software running, each with their own configuration
files.

## Debugging

I use Wireshark for debugging the network activity between this application
and my stereo receiver. To only inspect eISCP traffic, use this filter:

    tcp.port == 60128 && ip.addr == receiver_ip_addr

[0]: https://github.com/colinmurphy1/onkyo-remote/wiki
