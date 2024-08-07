module github.com/ops-center/james-go-client

go 1.22

toolchain go1.22.4

require (
	git.sr.ht/~rockorager/go-jmap v0.4.6
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/pkg/errors v0.9.1
	github.com/searchlight/james-go-client v0.0.0-20231113195825-c235670364db
	github.com/stretchr/testify v1.8.4
	golang.org/x/oauth2 v0.20.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace git.sr.ht/~rockorager/go-jmap => github.com/ops-center/go-jmap v0.4.7-0.20240424042927-c18631092ce1
