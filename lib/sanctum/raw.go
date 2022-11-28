package sanctum

type RawPrayer struct {
	// Input specifies which template to use.
	Input string
	// Obj is the data that will be passed into the template.
	Obj any
	// Args will be accessible within the template as arg0, arg1, argN...
	Args []any

	// Output defines the file name to place the output of this prayer.
	Output string
}

func (r *RawPrayer) Template() string {
	return r.Input
}

func (r *RawPrayer) Object() any {
	return r.Obj
}

func (r *RawPrayer) Arguments() []any {
	return r.Args
}

func (r *RawPrayer) Format(bytes []byte) ([]byte, error) {
	return bytes, nil
}

func (r *RawPrayer) FileName() string {
	return r.Output
}

var _ Prayer = (*RawPrayer)(nil)
