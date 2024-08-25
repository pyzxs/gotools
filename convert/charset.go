package convert

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
)

// convertCharset performs charset conversion between two encodings.
func convertCharset(input string, from encoding.Encoding, to encoding.Encoding) (string, error) {
	// Create a reader that will decode the input string from 'from' encoding to UTF-8
	decoder := transform.NewReader(bytes.NewReader([]byte(input)), from.NewDecoder())

	// Create a writer that will encode the input to 'to' encoding
	var buf bytes.Buffer
	encoder := transform.NewWriter(&buf, to.NewEncoder())

	// Copy from decoder to encoder
	if _, err := io.Copy(encoder, decoder); err != nil {
		return "", err
	}

	// Flush the encoder to ensure all data is written
	if err := encoder.Close(); err != nil {
		return "", err
	}

	return buf.String(), nil
}
