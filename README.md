# Myfind

## Contents
<h3 id="ex00">Exercise 00: Finding Things</h3>
As a first step,
As a first step, let's implement a `find`-like utility using Go. It needs to take a path and a set of command line options to be able to find different types of entries. We are interested in three types of entries, which are directories, regular files and symbolic links. So we should be able to run our program as follows:

```
# Find all files/directories/symlinks recursively in directory /foo
~$ ./myFind /foo
/foo/bar
/foo/bar/baz
/foo/bar/baz/deep/directory
/foo/bar/test.txt
/foo/bar/buzz -> /foo/bar/baz
/foo/bar/broken_sl -> [broken]
```

or specify `-sl`, `-d` or `-f` to print only symlinks, only directories or only files. Note that the user should be able to explicitly specify one, two or all three of these options, e.g. `./myFind -f -sl /path/to/dir` or `./myFind -d /path/to/other/dir`.

You should also implement another option — `-ext` (works ONLY if -f is specified) to allow users to print only files with a certain extension. An extension in this task can be thought of as the last part of the filename when we split it by a dot. So,

```
# Find only *.go files and ignore all others.
~$ ./myFind -f -ext 'go' /go
/go/src/github.com/mycoolproject/main.go
/go/src/github.com/mycoolproject/magic.go
```

You'll also need to resolve symlinks. So if `/foo/bar/buzz` is a symlink pointing to another place in FS, like `/foo/bar/baz`, print both paths separated by `->`, like in the example above. 

Another thing about symlinks is that they can be broken (pointing to a non-existent file node). In this case, your code should print `[broken]` instead of the path of a symlink destination.

Files and directories that the current user doesn't have access to (permission errors) should be skipped in the output and not cause a runtime error.