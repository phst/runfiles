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

package runfiles

import (
	"io/fs"
	"os"
)

// Open implements fs.FS.Open.
func (r *Runfiles) Open(name string) (fs.File, error) {
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{"open", name, fs.ErrInvalid}
	}
	p, err := r.Path(name)
	if err != nil {
		return nil, &fs.PathError{"open", name, err}
	}
	return os.Open(p)
}

// Stat implements fs.StatFS.Stat.
func (r *Runfiles) Stat(name string) (fs.FileInfo, error) {
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{"stat", name, fs.ErrInvalid}
	}
	p, err := r.Path(name)
	if err != nil {
		return nil, &fs.PathError{"stat", name, err}
	}
	return os.Stat(p)
}

// ReadFile implements fs.ReadFileFS.ReadFile.
func (r *Runfiles) ReadFile(name string) ([]byte, error) {
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{"open", name, fs.ErrInvalid}
	}
	p, err := r.Path(name)
	if err != nil {
		return nil, &fs.PathError{"open", name, err}
	}
	return os.ReadFile(p)
}
