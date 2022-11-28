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

// Package runfiles provides access to Bazel runfiles.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles instead.
package runfiles

import (
	"fmt"
	"path"
	"strings"

	"github.com/bazelbuild/rules_go/go/runfiles"
)

// Runfiles allows access to Bazel runfiles.  Use New to create Runfiles
// objects; the zero Runfiles object always returns errors.  See
// https://docs.bazel.build/skylark/rules.html#runfiles for some information on
// Bazel runfiles.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#Runfiles
// instead.
type Runfiles struct{ runfiles.Runfiles }

// New creates a given Runfiles object.  By default, it uses os.Args and the
// RUNFILES_MANIFEST_FILE and RUNFILES_DIR environmental variables to find the
// runfiles location.  This can be overwritten by passing some options.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#New instead.
func New(opts ...Option) (*Runfiles, error) {
	impl, err := runfiles.New(opts...)
	if err != nil {
		return nil, err
	}
	return &Runfiles{*impl}, err
}

// Path returns the absolute path name of a runfile.  The runfile name must be
// a relative path, using the slash (not backslash) as directory separator.  If
// r is the zero Runfiles object, Path always returns an error.  If the
// runfiles manifest maps s to an empty name (indicating an empty runfile not
// present in the filesystem), Path returns an error that wraps ErrEmpty.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#Runfiles.Rlocation
// instead.
func (r *Runfiles) Path(s string) (string, error) {
	if path.IsAbs(s) {
		return "", fmt.Errorf("runfiles: name %q may not be absolute", s)
	}
	if s == ".." || strings.HasPrefix(s, "../") {
		return "", fmt.Errorf("runfiles: name %q may not contain a parent directory", s)
	}
	return r.Rlocation(s)
}

// Option is an option for the New function to override runfiles discovery.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#Option
// instead.
type Option = runfiles.Option

// ProgramName is an Option that sets the program name.  If not set, New uses
// os.Args[0].
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#ProgramName
// instead.
type ProgramName = runfiles.ProgramName

// Error represents a failure to look up a runfile.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#Error instead.
type Error = runfiles.Error

// ErrEmpty indicates that a runfile isn’t present in the filesystem, but
// should be created as an empty file if necessary.
//
// Deprecated: use
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/runfiles#ErrEmpty
// instead.
var ErrEmpty = runfiles.ErrEmpty
