package router

import "errors"

var (
	ErrInvalidCommandArgsType     = errors.New("invalid command args type")
	ErrInvalidResponseReceived    = errors.New("invalid response received")
	ErrResponseTypeNotImplemented = errors.New("response type not implemented")
	ErrFailedToEncodeRequest      = errors.New("failed to encode request")
	ErrReceivedFromServer         = errors.New("received from server")
)
