package thai2karaoke

// convert applies a karaoke tag to a Thai word.
func convert(word, tag string) string {
	runes := []rune(word)
	get := func(i int) Cons {
		if i < len(runes) {
			return fw(string(runes[i]))
		}
		return Cons{}
	}
	f := get(0)
	s := get(1)
	t := get(2)

	switch tag {
	case "a":
		return f.First + "a"
	case "a1", "a2":
		return f.First + "a" + s.Spell
	case "a3":
		return f.First + "a" + t.Spell
	case "a4", "a5":
		return f.First + s.Spell + "a" + t.Spell
	case "a6":
		return f.First + "a" + s.Spell + t.Spell
	case "a7":
		return f.First + "a" + s.Spell
	case "a8":
		return f.First + s.Spell + "a" + t.Spell
	case "an":
		return f.First + "an"
	case "am":
		return f.First + "am"
	case "i":
		return f.First + "i"
	case "i1":
		return f.First + "i" + s.Spell
	case "i2", "i3":
		return f.First + s.Spell + "i" + t.Spell
	case "i4":
		return f.First + "i" + s.Spell
	case "ue":
		return f.First + "ue"
	case "ue1", "ue2":
		return f.First + "ue" + s.Spell
	case "ue3", "ue4":
		return f.First + s.Spell + "ue" + t.Spell
	case "u":
		return f.First + "u"
	case "u1":
		return f.First + "u" + s.Spell
	case "u2":
		return f.First + s.Spell + "u"
	case "u3":
		return f.First + s.Spell + "u" + t.Spell
	case "u4":
		return f.First + s.Spell + "u"
	case "e":
		return f.First + "e"
	case "e1", "e2":
		return f.First + "e" + s.Spell
	case "e3", "e4":
		return f.First + s.Spell + "e" + t.Spell
	case "ae":
		return f.First + "ae"
	case "ae1", "ae2":
		return f.First + "ae" + s.Spell
	case "o":
		return f.First + "o"
	case "o1":
		return f.First + "o"
	case "o2", "o3":
		return f.First + "o" + s.Spell
	case "o4", "o5":
		return f.First + s.Spell + "o" + t.Spell
	case "oe":
		return f.First + "oe"
	case "oe1":
		return f.First + s.Spell + "oe"
	case "oe2":
		return f.First + "oe" + s.Spell
	case "oe3":
		return f.First + s.Spell + "oe" + t.Spell
	case "oei":
		return f.First + "oei"
	case "ia":
		return f.First + "ia"
	case "ia1":
		return f.First + "ia" + s.Spell
	case "ia2":
		return f.First + s.Spell + "ia" + t.Spell
	case "uea":
		return f.First + "uea"
	case "uea1", "uea2":
		return f.First + "uea" + s.Spell
	case "uea3", "uea4":
		return f.First + s.Spell + "uea" + t.Spell
	case "ua":
		return f.First + "ua"
	case "ua1", "ua2":
		return f.First + "ua" + s.Spell
	case "ua3", "ua4":
		return f.First + s.Spell + "ua" + t.Spell
	case "uai":
		return f.First + s.Spell + "uai"
	case "ai":
		return f.First + "ai"
	case "ai1":
		return f.First + s.Spell + "ai"
	case "ao":
		return f.First + "ao"
	case "ao1", "ao2":
		return f.First + "ao" + s.Spell
	case "ui":
		return f.First + "ui"
	case "ui1":
		return f.First + s.Spell + "ui"
	default:
		return word
	}
}
