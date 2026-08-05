package pagination

const (
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100
)

type Request struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// Normalize applies default values and limits.
func (r *Request) Normalize() {

	if r.Page < 1 {
		r.Page = DefaultPage
	}

	if r.Limit < 1 {
		r.Limit = DefaultLimit
	}

	if r.Limit > MaxLimit {
		r.Limit = MaxLimit
	}
}

// Offset returns SQL offset.
func (r Request) Offset() int {
	return (r.Page - 1) * r.Limit
}
