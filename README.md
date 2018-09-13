# Roman Numeral Calculator

Start by writing a Roman Numeral converter and from that Roman Numerals. Then
write a Roman Numeral Calculator that takes a string i.e. "XXIV + XI" and
returns the result which would be "XXXV". Ideally it should perform all BODMAS
calculation so:
- Brackets: ()
- Order (power): ^
- Division: /
- Multiplication: *
- Addition: +
- Subtraction: -

# Building and Running

You can choose to either build a local version or a Docker container.
Afterwards, just navigate your browser http://localhost:8080.

## Local

1. Get Golang (https://golang.org/)
2. Get dep (https://github.com/golang/dep)
3. `dep ensure`
4. `go build`
5. `./roman-calc`

If you want to run tests:
`
go test test/*
`

## Docker

1. `docker build -t roman-calc .`
2. `docker run --rm -it -p 8080:8080 roman-calc`
