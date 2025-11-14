package server

import (
	"html/template"
	"time"
)

// reflection contains all information we can discover about the incoming request.
type reflection struct {
	Timestamp        time.Time           `json:"timestamp"`
	Method           string              `json:"method"`
	Proto            string              `json:"proto"`
	Scheme           string              `json:"scheme"`
	Host             string              `json:"host"`
	RequestURI       string              `json:"request_uri"`
	RemoteAddr       string              `json:"remote_addr"`
	RemoteIP         string              `json:"remote_ip"`
	RemotePort       string              `json:"remote_port"`
	TLS              *tlsDetails         `json:"tls,omitempty"`
	Headers          map[string][]string `json:"headers"`
	Query            map[string][]string `json:"query"`
	Cookies          []cookieDetails     `json:"cookies,omitempty"`
	ContentLength    int64               `json:"content_length"`
	TransferEncoding []string            `json:"transfer_encoding,omitempty"`
	BodyPreview      string              `json:"body_preview,omitempty"`
	ClientData       map[string]any      `json:"client_data,omitempty"`
}

type tlsDetails struct {
	Version     string `json:"version"`
	CipherSuite string `json:"cipher_suite"`
	ServerName  string `json:"server_name"`
	Negotiated  string `json:"alpn"`
}

type cookieDetails struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type keyValues struct {
	Key    string
	Values []string
}

type pageData struct {
	Reflection    reflection
	Headers       []keyValues
	Query         []keyValues
	ClientJSON    string
	HasClientData bool
	StatusMessage string
	StatusVariant string
	ClientScript  template.JS
}
