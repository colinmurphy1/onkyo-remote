package eiscp

import (
	"net"
)

// Struct that stores the connection to the receiver
type Connection struct {
	ip          string            // IP address of the Onkyo receiver
	port        int               // ISCP port of the receiver (default 60128)
	con         net.Conn          // Connection is stored here
	iscpVersion byte              // ISCP version (default 0x1) (should not need changed)
	iscpDest    byte              // ISCP destination (default 0x31)
	Status      OnkyoStatus       // Store state of receiver
	AlbumArt    albumArt          // Album art
	XmlData     string            // XML data for debugging
	Inputs      map[string]string // List of inputs
}

type receiverInfo struct {
	Brand        string // Brand (Onkyo, Integra, Pioneer?)
	ModelName    string // Model number
	FriendlyName string // Friendly name (user-definable at web interface)
}

// Struct that stores the current state of the receiver
type OnkyoStatus struct {
	Power    power        // Power status
	Input    input        // Input source
	Volume   volume       // Volume status
	SongInfo songInfo     // Song information
	Tuner    tuner        // Tuner status
	Info     receiverInfo // Information about the receiver
}

// Power status
type power struct {
	Status bool
}

// Current input information
type input struct {
	Name      string    // Input name
	HexCode   string    // Input HEX code (for debugging, and future custom naming)
	NetSource string    // NET Source (DLNA, AirPlay, Spotify, etc.) Leave blank if not in NET
	Info      inputInfo // Input source information
}

// Source information
type inputInfo struct {
	InputPort           string
	InputFormat         string
	SamplingFreq        string
	InputSignalChannel  string
	ListenMode          string
	OutputSignalChannel string
}

// Volume status
type volume struct {
	Level int  // Volume level
	Mute  bool // Mute status
	Max   uint // Maximum volume
}

// Song information
type songInfo struct {
	Title    string    // Song title
	Artist   string    // Song Artist
	Album    string    // Song Album
	AlbumArt bool      // Album art available
	Time     songTime  // Song time/position
	Track    songTrack // Track position
	Status   string    // Play pause etc.
}

// Song time position/length
type songTime struct {
	Current string // Position in HH:MM:SS
	Length  string // Length in HH:MM:SS
}

// Song track position
type songTrack struct {
	Current int
	Total   int
}

// Tuner status
type tuner struct {
	Frequency  int                    // Tuner frequency
	Preset     int64                  // Tuner preset
	PresetList map[string]tunerPreset // presets
}

type tunerPreset struct {
	Frequency string // Tuner frequency
	Band      int    // Band (0: No preset, 1: FM, 2: AM)
}

// =============================================================================
// NET Services
var netServices = map[string]string{
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
	"18": "AirPlay",
	"F0": "Front USB",
	"F1": "Rear USB",
	"F2": "Internet Radio",
	"F3": "NET",
	"F4": "Bluetooth",
}

// =============================================================================
// ALBUM ART
type albumArt struct {
	Data        []byte // Binary image data
	ContentType string // Content type (eg. image/jpeg)
}
