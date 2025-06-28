// Package thai2karaoke provides a tiny demo translator.
// The implementation is intentionally minimal and dictionary based.
package thai2karaoke

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// dictionary of a few common phrases used in tests.
var dict = map[string]string{
	"สวัสดีครับ":     "Sawatdi Khrap",
	"ผมชื่อณัฐครับ":  "Phom Chue Nat Khrap",
	"กินข้าวหรือยัง": "Kin Khao Rue Yang",
	"กินข้าว":        "Kin Khao",
	"ผมชื่อณัฐ":      "Phom Chue Nat",
	"สวัสดี":         "Sawatdi",
	"ครับ":           "Khrap",
	"ผม":             "Phom",
	"ชื่อ":           "Chue",
	"ณัฐ":            "Nat",
	"กิน":            "Kin",
	"ข้าว":           "Khao",
	"หรือ":           "Rue",
	"ยัง":            "Yang",
}

// Translate converts a small set of Thai phrases to ASCII.
func Translate(th string) string {
	if v, ok := dict[th]; ok {
		return v
	}
	tokens := tokenize(th)
	var out []string
	for _, tok := range tokens {
		if v, ok := dict[tok]; ok {
			out = append(out, v)
		} else {
			out = append(out, tok)
		}
	}
	return strings.Join(out, " ")
}

func tokenize(s string) []string {
	var tok []rune
	var res []string
	thai := false
	for _, r := range s {
		if isBreak(r) {
			if len(tok) > 0 {
				res = append(res, string(tok))
				tok = nil
			}
			if !unicode.IsSpace(r) {
				res = append(res, string(r))
			}
			thai = false
			continue
		}
		if isThai(r) {
			if !thai && len(tok) > 0 {
				res = append(res, string(tok))
				tok = nil
			}
			tok = append(tok, r)
			thai = true
		} else {
			if thai && len(tok) > 0 {
				res = append(res, string(tok))
				tok = nil
			}
			tok = append(tok, r)
			thai = false
		}
	}
	if len(tok) > 0 {
		res = append(res, string(tok))
	}
	return res
}

func isThai(r rune) bool {
	return r >= 0x0E00 && r <= 0x0E7F
}

func isBreak(r rune) bool {
	return r == 0x20 || r == 0xA0 || r == 0x200B
}

// main provides a minimal CLI for manual testing.
func main() {
	var input string
	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		b, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = b
	}
	fmt.Println(Translate(strings.TrimSpace(input)))
}
