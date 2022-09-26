package main

import (
	"crypto/tls"
	"log"
	"time"

	"misc/go-soap-client/gen"

	"github.com/hooklift/gowsdl/soap"
)

func main() {
	req := &gen.CheckCardByLaser{
		PID:       "3212542745619",
		FirstName: "อร",
		LastName:  "ใจดี",
		BirthDay:  "25300101",
		Laser:     "JT1232145632",
	}
	log.Printf("pid = %s", req.PID)
	log.Printf("firstname = %s", req.FirstName)
	log.Printf("lastname = %s", req.LastName)
	log.Printf("birthday = %s", req.BirthDay)
	log.Printf("laser = %s", req.Laser)
	log.Println("")
	client := soap.NewClient(
		"http://10.217.141.81:8090/DOPAMOCKUP/services/CheckCardBankServiceSoap",
		soap.WithTimeout(time.Second*5),
		//soap.WithBasicAuth("usr", "psw"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
	service := gen.NewCheckCardServiceSoap(client)
	reply, err := service.CheckCardByLaser(req)
	if err != nil {
		log.Fatalf("could't get trade prices: %v", err)
	}
	log.Printf("code = %d", reply.CheckCardByLaserResult.Code)
	log.Printf("desc = %s", reply.CheckCardByLaserResult.Desc)
	log.Printf("isError = %v", reply.CheckCardByLaserResult.IsError)
	log.Printf("errorMessage = %s", reply.CheckCardByLaserResult.ErrorMessage)

	req_ := &gen.CheckCardByCID{
		PID:    "4440900001894",
		ChipNo: "5364260c1b3033f4",
		Bp1no:  "300341258011",
	}
	log.Println("+++++++++++++++++++++++++++++++++++++")
	log.Printf("pid = %s", req_.PID)
	log.Printf("Chip No = %s", req_.ChipNo)
	log.Printf("Bp1no = %s", req_.Bp1no)
	log.Println("")
	reply_, err := service.CheckCardByCID(req_)
	if err != nil {
		log.Fatalf("could't get trade prices: %v", err)
	}
	log.Printf("code = %d", reply_.CheckCardByCIDResult.Code)
	log.Printf("desc = %s", reply_.CheckCardByCIDResult.Desc)
	log.Printf("isError = %v", reply_.CheckCardByCIDResult.IsError)
	log.Printf("errorMessage = %s", reply_.CheckCardByCIDResult.ErrorMessage)
}
