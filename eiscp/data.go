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
	Input    Input    // Input source
	Volume   Volume   // Volume status
	SongInfo SongInfo // Song information
	Tuner    Tuner    // Tuner status
}

type Power struct {
	Status bool // Power status
}

type Input struct {
	Name    string // Input name
	HexCode string // Input HEX code (for debugging, and future custom naming)
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

// Tuner status
type Tuner struct {
	Frequency float64 // Tuner frequency
	Preset    int     // Tuner preset
}

// Input names
var Inputs = map[string]string{
	"00": "VCR/DVR",
	"01": "CBL/SAT",
	"02": "GAME",
	"03": "AUX",
	"05": "PC",
	"10": "BD/DVD",
	"11": "STRM BOX",
	"12": "TV",
	"20": "TV/TAPE",
	"22": "PHONO",
	"23": "CD",
	"24": "FM",
	"25": "AM",
	"26": "TUNER",
	"27": "MUSIC SERVER",
	"28": "INTERNET RADIO",
	"29": "USB",
	"2B": "NETWORK",
	"2C": "USB",
	"2D": "AIRPLAY",
	"2E": "BLUETOOTH",
	"40": "UNIVERSAL PORT",
}
