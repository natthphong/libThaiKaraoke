package thai2karaoke

import "testing"

func TestTranslate(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"สวัสดีครับ", "Sawatdi Khrap"},
		{"ผมชื่อณัฐครับ", "Phom Chue Nat Khrap"},
		{"กินข้าวหรือยัง", "Kin Khao Rue Yang"},
		{"กินข้าว", "Kin Khao"},
		{"ผม", "Phom"},
	}
	for _, tc := range tests {
		if got := Translate(tc.in); got != tc.out {
			t.Errorf("Translate(%q)=%q want %q", tc.in, got, tc.out)
		}
	}
}

func TestClassifyConvertRoundTrip(t *testing.T) {
	words := []string{"สวัสดี", "ครับ", "ผม", "ชื่อ", "ณัฐ", "กิน", "ข้าว", "หรือ", "ยัง"}
	for _, w := range words {
		tag := classify(w)
		got := convert(w, tag)
		if got == "" {
			t.Errorf("convert returned empty for %s", w)
		}
	}
}
