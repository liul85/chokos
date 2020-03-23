package covtiser

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAddCorrectRoute(t *testing.T) {
	r := newRouter()
	handler := func(c *Context) {
		c.String(http.StatusOK, "index")
	}
	r.addRoute(http.MethodGet, "/index", handler)

	got, ok := r.handlers["GET-/index"]
	if !ok {
		t.Errorf("Expected index path handler with GET method, but didn't get")
	}

	sf1 := reflect.ValueOf(handler)
	sf2 := reflect.ValueOf(got)

	if sf1.Pointer() != sf2.Pointer() {
		t.Errorf("Expected handler, but didnt get")
	}
}
