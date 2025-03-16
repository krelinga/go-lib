package kfscopier_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/krelinga/go-lib/kfs/kfscopier"
	"github.com/krelinga/go-lib/pipe"
)

type tempDirPath string

func (tdp tempDirPath) Join(in string) string {
	return filepath.Join(string(tdp), in)
}

func (tdp tempDirPath) MkdirAll(t *testing.T, in string) {
	t.Helper()
	if err := os.MkdirAll(tdp.Join(in), 0755); err != nil {
		t.Fatal(err)
	}
}

func (tdp tempDirPath) CreateFile(t *testing.T, in string) {
	t.Helper()
	f, err := os.Create(tdp.Join(in))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(in); err != nil {
		t.Fatal(err)
	}
}

func (tdp tempDirPath) Check(t *testing.T, src, dest string) {
	t.Helper()
	data, err := os.ReadFile(tdp.Join(dest))
	if err != nil {
		t.Error(err)
	}
	if string(data) != src {
		t.Errorf("expected %q, got %q", src, string(data))
	}
}

func (tdp tempDirPath) RemoveAll(t *testing.T) {
	t.Helper()
	if err := os.RemoveAll(string(tdp)); err != nil {
		t.Fatal(err)
	}
}

func (tdp tempDirPath) Child(t *testing.T, in string) tempDirPath {
	t.Helper()
	newPath := filepath.Join(string(tdp), in)
	if err := os.MkdirAll(newPath, 0755); err != nil {
		t.Fatal(err)
	}
	return tempDirPath(newPath)
}

func newTempDirPath(t *testing.T) tempDirPath {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", "kfscopier_test")
	if err != nil {
		t.Fatal(err)
	}
	return tempDirPath(tmpDir)
}

func TestCopier(t *testing.T) {
	tests := []struct {
		name string
		opts kfscopier.Options
		init func(t *testing.T, tdp tempDirPath) *kfscopier.Req
		check func(t *testing.T, tdp tempDirPath)
		wantErr error
	}{
		{
			name: "Valid",
			init: func(t *testing.T, tdp tempDirPath) *kfscopier.Req {
				tdp.CreateFile(t, "src")
				return &kfscopier.Req{
					Src:  tdp.Join("src"),
					Dest: tdp.Join("dest"),
				}
			},
			check: func(t *testing.T, tdp tempDirPath) {
				tdp.Check(t, "src", "dest")
			},
		},
		{
			name: "Src does not exist",
			init: func(t *testing.T, tdp tempDirPath) *kfscopier.Req {
				return &kfscopier.Req{
					Src:  tdp.Join("src"),
					Dest: tdp.Join("dest"),
				}
			},
			wantErr: kfscopier.ErrReqSrcNotStat,
		},
		{
			name: "Src is a directory",
			init: func(t *testing.T, tdp tempDirPath) *kfscopier.Req {
				tdp.MkdirAll(t, "src")
				return &kfscopier.Req{
					Src:  tdp.Join("src"),
					Dest: tdp.Join("dest"),
				}
			},
			wantErr: kfscopier.ErrReqSrcNotFile,
		},
		{
			name: "Dest file exists",
			init: func(t *testing.T, tdp tempDirPath) *kfscopier.Req {
				tdp.CreateFile(t, "src")
				tdp.CreateFile(t, "dest")
				return &kfscopier.Req{
					Src:  tdp.Join("src"),
					Dest: tdp.Join("dest"),
				}
			},
			wantErr: kfscopier.ErrReqDestExists,
		},
		{
			name: "Dest file exists and is a directory",
			init: func(t *testing.T, tdp tempDirPath) *kfscopier.Req {
				tdp.CreateFile(t, "src")
				tdp.MkdirAll(t, "dest")
				return &kfscopier.Req{
					Src:  tdp.Join("src"),
					Dest: tdp.Join("dest"),
				}
			},
			wantErr: kfscopier.ErrReqDestExists,
		},
	}
	tdp := newTempDirPath(t)
	defer tdp.RemoveAll(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tdp := tdp.Child(t, tt.name)
			defer tdp.RemoveAll(t)
			req := tt.init(t, tdp)
			in := make(chan *kfscopier.Req, 1)
			in <- req
			close(in)
			errs := kfscopier.New(context.Background(), in, tt.opts)
			gotErrs := []error{}
			pipe.Wait(pipe.ToArrayFunc(errs, &gotErrs))
			switch {
				case tt.wantErr != nil && len(gotErrs) == 0:
					t.Errorf("expected error %v, got none", tt.wantErr)
				case tt.wantErr != nil && len(gotErrs) > 0 && !errors.Is(gotErrs[0], tt.wantErr):
					t.Errorf("expected error %v, got %v", tt.wantErr, gotErrs[0])
				case tt.wantErr == nil && len(gotErrs) > 0:
					t.Errorf("expected no error, got %v", gotErrs[0])
			}
			if len(gotErrs) > 1 {
				t.Errorf("expected at most 1 error, got %d", len(gotErrs))
			}
			if tt.check != nil {
				tt.check(t, tdp)
			}
		})
	}
}