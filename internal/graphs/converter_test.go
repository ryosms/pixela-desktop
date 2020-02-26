package graphs

import (
	"testing"
)

func TestTransform(t *testing.T) {
	t.Run("Can parse translate has no space char", func(t *testing.T) {
		x, y, err := transformXY("translate(100,200)")
		if err != nil {
			t.Fatal(err)
		}
		if x != 100 || y != 200 {
			t.Errorf("parse failed\n  x: expected is '100', actual is '%d'\n  y: expected is '200', actual is '%d'", x, y)
		}
	})
	t.Run("Can parse translate has space chars", func(t *testing.T) {
		x, y, err := transformXY("translate( 100 , 200 )")
		if err != nil {
			t.Fatal(err)
		}
		if x != 100 || y != 200 {
			t.Errorf("parse failed\n  x: expected is '100', actual is '%d'\n  y: expected is '200', actual is '%d'", x, y)
		}
	})
	t.Run("Can't parse unexpected string", func(t *testing.T) {
		_, _, err := transformXY("transform(100, 200)")
		if err == nil {
			t.Errorf("why can it be parsed?")
		}
	})
}

func TestFillToColor(t *testing.T) {
	t.Run("Can parse #ABCDEF", func(t *testing.T) {
		c, err := fillToColor("#ABCDEF")
		if err != nil {
			t.Fatal(err)
			return
		}
		if c.R != 0xAB || c.G != 0xCD || c.B != 0xEF {
			t.Errorf("invalid parse result(%+v)", c)
		}
	})
	t.Run("Can parse #ACE", func(t *testing.T) {
		c, err := fillToColor("#ACE")
		if err != nil {
			t.Fatal(err)
			return
		}
		if c.R != 0xAA || c.G != 0xCC || c.B != 0xEE {
			t.Errorf("invalid parse result(%+v)", c)
		}
	})
	t.Run("Can't parse invalid string", func(t *testing.T) {
		_, err := fillToColor("#1")
		if err == nil {
			t.Errorf("why can it be parsed?")
		}
	})
}
