package prayer

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Rust struct {
	// Input specifies which template to use.
	Input string
	// Obj is the data that will be passed into the template.
	Obj any
	// Args will be accessible within the template as arg0, arg1, argN...
	Args []any

	// Output defines the file path and name to place the output of this prayer.
	Output string
	// Edition is passed to rustfmt. Defaults to 2021.
	Edition string
}

func (r *Rust) Template() string {
	return r.Input
}

func (r *Rust) Object() any {
	return r.Obj
}

func (r *Rust) Arguments() []any {
	return r.Args
}

func (r *Rust) Format(src []byte) ([]byte, error) {
	edition := r.Edition
	if edition == "" {
		edition = "2021"
	}
	cmd := exec.Command("rustfmt", "--emit", "stdout", "--edition", edition)
	cmd.Stdin = bytes.NewReader(src)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("rustfmt: %w: %s", err, stderr.String())
	}
	return out, nil
}

func (r *Rust) FileName() string {
	return r.Output
}

var _ Prayer = (*Rust)(nil)
