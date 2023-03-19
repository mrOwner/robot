package util

import "bytes"

const (
	auth   = "Authorization"
	bearer = "Bearer"
)

func BearerToken(token string) (string, string) {
	buf := bytes.Buffer{}
	buf.Grow(len(bearer) + len(token) + 1)
	buf.WriteString(bearer)
	buf.WriteRune(' ')
	buf.WriteString(token)

	return auth, buf.String()
}
