package errors

type InvalidateDeleteRequest struct {
	ms string
}

func (i *InvalidateDeleteRequest) Error() string {
	return i.ms
}
