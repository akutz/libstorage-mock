package main

import (
	"fmt"
	"regexp"
)

var addrRX = regexp.MustCompile(
	`(?i)^((?:(?:tcp|udp|ip)[46]?)|(?:unix(?:gram|packet)?))://(.+)$`)

func parseProtoAddr(protoAddr string) (proto string, addr string, err error) {
	m := addrRX.FindStringSubmatch(protoAddr)
	if m == nil {
		return "", "", fmt.Errorf("invalid address: %v", protoAddr)
	}
	return m[1], m[2], nil
}
