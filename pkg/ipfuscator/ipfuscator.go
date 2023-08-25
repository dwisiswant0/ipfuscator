// Package ipfuscator provides implementations of IPv4 address obfuscator.
// This package is designed to offer various methods for generating
// alternative IPv4 address representations through conversion and
// manipulation techniques.
package ipfuscator

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"math/rand"
)

// IPFuscator represents a utility for converting and obfuscating IPv4 addresses.
type IPFuscator struct {
	// IP is IPv4 address in a 4-byte representation
	IP net.IP

	hex, ips, octal []string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New creates a new IPFuscator instance based on the provided IPv4 address.
func New(ip net.IP) (*IPFuscator, error) {
	if ip == nil {
		return nil, errors.New(errNilIP)
	}

	v4 := ip.To4()

	if v4 == nil {
		return nil, fmt.Errorf(errCannotConvert, ip.String())
	}

	ipf := new(IPFuscator)
	ipf.IP = v4
	ipf.hex = make([]string, len(v4))
	ipf.ips = make([]string, len(v4))
	ipf.octal = make([]string, len(v4))

	var builder strings.Builder

	for i, part := range v4 {
		builder.Reset()

		builder.WriteString("0x")
		builder.WriteString(strconv.FormatInt(int64(part), 16))
		ipf.hex[i] = builder.String()

		builder.Reset()
		builder.WriteString(strconv.Itoa(int(part)))
		ipf.ips[i] = builder.String()

		builder.Reset()
		builder.WriteString(strconv.FormatInt(int64(part), 8))
		ipf.octal[i] = builder.String()
	}

	return ipf, nil
}

// ToBaseWithPadding returns the fuscated IPv4 address using base conversion
// with random zero padding.
func (ipf *IPFuscator) ToBaseWithPadding() string {
	var out strings.Builder

	for i, part := range ipf.ips {
		if i > 0 {
			out.WriteString(".")
		}

		if ipf.ips[i] == "0" {
			for j := 0; j < rand.Intn(30)+1; j++ {
				out.WriteString("0")
			}
		}

		// pad := strings.Repeat("0", rand.Intn(30)+1)
		// if ipf.ips[i] == "0" {
		// 	out.WriteString(pad)
		// }
		out.WriteString(part)
	}

	return out.String()
}

// ToCircledDigits returns the IPv4 address in circled digits unicode form.
func (ipf *IPFuscator) ToCircledDigits() string {
	var out strings.Builder

	for _, digit := range ipf.IP.String() {
		if '0' <= digit && digit <= '9' {
			num := int(digit - '0')
			out.WriteRune(circledDigitsCharset[num])
		} else {
			out.WriteRune(digit)
		}
	}

	return out.String()
}

// ToDecimal returns the IPv4 address in decimal form.
func (ipf *IPFuscator) ToDecimal() string {
	dec := uint32(ipf.IP[0])<<24 |
		uint32(ipf.IP[1])<<16 |
		uint32(ipf.IP[2])<<8 |
		uint32(ipf.IP[3])

	return strconv.FormatUint(uint64(dec), 10)
}

// ToHex returns the IPv4 address in hexadecimal form.
func (ipf *IPFuscator) ToHex() string {
	return strings.Join(ipf.hex, ".")
}

// ToHexWithPadding returns the IPv4 address in hexadecimal form
// with random zero padding.
func (ipf *IPFuscator) ToHexWithPadding() string {
	var out strings.Builder

	for i := range ipf.hex {
		if i > 0 {
			out.WriteString(".")
		}

		pad := strings.Repeat("0", rand.Intn(30)+1)
		hex := strings.Split(ipf.hex[i], "0x")
		out.WriteString("0x")
		out.WriteString(pad)
		out.WriteString(hex[1])
	}

	return out.String()
}

// ToIPv6CompatibleV4 returns the IPv4 address embedded in IPv6 form.
func (ipf *IPFuscator) ToIPv6CompatibleV4() string {
	var out strings.Builder
	out.WriteString("[::ffff:")
	out.WriteString(ipf.IP.String())
	out.WriteString("]")

	return out.String()
}

// ToNoZeros returns the IPv4 address by dropping the zeros form.
func (ipf *IPFuscator) ToNoZeros() string {
	var out strings.Builder

	for i, part := range ipf.IP {
		if int(part) == 0 {
			continue
		}

		if i > 0 {
			out.WriteString(".")
		}

		out.WriteString(ipf.ips[i])
	}

	return out.String()
}

// ToOctal returns the IPv4 address in octal form.
func (ipf *IPFuscator) ToOctal() string {
	return "0" + strings.Join(ipf.octal, ".")
}

// ToOctalWithPadding returns the IPv4 address in octal form
// with random zero padding.
func (ipf *IPFuscator) ToOctalWithPadding() string {
	var out strings.Builder

	for i := range ipf.hex {
		if i > 0 {
			out.WriteString(".")
		}

		pad := strings.Repeat("0", rand.Intn(30)+1)
		out.WriteString(pad)
		out.WriteString(ipf.octal[i])
	}

	return out.String()
}

// ToRand8Bits returns the random last 8 bits IPv4 address.
func (ipf *IPFuscator) ToRand8Bits() string {
	mask := net.CIDRMask(8, 32)
	ip := make(net.IP, 4)
	copy(ip, ipf.IP.Mask(mask))
	ip[3] = byte(rand.Intn(255))

	return ip.String()
}

// ToRandBase returns the IPv4 address using a random base conversion.
func (ipf *IPFuscator) ToRandBase() string {
	return ipf.genRandBase()
}

// ToRandBaseWithPadding returns the IPv4 address using a random base
// conversion with random zero padding.
func (ipf *IPFuscator) ToRandBaseWithPadding() string {
	return ipf.genRandBaseWithPadding()
}

func (ipf *IPFuscator) genRandBase() string {
	var out strings.Builder

	for i := range ipf.ips {
		if i > 0 {
			out.WriteString(".")
		}

		switch rand.Intn(3) {
		case 0:
			out.WriteString(ipf.ips[i])
		case 1:
			out.WriteString(ipf.hex[i])
		case 2:
			if i == 0 {
				out.WriteString("0")
			}
			out.WriteString(ipf.octal[i])
		}
	}

	return out.String()
}

func (ipf *IPFuscator) genRandBaseWithPadding() string {
	var out strings.Builder

	switch rand.Intn(3) {
	case 0:
		out.WriteString(ipf.ToBaseWithPadding())
	case 1:
		out.WriteString(ipf.ToHexWithPadding())
	case 2:
		out.WriteString(ipf.ToOctalWithPadding())
	}

	return out.String()
}
