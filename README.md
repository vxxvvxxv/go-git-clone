# Git cloner

## Motivation

Every time, when I wanted to clone new project via git, I need to go my GOPATH folder and create sub-folders, for example:

1. I need to clone: `git@github.com:vxxvvxxv/go-git-clone.git`
2. Go to GOPATH folder, ex.: `cd ~/go/src`
3. Create new folder: `mkdir github.com && cd github.com`
4. Create owner folder: `mkdir vxxvvxxv && cd vxxvvxxv`
5. Clone new project: `git clone git@github.com:vxxvvxxv/go-git-clone.git && go-git-clone`

I want just run: `go-git-clone git@github.com:vxxvvxxv/go-git-clone.git`

This app, check and to create folders, run `git clone <url>` and go to the cloned project.

## Install

```shell
go install github.com/vxxvvxxv/go-git-clone
```

## Run

```shell
go-git-clone git@github.com:vxxvvxxv/go-git-clone.git
```
