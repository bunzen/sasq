package sasq

import "testing"

func TestParse(t *testing.T) {
	asrec, err := parse_line("17.112.45.19 | 714 | 17.112.0.0/16 | APPLE-ENGINEERING | US | APPLE.COM | APPLE COMPUTER INC ")
	if err != nil {
		println(err)
		t.FailNow()
	}
	if asrec.ASN != 714 {
		t.Fail()
	}
	if asrec.Prefix.IP.String() != "17.112.0.0" {
		if !testing.Short() {
			println("Prefix IP != 17.112.0.0")
		}
		t.Fail()
	}
	if asrec.Prefix.Mask.String() != "ffff0000" {
		if !testing.Short() {
			println("Mask != 0xffff00000")
		}
		t.Fail()
	}
	if asrec.Name != "APPLE-ENGINEERING" {
		t.Fail()
	}
	if asrec.CC != "US" {
		t.Fail()
	}
	if asrec.Domain != "APPLE.COM" {
		t.Fail()
	}
	if asrec.ISP != "APPLE COMPUTER INC" {
		t.Fail()
	}
}
