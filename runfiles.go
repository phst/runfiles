// Copyright 2020, 2021 Google LLC
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
//       urls = ["https://github.com/phst/runfiles/archive/60eb03913791bc1cd69c7b2d4316395927c087a8.zip"],
//       sha256 = "84d9bff5719e4d2a3daea5302dd61fed3539675f3211204a31f379db87df9e6a",
//       strip_prefix = "runfiles-60eb03913791bc1cd69c7b2d4316395927c087a8",
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
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

// Runfiles allows access to Bazel runfiles.  Use New to create Runfiles
// objects; the zero Runfiles object always returns errors.  See
// https://docs.bazel.build/skylark/rules.html#runfiles for some information on
// Bazel runfiles.
type Runfiles struct {
	// We don’t need concurrency control since Runfiles objects are
	// immutable once created.
	impl runfiles
	env  string
}

// New creates a given Runfiles object.  By default, it uses os.Args and the
// RUNFILES_MANIFEST_FILE and RUNFILES_DIR environmental variables to find the
// runfiles location.  This can be overwritten by passing some options.
func New(opts ...Option) (*Runfiles, error) {
	var o options
	for _, a := range opts {
		a.apply(&o)
	}
	if o.program == "" {
		o.program = ProgramName(os.Args[0])
	}
	if o.manifest == "" {
		o.manifest = ManifestFile(os.Getenv("RUNFILES_MANIFEST_FILE"))
	}
	if o.directory == "" {
		o.directory = Directory(os.Getenv("RUNFILES_DIR"))
	}
	// See section “Runfiles discovery” in
	// https://docs.google.com/document/d/e/2PACX-1vSDIrFnFvEYhKsCMdGdD40wZRBX3m3aZ5HhVj4CtHPmiXKDCxioTUbYsDydjKtFDAzER5eg7OjJWs3V/pub.
	manifest := ManifestFile(o.program + ".runfiles_manifest")
	if stat, err := os.Stat(string(manifest)); err == nil && stat.Mode().IsRegular() {
		return manifest.new()
	}
	dir := Directory(o.program + ".runfiles")
	if stat, err := os.Stat(string(dir)); err == nil && stat.IsDir() {
		return dir.new(), nil
	}
	if o.manifest != "" {
		return o.manifest.new()
	}
	if o.directory != "" {
		return o.directory.new(), nil
	}
	return nil, errors.New("runfiles: no runfiles found")
}

// Path returns the absolute path name of a runfile.  The runfile name must be
// a relative path, using the slash (not backslash) as directory separator.  If
// r is the zero Runfiles object, Path always returns an error.
func (r *Runfiles) Path(s string) (string, error) {
	// See section “Library interface” in
	// https://docs.google.com/document/d/e/2PACX-1vSDIrFnFvEYhKsCMdGdD40wZRBX3m3aZ5HhVj4CtHPmiXKDCxioTUbYsDydjKtFDAzER5eg7OjJWs3V/pub.
	if s == "" {
		return "", errors.New("runfiles: name may not be empty")
	}
	if path.IsAbs(s) {
		return "", fmt.Errorf("runfiles: name %s may not be absolute", s)
	}
	if s != path.Clean(s) {
		return "", fmt.Errorf("runfiles: name %s must be canonical", s)
	}
	if s == ".." || strings.HasPrefix(s, "../") {
		return "", fmt.Errorf("runfiles: name %s may not contain a parent directory", s)
	}
	impl := r.impl
	if impl == nil {
		return "", errors.New("runfiles: uninitialized Runfiles object")
	}
	p := impl.path(s)
	if p == "" {
		return "", &os.PathError{"stat", s, os.ErrNotExist}
	}
	return p, nil
}

// Env returns additional environmental variables to pass to subprocesses.
// Each element is of the form “key=value”.  Pass these variables to
// Bazel-built binaries so they can find their runfiles as well.  See the
// Runfiles example for an illustration of this.
//
// The return value is a newly-allocated slice; you can modify it at will.  If
// r is the zero Runfiles object, the return value is nil.
func (r *Runfiles) Env() []string {
	if r.env == "" {
		return nil
	}
	return []string{r.env}
}

// Option is an option for the New function to override runfiles discovery.
type Option interface {
	apply(*options)
}

// ProgramName is an Option that sets the program name.  If not set, New uses
// os.Args[0].
type ProgramName string

type options struct {
	program   ProgramName
	manifest  ManifestFile
	directory Directory
}

func (p ProgramName) apply(o *options)  { o.program = p }
func (m ManifestFile) apply(o *options) { o.manifest = m }
func (d Directory) apply(o *options)    { o.directory = d }

type runfiles interface {
	path(string) string
}
