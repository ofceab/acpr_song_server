package message

// For getting message
type OperationMessager interface {
	// Get Operation message
	MessagePrint() string
}

// Describe a message with successed
type SucessMessage struct {
	msg string
}

// For error message
type ErrorMessage struct {
	msg string
}

func (c *SucessMessage) MessagePrint() string {
	return (*c).msg
}

func (c *ErrorMessage) MessagePrint() string {
	return (*c).msg
}
