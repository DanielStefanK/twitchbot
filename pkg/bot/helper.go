package bot

import "strings"

const prefix = "!"

func isCmd(msg string) bool {
	return strings.HasPrefix(msg, prefix)
}

func getCommand(msg string) string {
	i := strings.Index(msg, " ")
	if i < 0 {
		i = len(msg)
	}
	return msg[1:i]
}
