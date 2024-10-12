package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HTTPUtil struct {
	handlers map[string]map[string]http.HandlerFunc
}

func NewHTTPUtil() *HTTPUtil {
	return &HTTPUtil{
		handlers: map[string]map[string]http.HandlerFunc{},
	}
}

func (h *HTTPUtil) addHandle(path string, handler http.HandlerFunc, method string) {
	if _, ok := h.handlers[path]; !ok {
		h.handlers[path] = map[string]http.HandlerFunc{}
	}

	h.handlers[path][method] = handler
}

func (h *HTTPUtil) Get(path string, handler http.HandlerFunc) {
	h.addHandle(path, handler, http.MethodGet)
}

func (h *HTTPUtil) Post(path string, handler http.HandlerFunc) {
	h.addHandle(path, handler, http.MethodPost)
}

func (h *HTTPUtil) GetHandler() http.Handler {
	mux := http.NewServeMux()

	for path, methods := range h.handlers {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			method := r.Method

			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			log.Printf("[APIGW Handler] path=%s, method=%s, req=%s", path, method, body)

			handler, ok := methods[method]
			if !ok {
				http.Error(w, fmt.Sprintf("method type is not supported, method=%v", method), http.StatusNotFound)
				return
			}

			handler(w, r)
		})

	}

	return mux
}
