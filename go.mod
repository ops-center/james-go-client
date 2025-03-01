module go.opscenter.dev/james-go-client

go 1.22.1

toolchain go1.23.1

require (
	git.sr.ht/~rockorager/go-jmap v0.5.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/oauth2 v0.23.0
	k8s.io/apimachinery v0.30.2
	xorm.io/xorm v1.3.9
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20240502163921-fe8a2dddb1d0 // indirect
)

replace git.sr.ht/~rockorager/go-jmap => github.com/ops-center/go-jmap v0.5.1-0.20250107173053-f306fc3d04c3
