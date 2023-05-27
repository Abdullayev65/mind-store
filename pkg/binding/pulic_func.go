package binding

func MapForm(ptr any, form map[string][]string) error {
	return mapForm(ptr, form)
}
func Validate(obj any) error {
	return validate(obj)
}
