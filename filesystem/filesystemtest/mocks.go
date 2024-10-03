package filesystemtest

import (
	"io/fs"
	"path/filepath"

	"github.com/krelinga/go-lib/filesystem"
)

func NewMockDirEntry(path string, isDir bool) filesystem.DirEntry {
	return &mockDirEntry{path: path, isDir: isDir}
}

type mockDirEntry struct {
	path  string
	isDir bool
}

func (m *mockDirEntry) Name() string {
	return filepath.Base(m.path)
}

func (m *mockDirEntry) IsDir() bool {
	return m.isDir
}

func (m *mockDirEntry) Type() fs.FileMode {
	if m.isDir {
		return fs.ModeDir
	}
	return 0
}

func (m *mockDirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

func (m *mockDirEntry) Path() string {
	return m.path
}
