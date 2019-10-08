package cmd

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/hooklift/gowsdl/soap"
	"github.com/k0kubun/pp"
	"github.com/motemen/go-loghttp"
	"github.com/satori/go.uuid"
	"github.com/tooolbox/cybersource-sdk-go/pkg/helper"
	processor "github.com/tooolbox/cybersource-sdk-go/pkg/soap"
)

func Run() {

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Need a merchantId and transactionKey!")
	}

	merchantId, transactionKey := args[0], args[1]

	log.Println("MerchantId: " + merchantId)

	httpClient := &http.Client{
		Transport: &loghttp.Transport{
			LogRequest: func(req *http.Request) {
				dump, _ := httputil.DumpRequestOut(req, true)
				log.Printf("Request: %s\n", string(dump))
			},
			LogResponse: func(resp *http.Response) {
				dump, _ := httputil.DumpResponse(resp, true)
				log.Printf("Response: %s\n", string(dump))
			},
		},
	}

	client := soap.NewClient("https://ics2wstest.ic3.com/commerce/1.x/transactionProcessor/", soap.WithHTTPClient(httpClient))
	client.AddHeader(helper.Header{Value: soap.NewWSSSecurityHeader(merchantId, transactionKey, "", "1")})

	service := processor.NewITransactionProcessor(client)

	req := &processor.RequestMessage{
		MerchantReferenceCode: uuid.NewV4().String(),
		MerchantID:            merchantId,
		CcAuthService:         &processor.CCAuthService{Run: true},
		CcCaptureService:      &processor.CCCaptureService{Run: true},
		BillTo: &processor.BillTo{
			FirstName:  "John",
			LastName:   "Doe",
			Street1:    "1295 Charleston Road",
			City:       "Mountain View",
			State:      "CA",
			PostalCode: "94043",
			Country:    "US",
			Email:      "null@cybersource.com",
			IpAddress:  "10.7.111.111",
		},
		Card: &processor.Card{
			AccountNumber:   "4111111111111111",
			ExpirationMonth: 12,
			ExpirationYear:  2020,
		},
		PurchaseTotals: &processor.PurchaseTotals{
			Currency:         "USD",
			GrandTotalAmount: Amount("90.01"),
		},
		TransactionLocalDateTime: time.Now().UTC().Format(helper.CyberSourceDateTimeFormat),
	}

	reply, err := service.RunTransaction(req)
	if err != nil {
		log.Fatalf("Error running transaction: %v", err)
	}

	pp.Printf("Reply:\n%v", reply)
}

func Amount(str string) *processor.Amount {
	amt := processor.Amount(str)
	return &amt
}
