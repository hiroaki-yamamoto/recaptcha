# Recaptcha Serverside Validation in Golang

[![CI]][CILink] [![MI]][MILink] [![TC]][TCLink]
[![GRC]][GRCLink]

[CI]: https://circleci.com/gh/hiroaki-yamamoto/recaptcha.svg?style=svg
[CILink]: https://circleci.com/gh/hiroaki-yamamoto/recaptcha
[MI]: https://api.codeclimate.com/v1/badges/c89eb28bd7dd782716e4/maintainability
[MILink]: https://codeclimate.com/github/hiroaki-yamamoto/recaptcha/maintainability
[TC]: https://api.codeclimate.com/v1/badges/c89eb28bd7dd782716e4/test_coverage
[TCLink]: https://codeclimate.com/github/hiroaki-yamamoto/recaptcha/test_coverage
[GRC]: https://goreportcard.com/badge/github.com/hiroaki-yamamoto/recaptcha
[GRCLink]: https://goreportcard.com/report/github.com/hiroaki-yamamoto/recaptcha

# What this?
This is an implementation of server-side validation of recaptcha,
implementing in Golang.

# Why this?
First I found some implementations, but they don't support multiple instance.
Moreover, there's no test code!!. This is why I started coding this.

# How to use?

Using this script is very simple:

```go
import "fmt"
import reecap "github.com/hiroaki-yamamoto/recaptcha"

// CheckHuman checks recaptcha token.
func CheckHuman(cliAddr string, token string) bool {
  re := recap.New("[Your secret key]") // Create a new instance
  resp, err := re.Check(cliAddr, token) // Start the validation
  if err != nil {
    fmt.Printf(err)
    return false
  }
  return resp.Success
}
```

Note that `resp` has more information not only `Success` field. For detail,
check [recaptcha.go]

[recaptcha.go]: recaptcha.go

# Contribution

Posting Issue & PR is welcome. If you post PR, it's more appreciated, but
the corresponding test code of the code you changed is mandatory (In the
case that you changed the doc, you don't need to write the test code. the doc
is the doc, not code).
