package enums

import (
	"encoding/json"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	if HttpMethodPatch.String() != "patch" {
		t.Fatalf("got %s", HttpMethodPatch)
	}
	c, err := ParseColor("blue")
	if err != nil || c != ColorBlue {
		t.Fatalf("got %v %v", c, err)
	}
	if _, err := ParseColor("pink"); err == nil {
		t.Fatal("expected error")
	}
	if Color(9).String() != "Color(9)" {
		t.Fatalf("got %s", Color(9))
	}

	b, _ := json.Marshal(map[string]LogLevel{"lvl": LogLevelWarn})
	if string(b) != `{"lvl":"warn"}` {
		t.Fatalf("got %s", b)
	}
	var m map[string]LogLevel
	if err := json.Unmarshal(b, &m); err != nil || m["lvl"] != LogLevelWarn {
		t.Fatalf("got %v %v", m, err)
	}
}
