package errors

type AppError struct {
	ErrorCode int
	Message   string
}

type SongError struct {
	ErrorCode int
	Message   string
}

type ReleaseVersionError struct {
	ErrorCode int
	Message   string
}

func (s *AppError) Error() string {
	return s.Message
}
func (s *SongError) Error() string {
	return s.Message
}
func (s *ReleaseVersionError) Error() string {
	return s.Message
}
