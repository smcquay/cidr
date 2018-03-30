# cidr

This tool reports ip address ranges for a collection of CIDRs, and then reports
if there is any overlap.

The official package name is [mcquay.me/cidr](https://mcquay.me/cidr?go-get=1).

## usage

```bash
# prints range
$ cidr 10.0.0.0/24
10.0.0.0         10.0.0.255              256

# range plus no error for overlap:
$ cidr 10.0.0.0/24 10.0.2.0/24
10.0.0.0         10.0.0.255              256
10.0.2.0         10.0.2.255              256

# error on overlap
$ cidr 10.0.0.0/24 10.0.0.5/32
10.0.0.0         10.0.0.255              256
10.0.0.5         10.0.0.5                  1
overlaping networks: 10.0.0.5/32 overlaps with 10.0.0.0/24
$ cidr 10.0.1.0/24 10.0.0.0/16
10.0.1.0         10.0.1.255              256
10.0.0.0         10.0.255.255          65536
overlaping networks: 10.0.1.0/24 overlaps with 10.0.0.0/16
```
