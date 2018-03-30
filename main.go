package main

import (
	"fmt"
	"net"
	"os"

	"github.com/apparentlymart/go-cidr/cidr"
)

const usage = "sn <cidr>"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %v\n", usage)
		os.Exit(1)
	}

	nets := []*net.IPNet{}
	for _, sn := range os.Args[1:] {
		_, net, err := net.ParseCIDR(sn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v invalid: %v\n", sn, err)
			os.Exit(1)
		}
		nets = append(nets, net)
		a, b := cidr.AddressRange(net)
		l := cidr.AddressCount(net)
		fmt.Printf("%-16s %-16s %10d\n", a, b, l)
	}
	if err := cidr.NoOverlap(nets); err != nil {
		fmt.Fprintf(os.Stderr, "overlaping networks: %v\n", err)
		os.Exit(1)
	}
}
