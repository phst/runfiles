#!/bin/bash

# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

shopt -s -o errexit noclobber noglob nounset pipefail
shopt -u -o braceexpand history
shopt -s failglob

IFS=''

sha256() {
  wget --no-verbose --output-document=- -- "${url:?}" \
    | sha256sum \
    | cut --fields=1 --delimiter=' '
}

shopt -s -o xtrace

rev="$(git rev-parse remotes/github/master)"
url="https://github.com/phst/runfiles/archive/${rev:?}.zip"
hash="$(sha256 "${url:?}")"

sed --sandbox --in-place --regexp-extended \
  --expression="s|^(// +urls = \\[\"https://github.com/phst/runfiles/archive/)([[:xdigit:]]*)(\\.zip\"\\],)\$|\1${rev:?}\3|" \
  --expression="s|^(// +strip_prefix = \"runfiles-)([[:xdigit:]]*)(\",)\$|\1${rev:?}\3|" \
  --expression="s|^(// +sha256 = \")([[:xdigit:]]*)(\",)\$|\1${hash:?}\3|" \
  -- runfiles/runfiles.go
