package log

import "fmt"

func WrapGreen(msg string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", msg)
}

func WrapRed(msg string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", msg)
}

func WrapYellow(msg string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", msg)
}

func WrapBlue(msg string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", msg)
}

func WrapMagenta(msg string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", msg)
}

func WrapCyan(msg string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", msg)
}

func WrapOrange(msg string) string {
	return fmt.Sprintf("\x1b[38;5;208m%s\x1b[0m", msg)
}

func WrapBrown(msg string) string {
	return fmt.Sprintf("\x1b[38;5;130m%s\x1b[0m", msg)
}
