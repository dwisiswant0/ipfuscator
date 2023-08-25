package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"strings"

	"github.com/dwisiswant0/ipfuscator/pkg/ipfuscator"
	"github.com/gobwas/glob"
)

var (
	method string
	list   bool

	methods []string
)

const (
	version = "1.0.0"
	banner  = `
  ipfuscator ` + version + `
  by @dwisiswant0
  --
  A blazing-fast, thread-safe, straightforward and zero memory allocations
  tool to swiftly generate alternative IP(v4) address representations in Go.
`
	usage = `Usage
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
`
)

func init() {
	flag.StringVar(&method, "m", "*", "")
	flag.StringVar(&method, "method", "*", "")

	flag.BoolVar(&list, "l", false, "")
	flag.BoolVar(&list, "list", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n%s\n", banner, usage)
	}

	flag.Parse()
	fmt.Fprintf(os.Stderr, "%s\n", banner)

	methods = append(methods, getMethods()...)
}

func main() {
	if list {
		println(strings.Join(methods, "\n"))
		return
	}

	ip := flag.Arg(0)
	if ip == "" {
		log.Fatal(`no IP addr provided, use "-h" flag.`)
	}

	if method == "" {
		log.Fatal(`no conversion to run, use "-h" flag.`)
	}

	g, err := glob.Compile(method)
	if err != nil {
		log.Fatalf("failed to compile method: %s", err.Error())
	}

	ipv4 := net.ParseIP(ip)
	ipf, err := ipfuscator.New(ipv4)
	if err != nil {
		log.Fatalf("failed to construct IP addr: %s", err.Error())
	}

	ipfVal := reflect.ValueOf(ipf)

	for _, m := range methods {
		if g.Match(m) {
			methodVal := ipfVal.MethodByName(m)
			result := methodVal.Call(nil)
			if len(result) > 0 {
				fmt.Println(result[0].Interface().(string))
			}
		}
	}
}

func getMethods() []string {
	ipf := new(ipfuscator.IPFuscator)
	ipfType := reflect.TypeOf(ipf)

	var out []string
	for i := 0; i < ipfType.NumMethod(); i++ {
		method := ipfType.Method(i)
		out = append(out, method.Name)
	}

	return out
}
