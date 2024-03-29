package async

import "errors"

var (
	ErrInvalidReply   = errors.New("invalid response")
	ErrInvalidMessage = errors.New("invalid request")

	ErrInvalidMessageDataType     = errors.New("invalid message data type")
	ErrResponseTypeNotImplemented = errors.New("response type not implemented")

	ErrFailedToEncodeMessage = errors.New("failed to encode request")
)
