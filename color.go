package color

func Ansi(ansicode string) func(string) string {
	return func(msg string) string {
		return "\033["+ansicode+"m"+msg+"\033[0m"
	}
}

var Red         = Ansi("31")
var Green       = Ansi("32")
var Yellow      = Ansi("33")
var Blue        = Ansi("34")
var Purple      = Ansi("35")
var Aqua        = Ansi("36")
var Lightgray   = Ansi("37")
var Gray        = Ansi("90")
var LightRed    = Ansi("91")
var LightGreen  = Ansi("92")
var LightYellow = Ansi("93")
var LightBlue   = Ansi("94")
var LightPurple = Ansi("95")
var LightAqua   = Ansi("96")
var White       = Ansi("97")
