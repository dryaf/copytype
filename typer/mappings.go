package typer

import "github.com/micmonay/keybd_event"

var KeysOSXus = map[string]KeyPressCombo{
	"a": {false, keybd_event.VK_A},
	"A": {true, keybd_event.VK_A},

	"b": {false, keybd_event.VK_B},
	"B": {true, keybd_event.VK_B},

	"c": {false, keybd_event.VK_C},
	"C": {true, keybd_event.VK_C},

	"d": {false, keybd_event.VK_D},
	"D": {true, keybd_event.VK_D},

	"e": {false, keybd_event.VK_E},
	"E": {true, keybd_event.VK_E},

	"f": {false, keybd_event.VK_F},
	"F": {true, keybd_event.VK_F},

	"g": {false, keybd_event.VK_G},
	"G": {true, keybd_event.VK_G},

	"h": {false, keybd_event.VK_H},
	"H": {true, keybd_event.VK_H},

	"i": {false, keybd_event.VK_I},
	"I": {true, keybd_event.VK_I},

	"j": {false, keybd_event.VK_J},
	"J": {true, keybd_event.VK_J},

	"k": {false, keybd_event.VK_K},
	"K": {true, keybd_event.VK_K},

	"l": {false, keybd_event.VK_L},
	"L": {true, keybd_event.VK_L},

	"m": {false, keybd_event.VK_M},
	"M": {true, keybd_event.VK_M},

	"n": {false, keybd_event.VK_N},
	"N": {true, keybd_event.VK_N},

	"o": {false, keybd_event.VK_O},
	"O": {true, keybd_event.VK_O},

	"p": {false, keybd_event.VK_P},
	"P": {true, keybd_event.VK_P},

	"q": {false, keybd_event.VK_Q},
	"Q": {true, keybd_event.VK_Q},

	"r": {false, keybd_event.VK_R},
	"R": {true, keybd_event.VK_R},

	"s": {false, keybd_event.VK_S},
	"S": {true, keybd_event.VK_S},

	"t": {false, keybd_event.VK_T},
	"T": {true, keybd_event.VK_T},

	"u": {false, keybd_event.VK_U},
	"U": {true, keybd_event.VK_U},

	"v": {false, keybd_event.VK_V},
	"V": {true, keybd_event.VK_V},

	"w": {false, keybd_event.VK_W},
	"W": {true, keybd_event.VK_W},

	"x": {false, keybd_event.VK_X},
	"X": {true, keybd_event.VK_X},

	"y": {false, keybd_event.VK_Y},
	"Y": {true, keybd_event.VK_Y},

	"z": {false, keybd_event.VK_Z},
	"Z": {true, keybd_event.VK_Z},

	"0": {false, keybd_event.VK_0},
	")": {true, keybd_event.VK_0},

	"1": {false, keybd_event.VK_1},
	"!": {true, keybd_event.VK_1},

	"2": {false, keybd_event.VK_2},
	"@": {true, keybd_event.VK_2},

	"3": {false, keybd_event.VK_3},
	"#": {true, keybd_event.VK_3},

	"4": {false, keybd_event.VK_4},
	"$": {true, keybd_event.VK_4},

	"5": {false, keybd_event.VK_5},
	"%": {true, keybd_event.VK_5},

	"6": {false, keybd_event.VK_6},
	"^": {true, keybd_event.VK_6},

	"7": {false, keybd_event.VK_7},
	"&": {true, keybd_event.VK_7},

	"8": {false, keybd_event.VK_8},
	"*": {true, keybd_event.VK_8},

	"9": {false, keybd_event.VK_9},
	"(": {true, keybd_event.VK_9},

	"-": {false, keybd_event.VK_MINUS},
	"_": {true, keybd_event.VK_MINUS},

	"=": {false, keybd_event.VK_EQUAL},
	"+": {true, keybd_event.VK_EQUAL},

	"[": {false, keybd_event.VK_LeftBracket},
	"{": {true, keybd_event.VK_LeftBracket},

	"]": {false, keybd_event.VK_RightBracket},
	"}": {true, keybd_event.VK_RightBracket},

	";": {false, keybd_event.VK_SEMICOLON},
	":": {true, keybd_event.VK_SEMICOLON},

	"'":  {false, keybd_event.VK_Quote},
	"\"": {true, keybd_event.VK_Quote},

	"\\": {false, keybd_event.VK_BACKSLASH},
	"|":  {true, keybd_event.VK_BACKSLASH},

	",": {false, keybd_event.VK_COMMA},
	"<": {true, keybd_event.VK_COMMA},

	".": {false, keybd_event.VK_Period},
	">": {true, keybd_event.VK_Period},

	"/": {false, keybd_event.VK_SLASH},
	"?": {true, keybd_event.VK_SLASH},

	"`": {false, keybd_event.VK_GRAVE},
	"~": {true, keybd_event.VK_GRAVE},

	" ": {false, keybd_event.VK_SPACE},
}
