package sasq

import (
	"fmt"
	"strings"
	"testing"
)

func TestOutputDownload(t *testing.T) {
	text := QueryWhois([]string{"195.88.54.16", "94.39.45.09"})

	for _, val := range text {
		fmt.Printf("Query  : %v\n", val.Query)
		fmt.Printf("Number : %v\n", val.ASN)
		fmt.Printf("Prefix : %v\n", val.Prefix)
		fmt.Printf("Name   : %v\n", val.Name)
		fmt.Printf("CN     : %v\n", val.CC)
		fmt.Printf("Domain : %v\n", val.Domain)
		fmt.Printf("ISP    : %v\n", val.ISP)
		fmt.Println("--------------------")
	}

}

func TestDownload(t *testing.T) {
	text := QueryWhois([]string{"94.39.45.09"})

	if text[0].Query != "94.39.45.09" {
		t.Fail()
	}
	if text[0].Prefix.IP.String() != "94.36.0.0" {
		t.Fail()
	}
	if text[0].Prefix.Mask.String() != "fffc0000" {
		t.Fail()
	}
	if text[0].Name != "TISCALI" {
		println(text[0].Name)
		t.Fail()
	}
	if text[0].CC != "IT" {
		t.Fail()
	}
	if text[0].Domain != "tiscali.it" {
		t.Fail()
	}
	if text[0].ISP != "Tiscalinet" {
		t.Fail()
	}
}

func TestWrap(t *testing.T) {
	wrap := wrapQuery([]string{"94.39.45.09"})
	if !strings.HasPrefix(wrap, "begin origin\n") {
		t.Fail()
	}
	if !strings.HasSuffix(wrap, "\nend\n") {
		t.Fail()
	}

}
