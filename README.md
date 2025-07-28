# Recaptcha Serverside Validation in Golang

[![CI]][CILink] [![MI]][MILink] [![TC]][TCLink]
[![GRC]][GRCLink]

[CI]: https://github.com/hiroaki-yamamoto/recaptcha/actions/workflows/test.yml/badge.svg
[CILink]: https://github.com/hiroaki-yamamoto/recaptcha/actions/workflows/test.yml
[MI]: https://qlty.sh/gh/hiroaki-yamamoto/projects/recaptcha/maintainability.svg
[MILink]: https://qlty.sh/gh/hiroaki-yamamoto/projects/recaptcha
[TC]: https://qlty.sh/gh/hiroaki-yamamoto/projects/recaptcha/coverage.svg
[TCLink]: https://qlty.sh/gh/hiroaki-yamamoto/projects/recaptcha
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
