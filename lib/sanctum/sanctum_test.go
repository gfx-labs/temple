package sanctum

import (
	"testing"

	"github.com/gfx-labs/temple/lib/prayer"
	"github.com/spf13/afero"
)

func TestFormatFailureKeepsPreviousOutput(t *testing.T) {
	fs := afero.NewMemMapFs()
	prev := []byte("package out\n")
	if err := afero.WriteFile(fs, "out/a.go", prev, 0o644); err != nil {
		t.Fatal(err)
	}
	s := NewWithFS(fs)
	s.RegisterTemplate("bad", "func (")
	s.Prepare(&prayer.Go{Input: "bad", Output: "out/a.go"})
	if err := s.Pray(); err == nil {
		t.Fatal("expected format error")
	}
	got, _ := afero.ReadFile(fs, "out/a.go")
	if string(got) != string(prev) {
		t.Fatalf("output was modified: %q", got)
	}
}

func TestTemplateEndingInBrace(t *testing.T) {
	fs := afero.NewMemMapFs()
	s := NewWithFS(fs)
	s.RegisterTemplate("t", "struct {")
	s.Prepare(&prayer.Raw{Input: "t", Output: "x"})
	if err := s.Pray(); err != nil {
		t.Fatal(err)
	}
	got, _ := afero.ReadFile(fs, "x")
	if string(got) != "struct {" {
		t.Fatalf("got %q", got)
	}
}

func TestParseErrorReturnsError(t *testing.T) {
	s := NewWithFS(afero.NewMemMapFs())
	s.RegisterTemplate("broken", "{{ if }")
	s.Prepare(&prayer.Raw{Input: "broken", Output: "x"})
	if err := s.Pray(); err == nil {
		t.Fatal("expected parse error")
	}
}

func TestFuncsNotShared(t *testing.T) {
	a := NewWithFS(afero.NewMemMapFs())
	a.RegisterFunc("only_a", func() string { return "" })
	b := NewWithFS(afero.NewMemMapFs())
	if _, ok := b.fm["only_a"]; ok {
		t.Fatal("func registered on one sanctum leaked to another")
	}
}

func TestRegisterMissingTemplateErrors(t *testing.T) {
	s := NewWithFS(afero.NewMemMapFs())
	if err := s.RegisterTemplateFile("nope.tmpl"); err == nil {
		t.Fatal("expected error for missing file")
	}
	if err := s.RegisterTemplateDir("nope"); err == nil {
		t.Fatal("expected error for missing dir")
	}
}

func TestArgs(t *testing.T) {
	fs := afero.NewMemMapFs()
	s := NewWithFS(fs)
	s.RegisterTemplate("t", "{{arg0}}{{arg1}}")
	s.Prepare(&prayer.Raw{Input: "t", Args: []any{"a", "b"}, Output: "x"})
	if err := s.Pray(); err != nil {
		t.Fatal(err)
	}
	got, _ := afero.ReadFile(fs, "x")
	if string(got) != "ab" {
		t.Fatalf("got %q", got)
	}
}
