package int2turkish

// Text gets an int64 and returns a string slice in Turkish.
func Text(num int64) []string {
	var ret []string

	{
		steps := []int64{100, 1000, 1000000, 1000000000, 1000000000000}
		strs := []string{"Yüz", "Bin", "Milyon", "Milyar", "Trilyon"}
		for i := len(strs) - 1; i >= 0; i-- {
			step := steps[i]

			if num >= step {
				p := num / step
				if p > 1 {
					ret = append(ret, Text(p)...)
				}
				ret = append(ret, strs[i])
				num = num - (p * step)
			}
		}
	}

	if step := int64(10); num >= step {
		p := num / step
		strs := []string{
			"On",
			"Yirmi",
			"Otuz",
			"Kırk",
			"Elli",
			"Altmış",
			"Yetmiş",
			"Seksen",
			"Doksan",
		}
		ret = append(ret, strs[p-1])
		num = num - (p * step)
	}
	if num > 0 && num < 10 {
		strs := []string{
			"Bir",
			"İki",
			"Üç",
			"Dört",
			"Beş",
			"Altı",
			"Yedi",
			"Sekiz",
			"Dokuz",
		}
		ret = append(ret, strs[num-1])
	}

	return ret
}
