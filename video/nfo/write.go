package nfo

import "io"

func WriteTo(nfo Nfo, out io.Writer) error {
	_, err := nfo.getDocument().WriteTo(out)
	return err
}
