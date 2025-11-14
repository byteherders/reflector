package server

import (
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"
)

func schemeFromRequest(r *http.Request) string {
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		return proto
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func clientIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		parts := strings.Split(forwarded, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	if realIP := r.Header.Get("X-Real-Ip"); realIP != "" {
		return realIP
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func clientPort(r *http.Request) string {
	_, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return port
}

func cloneHeader(h http.Header) map[string][]string {
	out := make(map[string][]string, len(h))
	var keys []string
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		out[k] = append([]string(nil), h[k]...)
	}
	return out
}

func queryValues(r *http.Request) map[string][]string {
	if len(r.URL.Query()) == 0 {
		return nil
	}
	out := make(map[string][]string, len(r.URL.Query()))
	keys := make([]string, 0, len(r.URL.Query()))
	for k := range r.URL.Query() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		out[k] = append([]string(nil), r.URL.Query()[k]...)
	}
	return out
}

func cookieValues(r *http.Request) []cookieDetails {
	cookies := r.Cookies()
	if len(cookies) == 0 {
		return nil
	}
	out := make([]cookieDetails, 0, len(cookies))
	for _, c := range cookies {
		out = append(out, cookieDetails{Name: c.Name, Value: c.Value})
	}
	return out
}

func tlsFromRequest(r *http.Request) *tlsDetails {
	if r.TLS == nil {
		return nil
	}
	details := &tlsDetails{
		CipherSuite: tlsCipherSuiteName(r.TLS.CipherSuite),
		Version:     tlsVersionName(r.TLS.Version),
		ServerName:  r.TLS.ServerName,
		Negotiated:  r.TLS.NegotiatedProtocol,
	}
	return details
}

// tlsCipherSuiteName mirrors the strings exported in crypto/tls for readability.
func tlsCipherSuiteName(id uint16) string {
	switch id {
	case 0x1301:
		return "TLS_AES_128_GCM_SHA256"
	case 0x1302:
		return "TLS_AES_256_GCM_SHA384"
	case 0x1303:
		return "TLS_CHACHA20_POLY1305_SHA256"
	case 0xc02f:
		return "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
	case 0xc030:
		return "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	case 0xc02b:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	case 0xc02c:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	case 0x009c:
		return "TLS_RSA_WITH_AES_128_GCM_SHA256"
	case 0x009d:
		return "TLS_RSA_WITH_AES_256_GCM_SHA384"
	default:
		return fmt.Sprintf("0x%04x", id)
	}
}

func tlsVersionName(v uint16) string {
	switch v {
	case 0x0301:
		return "TLS1.0"
	case 0x0302:
		return "TLS1.1"
	case 0x0303:
		return "TLS1.2"
	case 0x0304:
		return "TLS1.3"
	default:
		return fmt.Sprintf("0x%04x", v)
	}
}

func mapToPairs(m map[string][]string) []keyValues {
	if len(m) == 0 {
		return nil
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]keyValues, 0, len(keys))
	for _, k := range keys {
		out = append(out, keyValues{Key: k, Values: append([]string(nil), m[k]...)})
	}
	return out
}
