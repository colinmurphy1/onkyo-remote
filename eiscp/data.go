package eiscp

import "net"

// Struct that stores the connection to the receiver
type Connection struct {
	ip          string      // IP address of the Onkyo receiver
	con         net.Conn    // Connection is stored here
	iscpVersion byte        // ISCP version (default 0x1) (should not need changed)
	iscpDest    byte        // ISCP destination (default 0x31)
	Status      OnkyoStatus // Store status of receiver
}

// Struct that stores the general status of the receiver
type OnkyoStatus struct {
	Power    Power    // Power status
	Volume   Volume   // Volume status
	SongInfo SongInfo // Song information
}

type Power struct {
	Status bool // Power status
}

// Volume status
type Volume struct {
	Level int  // Volume level
	Mute  bool // Mute status
}

// Song information
type SongInfo struct {
	Title  string    // Song title
	Artist string    // Song Artist
	Album  string    // Song Album
	Time   SongTime  // Song time/position
	Track  SongTrack // Track position
}

// Song time position/length
type SongTime struct {
	Current string // Position in HH:MM:SS
	Length  string // Length in HH:MM:SS
}

// Song track position
type SongTrack struct {
	Current int
	Total   int
}
