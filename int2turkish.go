package int2turkish

var (
	Steps    = []int64{100, 1000, 1000000, 1000000000, 1000000000000}
	StepStrs = []string{"Yüz", "Bin", "Milyon", "Milyar", "Trilyon"}
	Tenths   = []string{"On", "Yirmi", "Otuz", "Kırk", "Elli", "Altmış", "Yetmiş", "Seksen", "Doksan"}
	Ones     = []string{"Sıfır", "Bir", "İki", "Üç", "Dört", "Beş", "Altı", "Yedi", "Sekiz", "Dokuz"}
)

// Text gets an int64 and returns a string slice in Turkish.
func Text(num int64) []string {
	var ret []string

	{
		for i := len(StepStrs) - 1; i >= 0; i-- {
			step := Steps[i]

			if num >= step {
				p := num / step
				if p > 1 {
					ret = append(ret, Text(p)...)
				}
				ret = append(ret, StepStrs[i])
				num = num - (p * step)
			}
		}
	}

	if step := int64(10); num >= step {
		p := num / step
		ret = append(ret, Tenths[p-1])
		num = num - (p * step)
	}
	if (len(ret) == 0 && num == 0) || (num > 0 && num < 10) {
		ret = append(ret, Ones[num])
	}

	return ret
}
