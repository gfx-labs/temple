package sanctum

import "os/exec"

type RustPrayer struct {
	// Input specifies which template to use.
	Input string
	// Obj is the data that will be passed into the template.
	Obj any
	// Args will be accessible within the template as arg0, arg1, argN...
	Args []any

	// Output defines the file path and name to place the output of this prayer.
	Output string
}

func (r *RustPrayer) Template() string {
	return r.Input
}

func (r *RustPrayer) Object() any {
	return r.Obj
}

func (r *RustPrayer) Arguments() []any {
	return r.Args
}

func (r *RustPrayer) Format(bytes []byte) ([]byte, error) {
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

func (r *RustPrayer) FileName() string {
	return r.Output
}

var _ Prayer = (*RustPrayer)(nil)
