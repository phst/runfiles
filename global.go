// Copyright 2020, 2021, 2022 Google LLC
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

package runfiles

import (
	"fmt"
	"path"
	"strings"

	"github.com/bazelbuild/rules_go/go/runfiles"
)

// Path returns the absolute path name of a runfile.  The runfile name must be
// a relative path, using the slash (not backslash) as directory separator.  If
// the runfiles manifest maps s to an empty name (indicating an empty runfile
// not present in the filesystem), Path returns an error that wraps ErrEmpty.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#Rlocation
// instead.
func Path(s string) (string, error) {
	if path.IsAbs(s) {
		return "", fmt.Errorf("runfiles: name %q may not be absolute", s)
	}
	if s == ".." || strings.HasPrefix(s, "../") {
		return "", fmt.Errorf("runfiles: name %q may not contain a parent directory", s)
	}
	return runfiles.Rlocation(s)
}

// Env returns additional environmental variables to pass to subprocesses.
// Each element is of the form “key=value”.  Pass these variables to
// Bazel-built binaries so they can find their runfiles as well.  See the
// Runfiles example for an illustration of this.
//
// The return value is a newly-allocated slice; you can modify it at will.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#Env instead.
func Env() ([]string, error) { return runfiles.Env() }
