package lib

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// Get album art from receiver's HTTP server. Returns a slice of bytes and an
// error if there is one.
func GetArt(url string) ([]byte, error) {
	// Make HTTP request for album art
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Close connection

	// We want a 200 OK from the server
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad response code")
	}

	// If the content length is 0, there is no album art available
	if resp.ContentLength == 0 {
		return nil, errors.New("no album art")
	}

	// Turn the io.Reader into raw bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Split the bytes using the newline character
	bodySplit := bytes.Split(body, []byte("\n"))

	// Remove the first 2 lines of the response, which removes the
	// Content-Type and Content-length headers which are part of the body
	// due to the Onkyo web server being buggy.
	image := bytes.Join(bodySplit[3:], []byte("\n"))

	return image, nil
}
