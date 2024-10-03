package video

import (
	"context"
	"path/filepath"
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
	type record = pipe.KV[string, state]
	type groupedRecord = pipe.KV[string, []state]

	filterTypes := func(entry filesystem.DirEntry) bool {
		ext := filepath.Ext(entry.Path())
		return ext == ".mkv" || ext == ".nfo" || ext == ".tcprofile"
	}
	toRecord := func(entry filesystem.DirEntry) *record {
		switch {
		case strings.HasSuffix(entry.Path(), ".mkv"):
			return &record{Key: entry.Path(), Val: foundMkv}
		case strings.HasSuffix(entry.Path(), ".nfo"):
			return &record{Key: NewPathsFromNfo(entry.Path()).Mkv(), Val: foundNfo}
		case strings.HasSuffix(entry.Path(), ".tcprofile"):
			return &record{Key: NewPathsFromTcProfile(entry.Path()).Mkv(), Val: foundTcProfile}
		default:
			panic("unexpected file type")
		}
	}

	groupedToFileInfo := func(in *groupedRecord) *FileInfo {
		out := &FileInfo{}
		for _, v := range in.Val {
			switch v {
			case foundMkv:
				out.MkvPath = in.Key
			case foundNfo:
				out.HasNfo = true
			case foundTcProfile:
				out.HasTcProfile = true
			}
		}
		return out
	}

	filterEmptyMkvPath := func(f *FileInfo) bool {
		return f.MkvPath != ""
	}

	filtered := pipe.ParDoFilter(ctx, 1, files, filterTypes)
	records := pipe.ParDo(ctx, 1, filtered, toRecord)
	grouped := pipe.GroupBy(ctx, records)
	fileInfos := pipe.ParDo(ctx, 1, grouped, groupedToFileInfo)
	return pipe.ParDoFilter(ctx, 1, fileInfos, filterEmptyMkvPath)
}
