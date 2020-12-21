load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_antlr_antlr4",
        importpath = "github.com/antlr/antlr4",
        sum = "h1:uA+3jkHfX7cieBgh5GBV++KZQbGdBz0VHJBWQTFvAss=",
        version = "v0.0.0-20201214011320-c79b0fd80c9d",
    )
