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

load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "directory.go",
        "global.go",
        "manifest.go",
        "runfiles.go",
    ],
    importpath = "github.com/phst/runfiles",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["runfiles_test.go"],
    data = [
        "test.txt",
        "//testprog",
    ],
    embed = [":go_default_library"],
    rundir = ".",
)

exports_files(
    ["test.txt"],
    visibility = ["//testprog:__pkg__"],
)

# gazelle:prefix github.com/phst/runfiles
gazelle(name = "gazelle")
