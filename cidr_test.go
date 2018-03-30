package main

import (
	"fmt"
	"net"
	"testing"
)

func TestAddressRange(t *testing.T) {
	type Case struct {
		Range string
		First string
		Last  string
	}

	cases := []Case{
		Case{
			Range: "192.168.0.0/16",
			First: "192.168.0.0",
			Last:  "192.168.255.255",
		},
		Case{
			Range: "192.168.0.0/17",
			First: "192.168.0.0",
			Last:  "192.168.127.255",
		},
		Case{
			Range: "fe80::/64",
			First: "fe80::",
			Last:  "fe80::ffff:ffff:ffff:ffff",
		},
	}

	for _, testCase := range cases {
		_, network, _ := net.ParseCIDR(testCase.Range)
		firstIP, lastIP := AddressRange(network)
		desc := fmt.Sprintf("AddressRange(%#v)", testCase.Range)
		gotFirstIP := firstIP.String()
		gotLastIP := lastIP.String()
		if gotFirstIP != testCase.First {
			t.Errorf("%s first is %s; want %s", desc, gotFirstIP, testCase.First)
		}
		if gotLastIP != testCase.Last {
			t.Errorf("%s last is %s; want %s", desc, gotLastIP, testCase.Last)
		}
	}

}

func TestAddressCount(t *testing.T) {
	type Case struct {
		Range string
		Count uint64
	}

	cases := []Case{
		Case{
			Range: "192.168.0.0/16",
			Count: 65536,
		},
		Case{
			Range: "192.168.0.0/17",
			Count: 32768,
		},
		Case{
			Range: "192.168.0.0/32",
			Count: 1,
		},
		Case{
			Range: "192.168.0.0/31",
			Count: 2,
		},
		Case{
			Range: "0.0.0.0/0",
			Count: 4294967296,
		},
		Case{
			Range: "0.0.0.0/1",
			Count: 2147483648,
		},
		Case{
			Range: "::/65",
			Count: 9223372036854775808,
		},
		Case{
			Range: "::/128",
			Count: 1,
		},
		Case{
			Range: "::/127",
			Count: 2,
		},
	}

	for _, testCase := range cases {
		_, network, _ := net.ParseCIDR(testCase.Range)
		gotCount := AddressCount(network)
		desc := fmt.Sprintf("AddressCount(%#v)", testCase.Range)
		if gotCount != testCase.Count {
			t.Errorf("%s = %d; want %d", desc, gotCount, testCase.Count)
		}
	}

}

func TestVerifyNetowrk(t *testing.T) {

	type testVerifyNetwork struct {
		CIDRBlock string
		CIDRList  []string
	}

	testCases := []*testVerifyNetwork{
		&testVerifyNetwork{
			CIDRList: []string{
				"192.168.8.0/24",
				"192.168.9.0/24",
				"192.168.10.0/24",
				"192.168.11.0/25",
				"192.168.11.128/25",
				"192.168.12.0/25",
				"192.168.12.128/26",
				"192.168.12.192/26",
				"192.168.13.0/26",
				"192.168.13.64/27",
				"192.168.13.96/27",
				"192.168.13.128/27",
			},
		},
	}
	failCases := []*testVerifyNetwork{
		&testVerifyNetwork{
			CIDRList: []string{
				"192.168.8.0/24",
				"192.168.9.0/24",
				"192.168.10.0/24",
				"192.168.11.0/25",
				"192.168.11.128/25",
				"192.168.12.0/25",
				"192.168.12.64/26",
				"192.168.12.128/26",
			},
		},
		&testVerifyNetwork{
			CIDRList: []string{
				"192.168.7.0/24",
				"192.168.9.0/24",
				"192.168.10.0/24",
				"192.168.11.0/25",
				"192.168.11.128/25",
				"192.168.12.0/25",
				"192.168.12.64/26",
				"192.168.12.128/26",
			},
		},
	}

	for _, tc := range testCases {
		subnets := make([]*net.IPNet, len(tc.CIDRList))
		for i, s := range tc.CIDRList {
			_, n, err := net.ParseCIDR(s)
			if err != nil {
				t.Errorf("Bad test data %s\n", s)
			}
			subnets[i] = n
		}
		test := NoOverlap(subnets)
		if test != nil {
			t.Errorf("Failed test with %v\n", test)
		}
	}
	for _, tc := range failCases {
		subnets := make([]*net.IPNet, len(tc.CIDRList))
		for i, s := range tc.CIDRList {
			_, n, err := net.ParseCIDR(s)
			if err != nil {
				t.Errorf("Bad test data %s\n", s)
			}
			subnets[i] = n
		}
		test := NoOverlap(subnets)
		if test == nil {
			t.Errorf("Test should have failed with CIDR %s\n", tc.CIDRBlock)
		}
	}
}
