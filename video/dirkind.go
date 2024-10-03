package video

import "os"

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
		if entry.Name() == "tvshow.nfo"{
			return DirKindShow, nil
		}
	}
	return DirKindMovie, nil
}