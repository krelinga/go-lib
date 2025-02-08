package nfo

import "io"

// WriteTo writes the NFO document to the given writer.
func WriteTo(nfo Nfo, out io.Writer) error {
	_, err := nfo.getDocument().WriteTo(out)
	return err
}
