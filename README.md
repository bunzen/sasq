# sasq

A Go library designed to work with the Shadowserver IP-BGP service API.

## Usage


```go
package main

import "github.com/bunzen/sasq"

func main() {
	addrs := []string{"94.39.45.9", "44.55.33.22"}
	asns := sasq.QueryWhois(addrs)

	for _, asn := range asns {
		fmt.Printf("%v %v\n", asn.Query, asn.ASN)
	}
}
```

## License

Copyright © 2015,2017 Geir Skjotskift <geir@pogostick.net>

Distributed under the 2-Clause BSD License
