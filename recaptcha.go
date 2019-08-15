package recaptcha

type Recaptcha struct {
	secKey string
}

func New(secretKey string) Recaptcha {
	return Recaptcha{secretKey}
}
