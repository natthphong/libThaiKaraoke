package thai2karaoke

// Cons represents a Thai consonant mapping.
type Cons struct {
	Th    string
	First string
	Spell string
}

// table of 44 consonants.
var table = []Cons{
	{"ก", "k", "k"},
	{"ข", "kh", "k"},
	{"ฃ", "kh", "k"},
	{"ค", "kh", "k"},
	{"ฅ", "kh", "k"},
	{"ฆ", "kh", "k"},
	{"ง", "ng", "ng"},
	{"จ", "ch", "t"},
	{"ฉ", "ch", "t"},
	{"ช", "ch", "t"},
	{"ซ", "s", "t"},
	{"ฌ", "ch", "t"},
	{"ญ", "y", "n"},
	{"ฎ", "d", "t"},
	{"ฏ", "t", "t"},
	{"ฐ", "th", "t"},
	{"ฑ", "th", "t"},
	{"ฒ", "th", "t"},
	{"ณ", "n", "n"},
	{"ด", "d", "t"},
	{"ต", "t", "t"},
	{"ถ", "th", "t"},
	{"ท", "th", "t"},
	{"ธ", "th", "t"},
	{"น", "n", "n"},
	{"บ", "b", "p"},
	{"ป", "p", "p"},
	{"ผ", "ph", "p"},
	{"ฝ", "f", "p"},
	{"พ", "ph", "p"},
	{"ฟ", "f", "p"},
	{"ภ", "ph", "p"},
	{"ม", "m", "m"},
	{"ย", "y", "y"},
	{"ร", "r", "n"},
	{"ล", "l", "n"},
	{"ว", "w", "w"},
	{"ศ", "s", "t"},
	{"ษ", "s", "t"},
	{"ส", "s", "t"},
	{"ห", "h", "h"},
	{"ฬ", "l", "n"},
	{"อ", "o", "o"},
	{"ฮ", "h", "h"},
}

// fw finds a consonant mapping.
func fw(ch string) Cons {
	for _, c := range table {
		if c.Th == ch {
			return c
		}
	}
	return Cons{Th: ch}
}
