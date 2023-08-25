# ipfuscator

**ipfuscator** is a _blazing-fast_, thread-safe, straightforward and zero memory allocations tool to swiftly generate alternative IP(v4) address representations through conversion and manipulation techniques.

## Install

### CLI

**with Binary**

Just grab the bin based on your arch from [release page](https://github.com/dwisiswant0/ipfuscator/releases).

**with Go**

> **Note**
> Installation of [Go](https://golang.org/doc/install) version 1.16 or newer is required.

```console
go install github.com/dwisiswant0/ipfuscator@latest
```

### Library

```console
go get github.com/dwisiswant0/ipfuscator/pkg/ipfuscator@latest
```

## Usage

### Tool

```
  ipfuscator 1.0.0
  by @dwisiswant0
  --
  A blazing-fast, thread-safe, straightforward and zero memory allocations
  tool to swiftly generate alternative IP(v4) address representations in Go.

Usage
  ipfuscator [OPTIONS] --method [METHOD] [ADDR]

Options:
  -m, --method <METHOD>    Specify the glob conversion method to use (default: "*")
  -l, --list               Display a list of available conversion methods

Examples:
  ipfuscator -l
  ipfuscator 127.0.0.1
  ipfuscator -m "To*" 127.0.0.1
  ipfuscator -m "*Padding" 127.0.0.1
  ipfuscator -m "ToHexWithPadding" 127.0.0.1
```

### Library

[![GoDoc](https://pkg.go.dev/static/frontend/badge/badge.svg)](http://pkg.go.dev/github.com/dwisiswant0/ipfuscator/pkg/ipfuscator)

See the [docs](http://pkg.go.dev/github.com/dwisiswant0/ipfuscator/pkg/ipfuscator).

**Examples**

```go
package main

import (
	"fmt"
	"net"

	"github.com/dwisiswant0/ipfuscator/pkg/ipfuscator"
)

func main() {
	ipv4 := net.ParseIP("127.0.0.1")
	ipf, err := ipfuscator.New(ipv4)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Base w/ Padding: \t%s\n", ipf.ToBaseWithPadding())
	fmt.Printf("Circled digits: \t%s\n", ipf.ToCircledDigits())
	fmt.Printf("Decimal: \t\t%s\n", ipf.ToDecimal())
	fmt.Printf("Hex w/ Padding: \t%s\n", ipf.ToHexWithPadding())
	fmt.Printf("Hex: \t\t\t%s\n", ipf.ToHex())
	fmt.Printf("IPv6 Embedded: \t\t%s\n", ipf.ToIPv6CompatibleV4())
	fmt.Printf("No zeros: \t\t%s\n", ipf.ToNoZeros())
	fmt.Printf("Octal w/ Padding: \t%s\n", ipf.ToOctalWithPadding())
	fmt.Printf("Octal: \t\t\t%s\n", ipf.ToOctal())
	fmt.Printf("Rand 8-bits: \t\t%s\n", ipf.ToRand8Bits())
	fmt.Printf("Rand w/ Padding: \t%s\n", ipf.ToRandBaseWithPadding())
	fmt.Printf("Rand: \t\t\t%s\n", ipf.ToRandBase())

	// Output:
	// Base w/ Padding:     127.00000000000000.000000000.1
	// Circled digits:      ①②⑦.⓪.⓪.①
	// Decimal:             2130706433
	// Hex w/ Padding:      0x0000000000007f.0x0000000000000000000000.0x0000000000000000000000.0x0000000001
	// Hex:                 0x7f.0x0.0x0.0x1
	// IPv6 Embedded:       [::ffff:127.0.0.1]
	// No zeros:            127.1
	// Octal w/ Padding:    00000000000177.000000.0000000000000000.0000000000000000001
	// Octal:               0177.0.0.1
	// Rand 8-bits:         127.0.0.119
	// Rand w/ Padding:     0x000000000000000000000007f.0x00000000000000000000000000.0x0000000.0x01
	// Rand:                0x7f.0x0.0x0.0x1
}
```

## Benchmark

```console
$ make bench
goos: linux
goarch: amd64
pkg: github.com/dwisiswant0/ipfuscator/pkg/ipfuscator
cpu: 11th Gen Intel(R) Core(TM) i9-11900H @ 2.50GHz
BenchmarkToBaseWithPadding-16        	 6483795	       178.3 ns/op
BenchmarkToCircledDigits-16          	 8223832	       142.3 ns/op
BenchmarkToDecimal-16                	43212138	        26.25 ns/op
BenchmarkToHex-16                    	31927065	        40.15 ns/op
BenchmarkToHexWithPadding-16         	 1927899	       635.0 ns/op
BenchmarkToIPv6CompatibleV4-16       	16217696	        72.04 ns/op
BenchmarkToNoZeros-16                	26995516	        45.83 ns/op
BenchmarkToOctal-16                  	22103026	        55.83 ns/op
BenchmarkToOctalWithPadding-16       	 3390346	       359.3 ns/op
BenchmarkToRand8Bits-16              	16113920	        75.93 ns/op
BenchmarkToRandBase-16               	 9400024	       128.6 ns/op
BenchmarkToRandBaseWithPadding-16    	 2534938	       472.0 ns/op
PASS
ok  	github.com/dwisiswant0/ipfuscator/pkg/ipfuscator	16.722s
```

## License

**ipfuscator** is released under [Apache-2.0](/LICENSE). Copyright (c) 2023 Dwi Siswanto.