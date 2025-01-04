# Copyright 2020, 2021, 2022 Google LLC
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

"""Defines an alternative runfiles library for Go."""

workspace(name = "com_github_phst_runfiles")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "0936c9bc3c4321ee372cb8f66dd972d368cb940ed01a9ba9fd7debcf0093f09b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.51.0/rules_go-v0.51.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.51.0/rules_go-v0.51.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    nogo = "@io_bazel_rules_go//:tools_nogo",
    version = "1.19.3",
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "a80893292ae1d78eaeedd50d1cab98f242a17e3d5741b1b9fb58b5fd9d2d57bc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.40.0/bazel-gazelle-v0.40.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.40.0/bazel-gazelle-v0.40.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()
