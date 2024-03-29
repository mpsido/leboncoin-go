# Fizzbuzz
 
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by “fizz”, all multiples of 5 by “buzz”, and all multiples of 15 by “fizzbuzz”. The output would look like this: “1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...”.

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers
 
Bonus question :
- Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request

## Install and run

[![Build Status](https://travis-ci.com/mpsido/leboncoin-go.svg?branch=master)](https://travis-ci.com/mpsido/leboncoin-go)

Download the code and select the directory
```
go get github.com/mpsido/leboncoin-go/...
cd $GOPATH/src/mpsido/leboncoin-go
```

Download dependencies(optional)
```
make dep
```

Run the application
```
make run
```

## Dev dependencies

This project is using [dep](https://github.com/golang/dep/blob/master/docs/installation.md) dependency manager.

```
go get -u github.com/golang/dep/cmd/dep
```

The go files in this repository are formated using [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)

```
go get golang.org/x/tools/cmd/goimports 
```

## Call the API

Using a webbrowser you can type the URL: `http://localhost:8080/?int1=2&int2=4&limit=18&str1=fizz&str2=buzz`

Alternatively you can use the following `curl` command `curl localhost:8080/?int1=3\&int2=5\&limit=16\&str1=fizz\&str2=buzz`

The application also supports `/stats` endpoint: `http://localhost:8080/stats`

## Using docker

First time you need to build the image:
```
docker build -t leboncoin-go .
```

Then you can run inside a docker container using the following command:
```
docker run -it -p 8080:8080 leboncoin-go
```

## Heroku deployment:

The application is deployed at this url: `https://protected-atoll-44592.herokuapp.com`

For example try:  `https://protected-atoll-44592.herokuapp.com/?int1=2&int2=4&limit=100&str1=fizz&str2=buzz`