package prayer

import (
	"os/exec"
	"testing"
)

func TestRustFormat(t *testing.T) {
	if _, err := exec.LookPath("rustfmt"); err != nil {
		t.Skip("rustfmt not installed")
	}
	out, err := (&Rust{}).Format([]byte("async fn f(){let x=1;}"))
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "async fn f() {\n    let x = 1;\n}\n" {
		t.Fatalf("got %q", out)
	}
	if _, err := (&Rust{}).Format([]byte("fn main( {")); err == nil {
		t.Fatal("expected error")
	}
}
