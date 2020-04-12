// Copyright 2020 Google LLC
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

import "sync"

// Path returns the absolute path name of a runfile.  The runfile name must be
// a relative path, using the slash (not backslash) as directory separator.
func Path(s string) (string, error) {
	r, err := g.get()
	if err != nil {
		return "", err
	}
	return r.Path(s)
}

// Env returns additional environmental variables to pass to subprocesses.
// Each element is of the form “key=value”.  Pass these variables to
// Bazel-built binaries so they can find their runfiles as well.
//
// The return value is a newly-allocated slice; you can modify it at will.
func Env() ([]string, error) {
	r, err := g.get()
	if err != nil {
		return nil, err
	}
	return r.Env(), nil
}

type global struct {
	mu sync.Mutex
	r  *Runfiles
}

func (g *global) get() (*Runfiles, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.r != nil {
		return g.r, nil
	}
	r, err := New()
	if err != nil {
		return nil, err
	}
	g.r = r
	return r, nil
}

var g global
