package common

var alphabet map[byte]string

func InitMap() {
	alphabet = make(map[byte]string, 127)
	alphabet[0x04] = "a"
	alphabet[0x05] = "b"
	alphabet[0x06] = "c"
	alphabet[0x07] = "d"
	alphabet[0x08] = "e"
	alphabet[0x09] = "f"
	alphabet[0x0a] = "g"
	alphabet[0x0b] = "h"
	alphabet[0x0c] = "i"
	alphabet[0x0d] = "j"
	alphabet[0x0e] = "k"
	alphabet[0x0f] = "l"
	alphabet[0x10] = "m"
	alphabet[0x11] = "n"
	alphabet[0x12] = "o"
	alphabet[0x13] = "p"
	alphabet[0x14] = "q"
	alphabet[0x15] = "r"
	alphabet[0x16] = "s"
	alphabet[0x17] = "t"
	alphabet[0x18] = "u"
	alphabet[0x19] = "v"
	alphabet[0x1a] = "w"
	alphabet[0x1b] = "x"
	alphabet[0x1c] = "y"
	alphabet[0x1d] = "z"
	alphabet[0x1e] = "1"
	alphabet[0x1f] = "2"
	alphabet[0x20] = "3"
	alphabet[0x21] = "4"
	alphabet[0x22] = "5"
	alphabet[0x23] = "6"
	alphabet[0x24] = "7"
	alphabet[0x25] = "8"
	alphabet[0x26] = "9"
	alphabet[0x27] = "0"
	//alphabet[0x28] = enter

}

func GetCharacter(value byte) string {
	return alphabet[value]
}
