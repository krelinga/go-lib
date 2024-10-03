package video

import (
	"context"
	"strings"

	"github.com/krelinga/go-lib/filesystem"
	"github.com/krelinga/go-lib/pipe"
)

// FileInfo contains information about .mkv file that exists on-disk, as well as
// the presence of related .nfo and .tcprofile files.
type FileInfo struct {
	MkvPath      string
	HasNfo       bool
	HasTcProfile bool
}

// Paths is used to convert between different file types.  See the NewPathsFrom* functions.
type Paths struct {
	mkv       string
	nfo       string
	tcProfile string
}

func (p *Paths) Mkv() string {
	return p.mkv
}

func (p *Paths) Nfo() string {
	return p.nfo
}

func (p *Paths) TcProfile() string {
	return p.tcProfile
}

// Returns an instance of the Paths type based on the given .mkv file path.
func NewPathsFromMkv(mkv string) *Paths {
	base, found := strings.CutSuffix(mkv, ".mkv")
	if !found {
		panic("expected path to end with .mkv")
	}
	return &Paths{mkv: mkv,
		nfo:       base + ".nfo",
		tcProfile: base + ".tcprofile"}
}

// Returns an instance of the Paths type based on the given .nfo file path.
func NewPathsFromNfo(nfo string) *Paths {
	base, found := strings.CutSuffix(nfo, ".nfo")
	if !found {
		panic("expected path to end with .nfo")
	}
	return &Paths{
		mkv:       base + ".mkv",
		nfo:       nfo,
		tcProfile: base + ".tcprofile",
	}
}

// Returns an instance of the Paths type based on the given .tcprofile file path.
func NewPathsFromTcProfile(tcProfile string) *Paths {
	base, found := strings.CutSuffix(tcProfile, ".tcprofile")
	if !found {
		panic("expected path to end with .tcprofile")
	}
	return &Paths{
		mkv:       base + ".mkv",
		nfo:       base + ".nfo",
		tcProfile: tcProfile,
	}
}

// BuildFileInfo reads from the given channel of filesystem.DirEntry objects, groups presence
// info for .mkv, .nfo, and .tcprofile files together, and produces a channel of FileInfo objects.
func BuildFileInfo(ctx context.Context, files <-chan filesystem.DirEntry) <-chan *FileInfo {
	type state int
	const (
		foundMkv state = iota
		foundNfo
		foundTcProfile
	)

	toRecords := func(in <-chan filesystem.DirEntry) <-chan *pipe.KV[string, state] {
		out := make(chan *pipe.KV[string, state])

		go func() {
			defer close(out)

			for entry := range in {
				if strings.HasSuffix(entry.Path(), ".mkv") {
					if !pipe.TryWrite(ctx, out, &pipe.KV[string, state]{Key: entry.Path(), Val: foundMkv}) {
						return
					}
				}
				if strings.HasSuffix(entry.Path(), ".nfo") {
					if !pipe.TryWrite(ctx, out, &pipe.KV[string, state]{Key: NewPathsFromNfo(entry.Path()).Mkv(), Val: foundNfo}) {
						return
					}
				}
				if strings.HasSuffix(entry.Path(), ".tcprofile") {
					if !pipe.TryWrite(ctx, out, &pipe.KV[string, state]{Key: NewPathsFromTcProfile(entry.Path()).Mkv(), Val: foundTcProfile}) {
						return
					}
				}
			}
		}()

		return out
	}

	groupedToFileInfo := func(in <-chan *pipe.KV[string, []state]) <-chan *FileInfo {
		out := make(chan *FileInfo)

		go func() {
			defer close(out)

			for group := range in {
				found := func(s state) bool {
					for _, v := range group.Val {
						if v == s {
							return true
						}
					}
					return false
				}
				if !found(foundMkv) {
					continue
				}
				info := &FileInfo{
					MkvPath:      group.Key,
					HasNfo:       found(foundNfo),
					HasTcProfile: found(foundTcProfile),
				}
				if !pipe.TryWrite(ctx, out, info) {
					return
				}
			}
		}()

		return out
	}

	records := toRecords(files)
	grouped := pipe.GroupBy(ctx, records)
	return groupedToFileInfo(grouped)
}
