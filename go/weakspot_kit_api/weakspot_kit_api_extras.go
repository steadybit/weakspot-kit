package weakspot_kit_api

func Ptr[T any](val T) *T {
	return &val
}
