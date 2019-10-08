# CyberSource Go Client
This is the Go client for the [CyberSource SOAP Toolkit API](http://www.cybersource.com/developers/getting_started/integration_methods/soap_toolkit_api).

## Prerequisites

- Go 1.13 or above
   - [curl](http://php.net/manual/en/book.curl.php), [openssl](http://php.net/manual/en/book.openssl.php), [soap](http://php.net/manual/en/book.soap.php) extensions must be enabled
- A CyberSource account. You can create an evaluation account [here](http://www.cybersource.com/register/).
- A CyberSource transaction key.  Instructions on obtaining a transaction key can be found [here](http://www.cybersource.com/developers/integration_methods/simple_order_and_soap_toolkit_api/soap_api/html/wwhelp/wwhimpl/js/html/wwhelp.htm#href=Intro.04.3.html).

## Installation
```
go get github.com/tooolbox/cybersource-sdk-go
```
By default, the WSDL file for the client is for API version 1.120 (the latest when this package was updated). Available WSDL file URLs can be browsed at the following locations:

- [test](https://ics2wstest.ic3.com/commerce/1.x/transactionProcessor/)
- [live](https://ics2ws.ic3.com/commerce/1.x/transactionProcessor/)


## Getting Started
See `cmd/test.go` for an example of how to use the client.


## Documentation

For more information about CyberSource services, see <http://www.cybersource.com/developers/documentation>

For all other support needs, see <http://www.cybersource.com/support>
