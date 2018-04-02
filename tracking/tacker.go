package tracking

import (
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.RemoteAddr))
}

func NewIPEchoHandler() http.HandlerFunc {
	return echoHandler
}
