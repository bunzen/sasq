package sasq

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type ASRecord struct {
	Query  string    `json:"query"`
	ASN    int       `json:"asn"`
	Prefix net.IPNet `json:"prefix"`
	Name   string    `json:"name"`
	CC     string    `json:"cc"`
	Domain string    `json:"domain"`
	ISP    string    `json:"isp"`
}

func parse_line(line string) (*ASRecord, error) {
	elems := strings.Split(line, " | ")
	_, ipnet, err := net.ParseCIDR(elems[2])
	if err != nil {
		return nil, fmt.Errorf("%v %v", err, line)
	}

	asn, err := strconv.Atoi(elems[1])
	if err != nil {
		return nil, err
	}
	return &ASRecord{
		strings.TrimSpace(elems[0]),
		asn,
		*ipnet,
		strings.TrimSpace(elems[3]),
		strings.TrimSpace(elems[4]),
		strings.TrimSpace(elems[5]),
		strings.TrimSpace(elems[6]),
	}, nil
}
