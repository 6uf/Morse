package Morse

var (
	Morse map[string]string = map[string]string{
		"a": ".-",
		"b": "-..",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
		"0": "-----",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		" ": "/",
	}

	ReverseMorse map[string]string = ReturnReverse()
)

func ReturnReverse() (Reverse map[string]string) {
	Reverse = make(map[string]string)
	for Va, char := range Morse {
		Reverse[char] = Va
	}
	return
}
