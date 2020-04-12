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

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "runfiles",
    srcs = glob(
        ["*.go"],
        exclude = [
            "*_test.go",
            "testprog.go",
        ],
    ),
    importpath = "github.com/phst/runfiles",
    visibility = ["//visibility:public"],
)

go_test(
    name = "runfiles_test",
    srcs = glob(["*_test.go"]),
    data = [
        "test.txt",
        "testprog",
    ],
    deps = [":runfiles"],
)

go_binary(
    name = "testprog",
    testonly = 1,
    srcs = ["testprog.go"],
    out = "testprog",
    data = ["test.txt"],
    deps = [":runfiles"],
)
