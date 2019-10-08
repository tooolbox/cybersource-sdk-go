package main

//go:generate rm -rf pkg/nvp
//go:generate rm -rf pkg/soap
//go:generate gobin -m -run github.com/tooolbox/gowsdl/cmd/gowsdl@v0.3.2 -o service.go -p soap https://ics2wstest.ic3.com/commerce/1.x/transactionProcessor/CyberSourceTransaction_1.120.wsdl
//go:generate gobin -m -run github.com/tooolbox/gowsdl/cmd/gowsdl@v0.3.2 -o service.go -p nvp https://ics2wstest.ic3.com/commerce/1.x/transactionProcessor/CyberSourceTransaction_NVP_1.120.wsdl
//go:generate mkdir -p pkg
//go:generate mv soap pkg/soap
//go:generate mv nvp pkg/nvp

import (
	"github.com/tooolbox/cybersource-sdk-go/cmd"
)

func main() {
	cmd.Run()
}
