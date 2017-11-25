# go-leveltravel-api
Golang client api for Level.Travel Api (V3).

[![Build Status](https://travis-ci.org/liderman/go-leveltravel-api.svg?branch=master)](https://travis-ci.org/liderman/go-leveltravel-api)&nbsp;[![GoDoc](https://godoc.org/github.com/liderman/go-leveltravel-api?status.svg)](https://godoc.org/github.com/liderman/go-leveltravel-api)


Installation
-----------
	go get github.com/liderman/go-leveltravel-api

Usage
-----------
Getting a list of hot tours:
```go
    api := NewLevelTravelApi("YOUR_APIKEY")
    reps, _ := api.HotTours("10.01.2018", "20.01.2018", "prices", HotToursFilter{})
    fmt.Println(reps);
```

Getting a list of airlines:
```go
    api := NewYandexRapsApi("YOUR_APIKEY")
    reps, err := api.Airlines()
    fmt.Println(reps);
```

Features
--------

* Support api version `3.0`
* Partial support references api
* Full support method `Hot tours`
* Without external dependencies

Requirements
------------

* Need at least `go1.2` or newer.

Documentation
-------------

You can read package documentation [here](http:godoc.org/github.com/liderman/go-leveltravel-api).
Official docs: [here](https://level.travel/affiliate/api_docs/json_api)

Road map
--------

* Full coverage by tests
* Full support references api
* Error handling from level travel api
* Support search methods
* Support orders methods
* Support pay methods (Probably not soon)

Testing
-------
Unit-tests:
```bash
go test -v
```