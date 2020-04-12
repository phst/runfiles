# Alternative runfiles library for Go

This package contains an alternative runfiles library for
[`rules_go`][rules_go].  This library follows the [runfiles libraries
proposal][] more closely.  In particular, it requires specifying a workspace
name for runfile paths, and handles the case where no runfiles-specific
environmental variables are set correctly.  You can use both this library and
the [official Bazel library][] in the same program or library; they don’t
interfere with each other.

See the [Go package documentation][] for instructions how to use this package.

[rules_go]: https://github.com/bazelbuild/rules_go
[runfiles libraries proposal]: https://docs.google.com/document/d/e/2PACX-1vSDIrFnFvEYhKsCMdGdD40wZRBX3m3aZ5HhVj4CtHPmiXKDCxioTUbYsDydjKtFDAzER5eg7OjJWs3V/pub
[official Bazel library]: https://pkg.go.dev/github.com/bazelbuild/rules_go/go/tools/bazel?tab=doc
[Go package documentation]: https://pkg.go.dev/github.com/phst/runfiles?tab=doc
