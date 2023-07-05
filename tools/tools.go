package tools

// This creats a package that can be imported using bazel (not the main package). This import will prevent go mod tidy
// from considering this an unused dependency.
