package server

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type Server struct {
	bodyCap int
	mux     *http.ServeMux
}

func New(bodyCap int) *Server {
	srv := &Server{bodyCap: bodyCap}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", srv.healthHandler)
	mux.HandleFunc("/", srv.reflectionHandler)
	mux.HandleFunc("/collect", srv.collectHandler)
	srv.mux = mux
	return srv
}

func (s *Server) Handler() http.Handler {
	return logRequests(s.mux)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, "ok")
}

func (s *Server) reflectionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := readRequestBody(r, s.bodyCap)
	if err != nil {
		log.Printf("read request body: %v", err)
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}
	s.renderResponse(w, r, body, nil)
}

func (s *Server) collectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	limit := s.bodyCap
	if limit <= 0 {
		limit = 1 << 14
	}

	body, err := readRequestBody(r, limit)
	if err != nil {
		log.Printf("read client payload: %v", err)
		http.Error(w, "failed to read client payload", http.StatusInternalServerError)
		return
	}

	var clientData map[string]any
	if len(body) > 0 {
		if err := json.Unmarshal(body, &clientData); err != nil {
			log.Printf("decode client payload: %v", err)
			http.Error(w, "invalid client payload", http.StatusBadRequest)
			return
		}
	}

	s.renderResponse(w, r, body, clientData)
}

func (s *Server) renderResponse(w http.ResponseWriter, r *http.Request, body []byte, clientData map[string]any) {
	data := reflection{
		Timestamp:        time.Now().UTC(),
		Method:           r.Method,
		Proto:            r.Proto,
		Scheme:           schemeFromRequest(r),
		Host:             r.Host,
		RequestURI:       r.RequestURI,
		RemoteAddr:       r.RemoteAddr,
		RemoteIP:         clientIP(r),
		RemotePort:       clientPort(r),
		Headers:          cloneHeader(r.Header),
		Query:            queryValues(r),
		Cookies:          cookieValues(r),
		ContentLength:    r.ContentLength,
		TransferEncoding: append([]string(nil), r.TransferEncoding...),
		TLS:              tlsFromRequest(r),
		ClientData:       clientData,
	}

	if len(body) > 0 {
		data.BodyPreview = string(body)
	}

	var clientJSON string
	if clientData != nil {
		if pretty, err := json.MarshalIndent(clientData, "", "  "); err == nil {
			clientJSON = string(pretty)
		}
	}

	statusMessage := "Collecting additional details from your browser..."
	statusVariant := "info"
	if clientData != nil {
		statusMessage = "Browser-supplied metadata is shown below."
		statusVariant = "success"
	}

	page := pageData{
		Reflection:    data,
		Headers:       mapToPairs(data.Headers),
		Query:         mapToPairs(data.Query),
		ClientJSON:    clientJSON,
		HasClientData: clientData != nil,
		StatusMessage: statusMessage,
		StatusVariant: statusVariant,
		ClientScript:  template.JS(clientCollectorScript),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := reflectionTemplate.Execute(w, page); err != nil {
		log.Printf("render response: %v", err)
	}
}

func readRequestBody(r *http.Request, limit int) ([]byte, error) {
	if r.Body == nil || r.Body == http.NoBody {
		return nil, nil
	}
	defer r.Body.Close()
	if limit <= 0 {
		return nil, nil
	}
	return io.ReadAll(io.LimitReader(r.Body, int64(limit)))
}
