package error

const (
	TokenInvalid = "token invalid"
)

type TokenDecoderError struct {
	Err string
}

func (err *TokenDecoderError) Error() string {
	return err.Err
}
