package prayer

import (
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

func (r *Rust) Format(bytes []byte) ([]byte, error) {
	fmt := exec.Command("rustfmt")
	in, err := fmt.StdinPipe()
	if err != nil {
		return nil, err
	}
	_, err = in.Write(bytes)
	if err != nil {
		return nil, err
	}
	err = in.Close()
	if err != nil {
		return nil, err
	}
	return fmt.Output()
}

func (r *Rust) FileName() string {
	return r.Output
}

var _ Prayer = (*Rust)(nil)
