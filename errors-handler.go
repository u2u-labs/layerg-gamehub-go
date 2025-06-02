package layerggamehub

import (
	"errors"
	"net"
	"net/http"
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

func shouldRetry(resp *http.Response, err error) bool {
	if err != nil {
		return isConnectionError(err)
	}

	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		return true
	}

	return false
}
