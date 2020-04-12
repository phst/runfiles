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

package runfiles_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/phst/runfiles"
)

func ExamplePath() {
	path, err := runfiles.Path("phst_runfiles/runfiles/test.txt")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// Output: hi!
}

func ExampleRunfiles() {
	r, err := runfiles.New()
	if err != nil {
		panic(err)
	}
	// The binary “testprog” is itself built with Bazel, and needs
	// runfiles.
	prog, err := r.Path("phst_runfiles/runfiles/testprog")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(prog)
	cmd.Env = append(os.Environ(), r.Env()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	// Output: hi!
}
