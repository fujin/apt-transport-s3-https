package main

import (
	"bufio"
	"fmt"
	// "github.com/rlmcpherson/s3gof3r"
	"os"
)

var codes = map[int]string{
	100: "Capabilities",
	102: "Status",
	200: "URI Start",
	201: "URI Done",
	600: "URI Acquire",
	601: "Configuration",
}

type Header map[string]string

var capabilitiesHeader = &Header{
	"Version":     "1.2",
	"Send-Config": "True",
	"Pipelining":  "True",
}

type Message struct {
	value   int
	headers Header
}

func (m Message) message() string {
	return codes[m.value]
}

func sendCapabilities() {
	fmt.Println(&Message{100, *capabilitiesHeader})
}

func main() {
	sendCapabilities()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
