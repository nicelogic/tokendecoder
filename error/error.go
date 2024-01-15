package error

const (
	TokenExpired        = "token expired"
	TokenInvalid        = "token invalid"
)

type TokenDecoderError struct {
	Err string
}

func (err TokenDecoderError) Error() string {
	return err.Err
}
