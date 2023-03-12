package anthropic

import (
	"bufio"
	"bytes"
	"io"

	"github.com/sourcegraph/sourcegraph/lib/errors"
)

const maxPayloadSize = 10 * 1024 * 1024 // 10mb

// Decoder decodes streaming events from a Server Sent Event stream. It only supports
// streams generated by the Anthropic completions API. IE this is not a fully
// compliant Server Sent Events decoder.
//
// Adapted from internal/search/streaming/http/decoder.go.
type Decoder struct {
	scanner *bufio.Scanner
	data    []byte
	err     error
}

func NewDecoder(r io.Reader) *Decoder {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 4096), maxPayloadSize)
	// bufio.ScanLines, except we look for \r\n\r\n which separate events.
	split := func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte("\r\n\r\n")); i >= 0 {
			return i + 4, data[:i], nil
		}
		// If we're at EOF, we have a final, non-terminated event. This should
		// be empty.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	}
	scanner.Split(split)
	return &Decoder{
		scanner: scanner,
	}
}

// Scan advances the decoder to the next event in the stream. It returns
// false when it either hits the end of the stream or an error.
func (d *Decoder) Scan() bool {
	if !d.scanner.Scan() {
		d.err = d.scanner.Err()
		return false
	}

	// data: json($data)|[DONE]
	data := d.scanner.Bytes()
	dataK, data := splitColon(data)
	if !bytes.Equal(dataK, []byte("data")) {
		d.err = errors.Errorf("malformed data, expected data: %s", dataK)
		return false
	}

	d.data = data
	return true
}

// Event returns the event data of the last decoded event
func (d *Decoder) Data() []byte {
	return d.data
}

// Err returns the last encountered error
func (d *Decoder) Err() error {
	return d.err
}

func splitColon(data []byte) ([]byte, []byte) {
	i := bytes.Index(data, []byte(":"))
	if i < 0 {
		return bytes.TrimSpace(data), nil
	}
	return bytes.TrimSpace(data[:i]), bytes.TrimSpace(data[i+1:])
}
