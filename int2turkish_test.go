package int2turkish

import (
	"strings"
	"testing"
)

func TestTurkish(t *testing.T) {
	cases := []struct {
		input  int64
		output string
	}{
		{
			input:  0,
			output: "Sıfır",
		},
		{
			input:  1,
			output: "Bir",
		},
		{
			input:  42,
			output: "Kırkİki",
		},
		{
			input:  110,
			output: "YüzOn",
		},
		{
			input:  1981,
			output: "BinDokuzYüzSeksenBir",
		},
		{
			input:  5150,
			output: "BeşBinYüzElli",
		},
		{
			input:  31337,
			output: "OtuzBirBinÜçYüzOtuzYedi",
		},
		{
			input:  105300,
			output: "YüzBeşBinÜçYüz",
		},
		{
			input:  110000042,
			output: "YüzOnMilyonKırkİki",
		},
	}

	for idx, cs := range cases {
		outSlc := Text(cs.input)
		out := strings.Join(outSlc, "")
		if out != cs.output {
			t.Errorf("Case #%v: Input: %v Want %q, got %q", idx, cs.input, cs.output, out)
		} else {
			t.Logf("Case #%v: Input: %v = %q", idx, cs.input, out)
		}
	}
}
