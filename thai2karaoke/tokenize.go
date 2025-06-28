package thai2karaoke

import (
	"unicode"
)

// tokenize splits text into Thai syllables or non-Thai chunks.
func tokenize(s string) []string {
	var res []string
	runes := []rune(s)
	i := 0
	for i < len(runes) {
		r := runes[i]
		if isSpaceRune(r) {
			res = append(res, string(r))
			i++
			continue
		}
		if !isThaiRune(r) {
			res = append(res, string(r))
			i++
			continue
		}
		start := i
		seenVowel := false
		i++
		for i < len(runes) {
			rr := runes[i]
			if isVowelRune(rr) {
				seenVowel = true
				i++
				continue
			}
			if isLeadingConsonant(rr) && !seenVowel && (rr == 'ร' || rr == 'ล' || rr == 'ว') {
				i++
				continue
			}
			if isLeadingConsonant(rr) && runes[i-1] != 0x0E3A {
				break
			}
			if !isThaiRune(rr) || isSpaceRune(rr) {
				break
			}
			i++
		}
		res = append(res, string(runes[start:i]))
	}
	return res
}

func isThaiRune(r rune) bool         { return r >= 0x0E00 && r <= 0x0E7F }
func isLeadingConsonant(r rune) bool { return r >= 0x0E01 && r <= 0x0E2E }
func isSpaceRune(r rune) bool        { return unicode.IsSpace(r) || r == 0x0E2F }

func isThaiString(s string) bool {
	for _, r := range s {
		if !isThaiRune(r) {
			return false
		}
	}
	return len(s) > 0
}

func isSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) && r != 0x0E2F {
			return false
		}
	}
	return true
}

func isPunct(s string) bool {
	if len(s) != 1 {
		return false
	}
	r := rune(s[0])
	return unicode.IsPunct(r)
}

func isVowelRune(r rune) bool {
	_, ok := vowel[string(r)]
	return ok || r == 'เ' || r == 'แ' || r == 'โ' || r == 'ใ' || r == 'ไ'
}
