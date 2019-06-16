package hello

// Msg returns a string with "Hello " + your message
func Msg(msg string) string {
	return "Hello " + msg
}

// World returns "Hello World"
func World() string {
	return Msg("world!")
}
