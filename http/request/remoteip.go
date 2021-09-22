package request

import (
	"net/http"
	"strings"
)

var (
	DefaultScanForwardHeaderKey = []string{"X-Forwarded-For", "X-Real-IP"}
)

func GetRemoteIP(r *http.Request) string {
	var ip string
	for _, key := range DefaultScanForwardHeaderKey {
		value := r.Header.Get(key)

		if strings.Contains(value, ", ") {
			i := strings.Index(value, ", ")
			if i == -1 {
				i = len(value)
			}

			ip = value[:i]
			break
		}

		if value != "" {
			ip = value
			break
		}
	}

	if ip != "" {
		return ip
	}

	addr := strings.Split(r.RemoteAddr, ":")
	if len(addr) == 1 {
		return addr[0]
	}

	return strings.Join(addr[0:len(addr)-1], ":")
}