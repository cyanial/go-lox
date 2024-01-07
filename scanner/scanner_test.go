package scanner

import (
	"fmt"
	"github.com/cyanial/go-lox/token"
	"testing"
)

func TestScanner(t *testing.T) {

	source := `
var age = 12;
var tall = 1.7;
var hello = "world";

var _hell_12 = "zxc'";

for {

}

while {

}

func do() {

}

// this is comment

var x = 1;
var y = 2;
var z = (x+y);
var a = false;
var b = true;
var c = a and b;
var d = a or b;

x, y = y, x;

class student {

};
`

	sc := New(source)
	for tok := sc.ScanToken(); tok.Type != token.EOF; tok = sc.ScanToken() {
		fmt.Printf("%#v\n", *tok)
	}

}

func TestScannerUnTerminatedString(t *testing.T) {

	source := `
"asf`
	sc := New(source)
	for tok := sc.ScanToken(); tok.Type != token.EOF; tok = sc.ScanToken() {
		fmt.Printf("%#v\n", *tok)
	}

}

func TestScannerNumber(t *testing.T) {

	source := `123`
	sc := New(source)
	for tok := sc.ScanToken(); tok.Type != token.EOF; tok = sc.ScanToken() {
		fmt.Printf("%#v\n", *tok)
	}
}

func TestScannerDotNumber(t *testing.T) {
	source := `2.1`
	sc := New(source)
	for tok := sc.ScanToken(); tok.Type != token.EOF; tok = sc.ScanToken() {
		fmt.Printf("%#v\n", *tok)
	}
}

func TestScannerIdentity(t *testing.T) {

	source := `hello`
	sc := New(source)
	for tok := sc.ScanToken(); tok.Type != token.EOF; tok = sc.ScanToken() {
		fmt.Printf("%#v\n", *tok)
	}
}
