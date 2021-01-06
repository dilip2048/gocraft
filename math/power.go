package main

import (
	"github.com/conversenow/utils"
	"github.com/nexmo-community/nexmo-go"
	"github.com/rs/zerolog/log"
	"net/http"
)

/* Program to test function power */
func main() {
	//http.HandleFunc("/v1/tracker/feedback", func(w http.ResponseWriter, r *http.Request) {
	//	params := r.URL.Query()
	//	fmt.Println("SMS from " + params["msisdn"][0] + ": " + params["text"][0])
	//})
	//
	//http.ListenAndServe(":3000", nil)

	//todo: test rating starts
	//rating := "Hello"
	//feedback, err := strconv.ParseFloat(rating, 64)
	//if err != nil || (feedback < 0 && feedback > 5) {
	//	// if error occur, must be due to use send text sms instead of rating
	//	print("Successfully updated tracker with customer feedback")
	//}
	//if feedback > 0 && feedback <= 5 {
	//	println("Successfully updated tracker with customer rating")
	//}
	//todo: test raing end


	//todo: sms start
	subLogger := log.With().Str("module", "nexmo").Logger()

	auth := nexmo.NewAuthSet()
	auth.SetAPISecret("82b7a4c8", "62U2DdMk6rKtNxBg")

	client := nexmo.NewClient(http.DefaultClient, auth)

	sI, _, err := client.SMS.SendSMS(nexmo.SendSMSRequest{
		From: "12012108254",
		//To:   "15123332836", //skype
		To:"12013508101",
		//To:   "19293693281", //Manish
		Type: "text",
		Text: "5",
	})
	if err != nil {
		subLogger.Error().Err(err).Msg("Error sending sms")
	}
	if sI.Messages[0].ErrorText != "" {
		subLogger.Error().Msgf("SMS send failed due to %s", sI.Messages[0].ErrorText)
	}
	subLogger.Debug().RawJSON("sendSMSResponse", utils.ConvertToByteString(sI)).Msg("Feedback SMS details")
	////todo: sms ends



	//apiName := "PQO38FSL"
	//apiKey := "997qF6qW575HMf2U"
	//apiMode := "test"
	//
	//authorizenet.New(apiName,apiKey,true)
}
