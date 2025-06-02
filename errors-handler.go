package layerggamehub

import (
	"errors"
	"net"
	"net/url"
	"strings"
)

func isConnectionError(err error) bool {
	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		if _, ok := urlErr.Err.(net.Error); ok {
			return true
		}
	}

	if strings.Contains(err.Error(), "connection refused") ||
		strings.Contains(err.Error(), "connection reset") {
		return true
	}

	return false
}
