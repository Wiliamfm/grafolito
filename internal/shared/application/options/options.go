package options

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

const (
	issuer           = "jwt:issuer"
	audience         = "jwt:audience"
	secret           = "jwt:secret"
	expiresInMinutes = "jwt:expiresInMinutes"
)

type jwtOptions struct {
	Issuer           string
	Audience         string
	Secret           string
	ExpiresInMinutes int
}

var lock = &sync.Mutex{}
var options *jwtOptions

func JwtOptions() *jwtOptions {
	if options != nil {
		return options
	}

	lock.Lock()
	defer lock.Unlock()

	issuer, ok := os.LookupEnv(issuer)
	if !ok {
		panic(fmt.Sprintf("%s not set", issuer))
	}
	audience, ok := os.LookupEnv(audience)
	if !ok {
		panic(fmt.Sprintf("%s not set", audience))
	}
	secret, ok := os.LookupEnv(secret)
	if !ok {
		panic(fmt.Sprintf("%s not set", secret))
	}
	expires, ok := os.LookupEnv(expiresInMinutes)
	if !ok {
		panic(fmt.Sprintf("%s not set", expiresInMinutes))
	}
	expiresInMinutes, err := strconv.Atoi(expires)
	if err != nil {
		panic(err)
	}

	options := &jwtOptions{
		Issuer:           issuer,
		Audience:         audience,
		Secret:           secret,
		ExpiresInMinutes: expiresInMinutes,
	}

	return options
}
