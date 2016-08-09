package secret

import (
	"fmt"
	"strings"
)

var secrets = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code int) (handshake []string) {
	if code >= 32 {
		code = code % 32
	}

	if code <= 0 {
		return nil
	}

	converted := fmt.Sprintf("%b", code)
	var res []string
	sixteen := false
	for i := len(converted) - 1; i > -1; i-- {
		res = append(res, fmt.Sprintf("%c", converted[i]))
	}

	rev := strings.Join(res, "")

	if len(rev) == 5 {
		rev = rev[0:4]
		sixteen = true
	}

	for i, c := range rev {
		if c == 49 {
			handshake = append(handshake, secrets[i])
		}
	}
	if sixteen {
		for i, j := 0, len(handshake)-1; i < j; i, j = i+1, j-1 {
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}
	}
	return
}
