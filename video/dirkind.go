package video

import (
	"encoding/json"
	"fmt"
	"os"
)

// DirKind is an enum used to represent the type of video files in a given directory: either a movie or a tv show.
// DirKind implements the json.Marshaler and json.Unmarshaler interfaces, so human-readable names will be used when serializing to JSON.
type DirKind int

const (
	DirKindMovie DirKind = iota
	DirKindShow
)

// Examines metadata under dir to determine if it is a movie or a TV show.
func GetDirKind(dir string) (DirKind, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}
	for _, entry := range entries {
		if entry.Name() == "tvshow.nfo" {
			return DirKindShow, nil
		}
	}
	return DirKindMovie, nil
}

func (dk DirKind) MarshalJSON() ([]byte, error) {
	var kind string
	switch dk {
	case DirKindMovie:
		kind = "movie"
	case DirKindShow:
		kind = "show"
	default:
		return nil, fmt.Errorf("unknown DirKind: %d", dk)
	}
	return json.Marshal(kind)
}

func (dk *DirKind) UnmarshalJSON(data []byte) error {
	var kind string
	if err := json.Unmarshal(data, &kind); err != nil {
		return err
	}
	switch kind {
	case "movie":
		*dk = DirKindMovie
	case "show":
		*dk = DirKindShow
	default:
		return fmt.Errorf("unknown DirKind: %s", kind)
	}
	return nil
}
