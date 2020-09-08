# H1-credentials-helper-go

## Providers

### Passport provider

Passport provider is based on passport file which can be generated using [h1-cli](https://github.com/hyperonecom/h1-cli).

#### Usage

```go
package main


import (
    "log"

    "github.com/kuskoman/h1-credentials-helper-go"
)

func main() {
    provider, err := h1.GetPassportCredentialsHelper("") // empty string means that the library should look for passport file in ~/.h1/passport.json
    // if you have this file in different location you can pass it to this function

    if err != nil {
        log.Panic(err)
    }

    token, err := provider.GetToken("exampleAudience")

    // [...]
}
```
