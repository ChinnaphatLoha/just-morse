package main

import (
	"fmt"
	"os"
	"strings"
)

var morseCodeMap = map[rune]string{
	'A': ".-",    'B': "-...",  'C': "-.-.",  'D': "-..",   'E': ".",     'F': "..-.",
	'G': "--.",   'H': "....",  'I': "..",    'J': ".---",  'K': "-.-",   'L': ".-..",
	'M': "--",    'N': "-.",    'O': "---",   'P': ".--.",  'Q': "--.-",  'R': ".-.",
	'S': "...",   'T': "-",     'U': "..-",   'V': "...-",  'W': ".--",   'X': "-..-",
	'Y': "-.--",  'Z': "--..",  '1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----",
	' ': "/",     ',': "--..--", '.': ".-.-.-", '?': "..--..", '!': "-.-.--", '@': ".--.-.",
}

var textMap = reverseMap(morseCodeMap)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No command provided. Use 'morse \"<message>\"' or 'morse -t \"<morse code>\"'.")
		return
	}

	switch os.Args[1] {
	case "--translate", "-t":
		if len(os.Args) < 3 {
			fmt.Println("Error: No Morse code provided to translate.")
			return
		}
		translatedText, err := morseToText(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Translated text:", translatedText)
		}
	default:
		message := strings.Join(os.Args[1:], " ")
		morseCode, err := textToMorse(message)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Morse code:", morseCode)
		}
	}
}

func textToMorse(text string) (string, error) {
	var morseCode []string
	for _, char := range strings.ToUpper(text) {
		code, found := morseCodeMap[char]
		if !found {
			return "", fmt.Errorf("unsupported character '%c' in message", char)
		}
		morseCode = append(morseCode, code)
	}
	return strings.Join(morseCode, " "), nil
}

func morseToText(morse string) (string, error) {
	var text []string
	for _, code := range strings.Split(morse, " ") {
		char, found := textMap[code]
		if !found {
			return "", fmt.Errorf("unsupported Morse code '%s' in input", code)
		}
		text = append(text, string(char))
	}
	return strings.Join(text, ""), nil
}

func reverseMap(m map[rune]string) map[string]rune {
	n := make(map[string]rune)
	for k, v := range m {
		n[v] = k
	}
	return n
}