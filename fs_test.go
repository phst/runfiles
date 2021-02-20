// Copyright 2021 Google LLC
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

// +build go1.16

package runfiles_test

import (
	"io/fs"
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
	if err := fstest.TestFS(fsys, "com_github_phst_runfiles/test.txt", "com_github_phst_runfiles/testprog/testprog"); err != nil {
		t.Error(err)
	}
}
