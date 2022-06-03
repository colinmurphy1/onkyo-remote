package main

import (
	"bytes"
	"encoding/binary"
)

type OnkyoCommand struct {
	Version     byte
	Destination byte
	Command     []byte
}

// Generates an ISCP command, to be placed into an eISCP packet
func (cmd *OnkyoCommand) IscpCommand() []byte {
	buf := bytes.Buffer{}
	buf.WriteRune('!')
	buf.WriteByte(cmd.Destination) // Most times this will be 0x31 = 1
	buf.Write(cmd.Command)
	buf.Write([]byte{0x0D}) // A command ends with a carriage return
	return buf.Bytes()
}

// Generate an eISCP packet
func (cmd *OnkyoCommand) EiscpCommand() []byte {

	// Iscp command
	command := cmd.IscpCommand()

	// Make 4 bytes
	sizebuf := make([]byte, 4)

	buf := bytes.Buffer{}
	buf.Write([]byte("ISCP"))                                 // Start of message
	buf.Write([]byte{0, 0, 0, 0x10})                          // Header size
	binary.BigEndian.PutUint32(sizebuf, uint32(len(command))) // Length of the ISCP command
	buf.Write(sizebuf)                                        // Data size
	buf.WriteByte(cmd.Version)                                // Version
	buf.Write([]byte{0, 0, 0})                                // Reserved
	buf.Write(command)                                        // ISCP Command
	return buf.Bytes()
}
