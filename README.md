sandbox
=======

[![GitHub release](https://img.shields.io/github/release/keltia/sandbox.svg)](https://github.com/keltia/sandbox/releases)
[![GitHub issues](https://img.shields.io/github/issues/keltia/sandbox.svg)](https://github.com/keltia/sandbox/issues)
[![Go Version](https://img.shields.io/badge/go-1.10-blue.svg)](https://golang.org/dl/)
[![Build Status](https://travis-ci.org/keltia/sandbox.svg?branch=master)](https://travis-ci.org/keltia/sandbox)
[![GoDoc](http://godoc.org/github.com/keltia/sandbox?status.svg)](http://godoc.org/github.com/keltia/sandbox)
[![SemVer](http://img.shields.io/SemVer/2.0.0.png)](https://semver.org/spec/v2.0.0.html)
[![License](https://img.shields.io/pypi/l/Django.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Go Report Card](https://goreportcard.com/badge/github.com/keltia/sandbox)](https://goreportcard.com/report/github.com/keltia/sandbox)

Go library to create & manage lightweight sandboxes (like temporary files).

## Requirements

* Go >= 1.10

## Installation

This is a pure library, there is no associated command..

Installation is like many Go libraries with a simple

    go get github.com/keltia/sandbox

`sandbox` also has `vgo` support & metadata (see the articles on [vgo](https://research.swtch.com/vgo-intro)).  It respects the [Semantic Versioning](https://research.swtch.com/vgo-import) principle with tagged releases.

You can thus use

    vgo install github.com/keltia/sandbox

to install the library.

## API Usage

    snd, err := sandbox.New("foo")
    if err != nil {
        log.Fatalf("can not create sandbox: %v", err)
    }
    defer snd.Cleanup()

    err := snd.Enter()
    fmt.Println("Inside sandbox")
    err := snd.Exit()

You can also sandbox a single `func` with `Run()`:

    err := snd.Run(func() error {
        fmt.Printf("I am in %d", snd.Cwd())
    })

## License

This is under the 2-Clause BSD license, see `LICENSE.md`.

## History

I originally wrote this code for various projects of mine and have re-used it enough time to think about putting it into its own module.

## Contributing

Please see CONTRIBUTING.md for some simple rules.
