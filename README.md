# Git cloner

## Motivation

Every time, when I wanted to clone new project via git, I need to go my GOPATH folder and create sub-folders.


### Example

I need to clone: `git@github.com:vxxvvxxv/go-git-clone.git`

My steps:

1. Go to GOPATH folder, ex.: `cd ~/go/src`
2. Create new folder: `mkdir github.com && cd github.com`
3. Create owner folder: `mkdir vxxvvxxv && cd vxxvvxxv`
4. Clone new project: `git clone git@github.com:vxxvvxxv/go-git-clone.git`

I want just run: `go-git-clone git@github.com:vxxvvxxv/go-git-clone.git`

This app, check and to create folders, and run `git clone <url>`.

## Install

```shell
go install github.com/vxxvvxxv/go-git-clone
```

## Run

```shell
go-git-clone git@github.com:vxxvvxxv/go-git-clone.git
```
