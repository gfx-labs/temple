package prayer

type Prayer interface {
	Template() string
	Object() any
	Arguments() []any

	Format([]byte) ([]byte, error)

	FileName() string
}
