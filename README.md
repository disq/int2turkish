# int2turkish

Convert numbers to Turkish text.

## Signature

    func Text(num int64) []string

## Usage

```go
    import "github.com/disq/int2turkish"
```

Then:

```go
    int2turkish.Text("42") // []string{ "Kırk", "İki" }

    strings.Join(int2turkish.Text("42"), "") // Kırkİki
    strings.Join(int2turkish.Text("42"), " ") // Kırk İki
```

## Example Results

    go test -v .
    === RUN   TestTurkish
    --- PASS: TestTurkish (0.00s)
        int2turkish_test.go:57: Case #0: Input: 0 = "Sıfır"
        int2turkish_test.go:57: Case #1: Input: 1 = "Bir"
        int2turkish_test.go:57: Case #2: Input: 42 = "Kırkİki"
        int2turkish_test.go:57: Case #3: Input: 110 = "YüzOn"
        int2turkish_test.go:57: Case #4: Input: 1981 = "BinDokuzYüzSeksenBir"
        int2turkish_test.go:57: Case #5: Input: 5150 = "BeşBinYüzElli"
        int2turkish_test.go:57: Case #6: Input: 31337 = "OtuzBirBinÜçYüzOtuzYedi"
        int2turkish_test.go:57: Case #7: Input: 105300 = "YüzBeşBinÜçYüz"
        int2turkish_test.go:57: Case #8: Input: 110000042 = "YüzOnMilyonKırkİki"
    PASS
    ok      github.com/disq/int2turkish     0.007s
