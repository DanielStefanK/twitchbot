package bot

import (
	"strings"
)

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

func getParams(msg string) []string {
	splitted := strings.Split(msg, " ")
	if len(splitted) == 1 {
		return make([]string, 0, 0)
	}
	return splitted[1:]
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.ToLower(a) == strings.ToLower(e) {
			return true
		}
	}
	return false
}
