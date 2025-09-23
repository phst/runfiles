// Copyright 2021, 2022, 2025 Google LLC
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

import "io/fs"

// Open implements fs.FS.Open.
func (r *Runfiles) Open(name string) (fs.File, error) {
	return r.Runfiles.Open(name)
}

// Stat implements fs.StatFS.Stat.
func (r *Runfiles) Stat(name string) (fs.FileInfo, error) {
	return fs.Stat(&r.Runfiles, name)
}

// ReadFile implements fs.ReadFileFS.ReadFile.
func (r *Runfiles) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(&r.Runfiles, name)
}
