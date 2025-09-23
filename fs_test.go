// Copyright 2021, 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runfiles_test

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/phst/runfiles"
)

func TestFS(t *testing.T) {
	fsys, err := runfiles.New()
	if err != nil {
		t.Fatal(err)
	}
	// Ensure that the Runfiles object implements FS interfaces.
	var _ fs.FS = fsys
	var _ fs.StatFS = fsys
	var _ fs.ReadFileFS = fsys
	if err := fstest.TestFS(fsys, "_main/test.txt", "_main/testprog/testprog"); err != nil {
		t.Error(err)
	}
}

func TestFS_empty(t *testing.T) {
	dir := t.TempDir()
	manifest := filepath.Join(dir, "manifest")
	if err := os.WriteFile(manifest, []byte("__init__.py \n"), 0600); err != nil {
		t.Fatal(err)
	}
	fsys, err := runfiles.New(runfiles.ManifestFile(manifest), runfiles.ProgramName("/invalid"), runfiles.Directory("/invalid"))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Open", func(t *testing.T) {
		fd, err := fsys.Open("__init__.py")
		if err != nil {
			t.Fatal(err)
		}
		defer fd.Close()
		got, err := io.ReadAll(fd)
		if err != nil {
			t.Error(err)
		}
		if len(got) != 0 {
			t.Errorf("got nonempty contents: %q", got)
		}
	})
	t.Run("Stat", func(t *testing.T) {
		got, err := fsys.Stat("__init__.py")
		if err != nil {
			t.Fatal(err)
		}
		if got.Name() != "__init__.py" {
			t.Errorf("Name: got %q, want %q", got.Name(), "__init__.py")
		}
		if got.Size() != 0 {
			t.Errorf("Size: got %d, want %d", got.Size(), 0)
		}
		if !got.Mode().IsRegular() {
			t.Errorf("IsRegular: got %v, want %v", got.Mode().IsRegular(), true)
		}
		if got.IsDir() {
			t.Errorf("IsDir: got %v, want %v", got.IsDir(), false)
		}
	})
	t.Run("ReadFile", func(t *testing.T) {
		got, err := fsys.ReadFile("__init__.py")
		if err != nil {
			t.Error(err)
		}
		if len(got) != 0 {
			t.Errorf("got nonempty contents: %q", got)
		}
	})
}
