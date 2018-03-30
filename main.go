package main

import (
	"fmt"
	"net"
	"os"
)

const usage = "cidr <cidr>[, cidr2,...,cidrN]"

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
		a, b := AddressRange(net)
		l := AddressCount(net)
		fmt.Printf("%-16s %-16s %10d\n", a, b, l)
	}
	if err := NoOverlap(nets); err != nil {
		fmt.Fprintf(os.Stderr, "overlaping networks: %v\n", err)
		os.Exit(1)
	}
}
