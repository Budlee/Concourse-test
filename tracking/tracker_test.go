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
	handler := NewTrackingHandler()
	handler.ServeHTTP(testRecorder, r)
	if testRecorder.Body.String() != remoteAddress {
		t.Errorf("Remote address expected: [%v], obtained: [%s]", remoteAddress, testRecorder.Body.String())
	}

}
