package basic

type Filter struct {
	Limit       *int
	Offset      *int
	Order       *string
	Deleted     bool
	WithDeleted bool
}
