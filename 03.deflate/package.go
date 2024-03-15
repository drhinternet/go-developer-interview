package deflate

// DeflateToBase64 accepts a string as input, runs DEFLATE on it (as implemented
// by the +compress/flate+ standard library, flate.BestCompression), then encodes
// it to Base64 (as implemented by the encoding/base64 standard library, RawStdEncoding).
func DeflateToBase64(input string) string {
	return "implement me"
}
