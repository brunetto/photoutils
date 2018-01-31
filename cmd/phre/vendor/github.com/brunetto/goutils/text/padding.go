package text

import(
	"log"
	"strings"
)

// LeftPad returns the string padded filling remaining left spaces to `length` with `pad`.
func LeftPad( str, pad string, length int ) ( string ) {
	var repeat int = nPadRepeat( str, pad, length )
	return strings.Repeat(pad, repeat) + str
}

// RightPad returns the string padded filling remaining left spaces to `length` with `pad`.
func RightPad( str, pad string, length int ) ( string ) {
	var repeat int = nPadRepeat( str, pad, length )
	return str + strings.Repeat(pad, repeat)
}

func nPadRepeat( str, pad string, length int ) ( int ) {
	var repeat int
	if (length - len(str)) % len(pad) != 0 {
		log.Fatal("Can't pad ", str, " with ", pad, " to length ", length)
	} else {
		repeat = (length - len(str)) / len(pad)
	}
	if repeat < 0 {
		repeat = 0
	}
	return repeat
}
