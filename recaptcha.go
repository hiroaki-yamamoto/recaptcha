package recaptcha

const veryfyURL = "https://www.google.com/recaptcha/api/siteverify"

// Recaptcha is a structure to handle recaptcha
type Recaptcha struct {
	secKey string
}

// Check whether the response is from
// human (returns true) or not (returns false).
func (r Recaptcha) Check(remoteIP string, response string) (bool, err error) {

}

// New returns an instance of Recaptcha.
func New(secretKey string) Recaptcha {
	return Recaptcha{secKey: secretKey}
}
