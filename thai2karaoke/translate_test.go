package thai2karaoke

import "testing"

func TestTranslate(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"สวัสดีครับ", "Sawatdi Khrap"},
		{"ดีจ้าเด็กน้อยย", "Di Cha Dek Noi"},
		{"กินข้าวหรือยัง ?", "Kin Khao Rue Yang ?"},
		{"ผมชื่อณัฐ ครับ :)", "Phom Chue Nat Khrap :)"},
	}
	for _, tc := range tests {
		if got := Translate(tc.in); got != tc.out {
			t.Errorf("Translate(%q)=%q want %q", tc.in, got, tc.out)
		}
	}
}
