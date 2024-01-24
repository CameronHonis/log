package log

import "fmt"

var WrapGreen = func(msg string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", msg)
}

var WrapRed = func(msg string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", msg)
}

var WrapYellow = func(msg string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", msg)
}

var WrapBlue = func(msg string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", msg)
}

var WrapMagenta = func(msg string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", msg)
}

var WrapCyan = func(msg string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", msg)
}

var WrapOrange = func(msg string) string {
	return fmt.Sprintf("\x1b[38;5;208m%s\x1b[0m", msg)
}

var WrapBrown = func(msg string) string {
	return fmt.Sprintf("\x1b[38;5;130m%s\x1b[0m", msg)
}
