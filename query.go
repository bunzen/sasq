// sasq (Shadowserver AS Query) is a small library to
// query the Shadowserver IP-BGP Whois database; using the
// bulk query API
package sasq

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// QueryWhois resolves a list of IP-addresses into
// AS number with information.
func QueryWhois(addresses []string) []ASRecord {
	ret := []ASRecord{}

	conn, err := net.Dial("tcp", "asn.shadowserver.org:43")
	if err != nil {
		println("Unable to contact asn.shadowserver.org")
		println(err.Error())
		return ret
	}
	defer conn.Close()

	fmt.Fprintf(conn, wrapQuery(addresses))
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		println(err.Error())
		return ret
	}
	err_count := 0
	for line != "" {
		rec, err := parse_line(line)

		if err == nil {
			ret = append(ret, *rec)
		} else {
			err_count++
		}

		line, err = reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				log.Println(err)
				return ret
			}
		}
	}
	if err_count > 0 {
		log.Printf("Unable to resolve %v addresses\n", err_count)
	}

	return ret
}

func wrapQuery(address []string) string {
	query := "begin origin\n"
	query += strings.Join(address, "\n")
	query += "\nend\n"
	return query
}
