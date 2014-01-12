# go-patricia

**Build Status**: [![Build
Status](https://travis-ci.org/tchap/go-patricia.png?branch=master)](https://travis-ci.org/tchap/go-patricia)
**Test Coverate**: Comming as soon as Drone.io people update their Go.

#### About

A generic patricia trie (also called radix tree) implemented in Go (Golang).

The patricia trie as implemented in this library enables fast visitiong of
items. It is possible to

1. visit all items saved in the tree,
2. visit all items matching particular prefix (visit subtree), or
3. given a string, visit all items matching some prefix of the string.

`[]byte` type is used for keys, `interface{}` for values.

`Trie` is not thread safe. Synchronize the access yourself.

## State of the Project

This project is very much still in development and the API can change any time.

## Documentation

Check the generated documentation at [GoDoc](http://godoc.org/github.com/tchap/go-patricia/patricia).

## License

MIT, check the `LICENSE` file.

[![Gittip
Badge](http://img.shields.io/gittip/alanhamlett.png)](https://www.gittip.com/tchap/
"Gittip Badge")

[![Bitdeli
Badge](https://d2weczhvl823v0.cloudfront.net/tchap/go-patricia/trend.png)](https://bitdeli.com/free
"Bitdeli Badge")
