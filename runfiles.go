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

// Package runfiles provides access to Bazel runfiles.  It is an alternative to
// the official Bazel package
// https://pkg.go.dev/github.com/bazelbuild/rules_go/go/tools/bazel.
//
// Installation
//
// To use this package, first set up rules_go as described in
// https://github.com/bazelbuild/rules_go#setup.  Then add the following
// snippet to your Bazel WORKSPACE file:
//
//   load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
//   http_archive(
//       name = "com_github_phst_runfiles",
//       urls = ["https://github.com/phst/runfiles/archive/f8065aa0cb28b5cc0fffa7d0b5e9ea1a92add4bb.zip"],
//       sha256 = "8f0502d14cc35e8857d67ac02a3b8d46a496bb7e3ddf723ebb73ce71c4c0cd6d",
//       strip_prefix = "runfiles-f8065aa0cb28b5cc0fffa7d0b5e9ea1a92add4bb",
//   )
//
// Usage
//
// This package has two main entry points, the global functions Path and Env,
// and the Runfiles type.
//
// Global functions
//
// For simple use cases that don’t require hermetic behavior, use the Path and
// Env functions to access runfiles.  Use Path to find the filesystem location
// of a runfile, and use Env to obtain environmental variables to pass on to
// subprocesses.
//
// Runfiles type
//
// If you need hermetic behavior or want to change the runfiles discovery
// process, use New to create a Runfiles object.  New accepts a few options to
// change the discovery process.  Runfiles objects have methods Path and Env,
// which correspond to the package-level functions.  On Go 1.16, *Runfiles
// implements fs.FS, fs.StatFS, and fs.ReadFileFS.
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
