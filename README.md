# aozora

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