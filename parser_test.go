package sasq

import "testing"

func TestParse(t *testing.T) {
	asrec, err := parse_line("17.112.45.19 | 714 | 17.112.0.0/16 | APPLE-ENGINEERING | US | APPLE COMPUTER INC ")
	if err != nil {
		println(err)
		t.FailNow()
	}
	if asrec.ASN != 714 {
		t.Fail()
	}
	if asrec.Prefix != "17.112.0.0/16" {
		if !testing.Short() {
			println("Prefix IP != 17.112.0.0/16")
		}
		t.Fail()
	}
	if asrec.Name != "APPLE-ENGINEERING" {
		t.Fail()
	}
	if asrec.CC != "US" {
		t.Fail()
	}
	if asrec.ISP != "APPLE COMPUTER INC" {
		t.Fail()
	}
}
