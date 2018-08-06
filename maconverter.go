package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jessevdk/go-flags"
)

// Function to insert a character(rune) every N character in a string
// Usage example: fmt.Println(insertInto("helloworldhelloworldhelloworld", 5, '-'))
func insertInto(s string, interval int, sep rune) string {
	var buffer bytes.Buffer
	before := interval - 1
	last := len(s) - 1
	for i, char := range s {
		buffer.WriteRune(char)
		if i%interval == before && i != last {
			buffer.WriteRune(sep)
		}
	}
	return buffer.String()
}

type options struct {
	Help   bool `short:"h" long:"help" description:"Show this help message"`
	Cisco  bool `short:"C" long:"cisco" description:"Cisco format, xxxx.xxxx.xxxx"`
	ESS    bool `short:"E" long:"ess" description:"ESS format, xx:xx:xx:xx:xx:xx"`
	Huawei bool `short:"H" long:"huawei" description:"Huawei format, xxxx-xxxx-xxxx"`
}

func main() {
	// Put content of clipboard into macaddr variable
	macaddr, _ := clipboard.ReadAll()

	// Make a Regex for MAC-addresses, (a-f,A-F,0-9)
	MACregex, err := regexp.Compile("[^a-fA-F0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	// Run MACregex on macaddr variable and put into processedMAC variable
	processedMAC := strings.ToLower(MACregex.ReplaceAllString(macaddr, ""))

	var opts options
	p := flags.NewParser(&opts, flags.Default&^flags.HelpFlag)
	_, err = p.Parse()
	if err != nil {
		fmt.Printf("fail to parse args: %v", err)
		os.Exit(1)
	}

	if opts.Help {
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if len(processedMAC) == 12 {
		if opts.Cisco {
			// fmt.Println(processedMAC, "is a valid MAC address")
			// fmt.Println(insertInto(processedMAC, 4, '.'))
			clipboard.WriteAll(insertInto(processedMAC, 4, '.'))
		} else if opts.Huawei {
			// fmt.Println(processedMAC, "is a valid MAC address")
			// fmt.Println(insertInto(processedMAC, 4, '-'))
			clipboard.WriteAll(insertInto(processedMAC, 4, '-'))
		} else if opts.ESS {
			// fmt.Println(processedMAC, "is a valid MAC address")
			// fmt.Println(insertInto(processedMAC, 2, ':'))
			clipboard.WriteAll(insertInto(processedMAC, 2, ':'))
		}
	} else {
		fmt.Println(macaddr, "is not a valid MAC address")
	}

	// Print the values of the flags
	// fmt.Printf("Options: %v\n", opts)
}
