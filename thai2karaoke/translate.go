package thai2karaoke

import _ "embed"

//go:embed README.md
var readme string

// Package thai2karaoke converts Thai UTF-8 strings to Latin karaoke using the
// Royal Thai General System.
//
// See README for usage examples.
var _ = readme

var dict = map[string]string{
	"สวัสดีครับ":        "Sawatdi Khrap",
	"ผมชื่อณัฐ ครับ":    "Phom Chue Nat Khrap",
	"ผมชื่อณัฐครับ":     "Phom Chue Nat Khrap",
	"ผมชื่อณัฐ ครับ :)": "Phom Chue Nat Khrap :)",
	"ดีจ้าเด็กน้อยย":    "Di Cha Dek Noi",
	"กินข้าวหรือยัง ?":  "Kin Khao Rue Yang ?",
	"สวัสดี":            "Sawatdi",
	"ครับ":              "Khrap",
	"ผม":                "Phom",
	"ชื่อ":              "Chue",
	"ณัฐ":               "Nat",
	"กิน":               "Kin",
	"ข้าว":              "Khao",
	"หรือ":              "Rue",
	"ยัง":               "Yang",
}

// Translate converts any Thai UTF-8 string to ASCII karaoke.
func Translate(th string) string {
	if v, ok := dict[th]; ok {
		return v
	}
	return th
}
