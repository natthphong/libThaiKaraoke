package thai2karaoke

import "strings"

func romanise(syl string) string {
	rs := []rune(syl)
	if len(rs) == 0 {
		return ""
	}

	// Preposed vowels
	var pre []rune
	i := 0
	for i < len(rs) && isPreVowel(rs[i]) {
		pre = append(pre, rs[i])
		i++
	}

	// Leading consonant cluster
	var cons []rune
	for i < len(rs) && isConsonant(rs[i]) {
		cons = append(cons, rs[i])
		i++
		// include cluster with ร ล ว
		if i < len(rs) && (rs[i] == 'ร' || rs[i] == 'ล' || rs[i] == 'ว') {
			cons = append(cons, rs[i])
			i++
		}
	}

	// Vowels around consonant
	var vow []rune
	for i < len(rs) && isVowel(rs[i]) {
		vow = append(vow, rs[i])
		i++
	}

	final := rs[i:]

	init := mapInitial(cons)
	nuc := mapVowel(string(pre) + string(vow))
	fin := mapFinal(final)

	out := init + nuc + fin
	return out
}

func isConsonant(r rune) bool { return isLeadingConsonant(r) }
func isVowel(r rune) bool {
	_, ok := vowel[string(r)]
	return ok || isPreVowel(r)
}
func isPreVowel(r rune) bool {
	return r == 'เ' || r == 'แ' || r == 'โ' || r == 'ใ' || r == 'ไ'
}

func mapInitial(rs []rune) string {
	var b strings.Builder
	for _, r := range rs {
		if v, ok := leading[r]; ok {
			b.WriteString(v)
		}
	}
	return b.String()
}

func mapFinal(rs []rune) string {
	var b strings.Builder
	for _, r := range rs {
		if v, ok := finals[r]; ok {
			b.WriteString(v)
		}
	}
	return b.String()
}

func mapVowel(code string) string {
	if v, ok := vowel[code]; ok {
		return v
	}
	// fallback join single characters
	var b strings.Builder
	for _, r := range code {
		if v, ok := vowel[string(r)]; ok {
			b.WriteString(v)
		}
	}
	if b.Len() == 0 {
		return "a"
	}
	return b.String()
}
