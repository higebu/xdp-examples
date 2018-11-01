package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	bpf "github.com/iovisor/gobpf/bcc"
)

const BPF_PROG_TYPE_XDP = 6

var (
	src      = flag.String("src", "", "source file path")
	progName = flag.String("name", "xdp_prog", "xdp prog name")
)

func main() {
	flag.Parse()
	if *src == "" || *progName == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	data, err := ioutil.ReadFile(*src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	module := bpf.NewModule(string(data), []string{
		"-w",
	})
	defer module.Close()

	_, err = module.Load(*progName, BPF_PROG_TYPE_XDP, 1, 65536)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load xdp prog: %v\n", err)
		os.Exit(1)
	}
}
