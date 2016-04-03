# minion [![Join the chat at https://gitter.im/HaMster21/minion](https://badges.gitter.im/HaMster21/minion.svg)](https://gitter.im/HaMster21/minion?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

:wrench: Work in progress task management

[![wercker status](https://app.wercker.com/status/aeebd6b26c9708c4ecc1050177a014aa/s "wercker status")](https://app.wercker.com/project/bykey/aeebd6b26c9708c4ecc1050177a014aa)
[![GoDoc](https://godoc.org/github.com/HaMster21/minion?status.svg)](https://godoc.org/github.com/HaMster21/minion)
[![Go Report Card](https://goreportcard.com/badge/github.com/hamster21/minion)](https://goreportcard.com/report/github.com/hamster21/minion)
[![codebeat badge](https://codebeat.co/badges/5f1222b7-c537-4ed9-89e7-85b09eaab33d)](https://codebeat.co/projects/github-com-hamster21-minion)
[![codecov.io](https://codecov.io/github/HaMster21/minion/coverage.svg?branch=master)](https://codecov.io/github/HaMster21/minion?branch=master)

Poor mans attempt to build something like Taskwarrior, but less restrictive and with Windows support

## Goals

Tasks are difficult to handle for humans and everyone seems to handle them
differently. I want minion to be flexible, so I try to provide and interface to
let minion do its work and implement some defaults that can be tweaked
individually.

My minion should be a commandline tool that helps me structuring my work. As
such, it is intended to be a single person, local machine application without a
fancy GUI.

## Assumptions

I will assume that there are two kinds of structures that my minion will
handle: Items and Collections thereof.

An Item can be something like a task, a habit, a note, a checklist entry or
whatever. By default it will be stored as a plain text file, so it can be
easily manipulated by anybody else. Every Item has (at least) the following
fields:

- Heading
- Description
- Collection it belongs to
- Position in the collection, for ordering
- Creation date/time
- Last change date/time
- Additional information, which is a key-value store of whatever you want

A Collection contains Items and adds ordering and a name to them. A collection
has (at least) the following fields:

- Name
- Creation date/time
- Last change date/time
- Archive flag, to exclude it from being rendered as active
- Archive date/time

## Roadmap

- Task and Collection interfaces
- Reading/Writing tasks from/to JSON|TOML|Markdown|Textile|AsciiDoc
- [TermUI Dashboard](https://github.com/gizak/termui)
- Output customization with go templates
- Item and Collection hooks that get triggered on events in the Item/Collection
- Storage backends other than the file system, like git (e.g. like
  [git-appraise](https://github.com/google/git-appraise)) or Redis or whatever
- rkt/docker image to keep items and collections in a single container
