# envcd

environment configurations detector/discovery/dictionary

[![license card](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg?label=license)](https://github.com/acmestack/envcd/blob/main/LICENSE)
[![go version](https://img.shields.io/github/go-mod/go-version/acmestack/envcd)](#)
[![go report](https://goreportcard.com/badge/github.com/acmestack/envcd)](https://goreportcard.com/report/github.com/acmestack/envcd)
[![codecov report](https://codecov.io/gh/acmestack/envcd/branch/main/graph/badge.svg)](https://codecov.io/gh/acmestack/envcd)
[![workflow](https://github.com/acmestack/envcd/actions/workflows/go.yml/badge.svg?event=push)](#)
[![lasted release](https://img.shields.io/github/v/release/acmestack/envcd?label=lasted)](https://github.com/acmestack/envcd/releases)

## Software Architecture
![Envcd Architecture](envcd.svg)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/acmestack/envcd.svg)](https://starchart.cc/acmestack/envcd)

## Contribute and Support

- [How to Contribute](https://acmestack.org/docs/contributing/guide/)

## Code Comment Polish

* you can comment with idea plugin `Gonano`

```go
// NewAsyncWriter Write data with Buffer, this Writer and Closer is thread safety, but WriteCloser parameters not safety.
//  @param w       Writer
//  @param bufSize accept buffer max length
//  @param block   if true, overflow buffer size, will blocking, if false will occur error
//  @return *AsyncLogWriter
func NewAsyncWriter(w io.Writer, bufSize int, block bool) *AsyncLogWriter {
}
```

