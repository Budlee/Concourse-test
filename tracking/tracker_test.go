package tracking

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIPAddressIsLogged(t *testing.T) {
	remoteAddress := "1.2.3.4:9999"
	r, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.RemoteAddr = remoteAddress

	testRecorder := httptest.NewRecorder()
	handler := NewIPEchoHandler()
	handler.ServeHTTP(testRecorder, r)
	if testRecorder.Body.String() != remoteAddress {
		t.Errorf("Remote address expected: [%v], obtained: [%s]", remoteAddress, testRecorder.Body.String())
	}
}

func BenchmarkIPEcho(b *testing.B) {
	remoteAddress := "1.2.3.4:9999"
	r, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		b.Fatal(err)
	}
	r.RemoteAddr = remoteAddress
	for i := 0; i < b.N; i++ {
		echoHandler(httptest.NewRecorder(), r)
	}
}
