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
	Name      string // Input name
	HexCode   string // Input HEX code (for debugging, and future custom naming)
	NetSource string // NET Source (DLNA, AirPlay, Spotify, etc.) Leave blank if not in NET
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
	"31": "XM",
	"32": "SIRIUS",
	"33": "DAB",
	"2B": "NETWORK",
	"2C": "USB",
	"2D": "AIRPLAY",
	"2E": "BLUETOOTH",
	"40": "UNIVERSAL PORT",
}

// NET Services
var NetServices = map[string]string{
	"00": "DLNA",
	"01": "My Favorite",
	"02": "vTuner",
	"03": "SiriusXM",
	"04": "Pandora",
	"05": "Rhapsody",
	"06": "Last.fm",
	"07": "Slacker",
	"0A": "Spotify",
	"0B": "AUPEO!",
	"0C": "radiko",
	"0D": "e-onkyo",
	"0E": "TuneIn",
	"0F": "MP3tunes",
	"10": "Simfy",
	"11": "Home Media",
	"12": "Deezer",
	"13": "iHeartRadio",
	"18": "Airplay",
	"F0": "Front USB",
	"F1": "Rear USB",
	"F2": "Internet Radio",
	"F3": "NET",
	"F4": "Bluetooth",
}
