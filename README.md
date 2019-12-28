# aozora

[![Build Status](https://travis-ci.org/tgfjt/aozora.png?branch=master)](https://travis-ci.org/tgfjt/aozora)
[![GoDoc](https://godoc.org/github.com/tgfjt/aozora?status.svg)](https://godoc.org/github.com/tgfjt/aozora)

青空文庫に感謝して使う。

## Usage

```go
list, err := aozora.ListOfWorks()

if err != nil {
  log.Fatal(err)
}

fmt.Println(list[88].OriginBook()[0].ParentBook.Title)

// Output:
// 筑摩全集類聚版芥川龍之介全集
```

## Installation

```
$ go get github.com/tgfjt/aozora
```
