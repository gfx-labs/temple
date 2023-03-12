package prayer

type Raw struct {
	// Input specifies which template to use.
	Input string
	// Obj is the data that will be passed into the template.
	Obj any
	// Args will be accessible within the template as arg0, arg1, argN...
	Args []any
	// Formatter will be called to format and validate the output of the template.
	Formatter func([]byte) ([]byte, error)

	// Output defines the file name to place the output of this prayer.
	Output string
}

func (r *Raw) Template() string {
	return r.Input
}

func (r *Raw) Object() any {
	return r.Obj
}

func (r *Raw) Arguments() []any {
	return r.Args
}

func (r *Raw) Format(bytes []byte) ([]byte, error) {
	if r.Formatter == nil {
		return bytes, nil
	}

	return r.Formatter(bytes)
}

func (r *Raw) FileName() string {
	return r.Output
}

var _ Prayer = (*Raw)(nil)
