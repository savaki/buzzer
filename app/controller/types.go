package controller

import "encoding/xml"

type Form struct {
	CallSid       string
	Caller        string
	Called        string
	AccountSid    string
	CallStatus    string
	CallerCity    string
	CallerState   string
	CallerZip     string
	CallerCountry string
	CalledCity    string
	CalledState   string
	CalledZip     string
	CalledCountry string
}

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
}
