package deflate_test

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
	"testing"

	"deflate"
)

func TestDeflateToBase64(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"", "AQAA//8"},
		{"x", "qgAEAAD//w"},
		{"hello world", "ykjNyclXKM8vykkBBAAA//8"},
		{"This is a longer text, but still not very long.", "CsnILFbILFZIVMjJz0tPLVIoSa0o0VFIKi1RKC7JzMlRyMsvUShLLaoEy+sBAgAA//8"},
	}

	for _, test := range tests {
		output := deflate.DeflateToBase64(test.input)

		if output != test.expect {
			t.Errorf("expected %v, got %v", test.expect, output)
			continue
		}

		compressed, err := base64.RawStdEncoding.DecodeString(output)

		if err != nil {
			t.Errorf("error base64 decoding: %v", err)
			continue
		}

		decompressed, err := io.ReadAll(flate.NewReader(bytes.NewBuffer(compressed)))

		if err != nil {
			t.Errorf("error decompressing: %v", err)
			continue
		}

		if string(decompressed) != test.input {
			t.Errorf("expected decompression %v, got %s", test.input, decompressed)
			continue
		}
	}
}
