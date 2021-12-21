package translit

import "bytes"

var mapRuEn = map[string]string{
	"а": "a",
	"А": "A",
	"Б": "B",
	"б": "b",
	"В": "V",
	"в": "v",
	"Г": "G",
	"г": "g",
	"Д": "D",
	"д": "d",
	"Е": "E",
	"е": "e",
	"Ё": "E",
	"ё": "e",
	"Ж": "ZH",
	"ж": "zh",
	"З": "Z",
	"з": "z",
	"И": "I",
	"и": "i",
	"Й": "I",
	"й": "i",
	"К": "K",
	"к": "k",
	"Л": "L",
	"л": "l",
	"М": "M",
	"м": "m",
	"Н": "N",
	"н": "n",
	"О": "O",
	"о": "o",
	"П": "P",
	"п": "p",
	"Р": "R",
	"р": "r",
	"С": "S",
	"с": "s",
	"Т": "T",
	"т": "t",
	"У": "U",
	"у": "u",
	"Ф": "F",
	"ф": "f",
	"Х": "KH",
	"х": "kh",
	"Ц": "C",
	"ц": "c",
	"Ч": "CH",
	"ч": "ch",
	"Ш": "SH",
	"ш": "sh",
	"Щ": "SH",
	"щ": "sh",
	"Ъ": "",
	"ъ": "",
	"Ы": "",
	"ы": "",
	"Ь": "",
	"ь": "",
	"Э": "E",
	"э": "e",
	"Ю": "Yu",
	"ю": "yu",
	"Я": "JA",
	"я": "ya",
}

func Encode(text string) string {
	if text == "" {
		return ""
	}

	var input = bytes.NewBufferString(text)
	var output = bytes.NewBuffer(nil)

	var enStr string
	var ok bool

	for {
		rRune, _, err := input.ReadRune()

		if err != nil {
			break
		}

		if !isRussianChar(rRune) {
			output.WriteRune(rRune)
			continue
		}

		enStr, ok = mapRuEn[string(rRune)]

		if ok {
			output.WriteString(enStr)
			continue
		}
	}

	return output.String()
}

func isRussianChar(r rune) bool {
	switch {
	case r >= 1040 && r <= 1103,
		r == 1105, r == 1025:
		return true
	}

	return false
}
