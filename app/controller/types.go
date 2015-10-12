package controller

import "encoding/xml"

type Form struct {
	Attempt       int
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
	XMLName  xml.Name  `xml:"Response"`
	Say      string    `xml:",omitempty"`
	Play     *Play     `xml:",omitempty"`
	Gather   *Gather   `xml:",omitempty"`
	Redirect *Redirect `xml:",omitempty"`
}

type Gather struct {
	XMLName xml.Name `xml:"Gather"`
	Action  string   `xml:"action,attr,omitempty"`
	Method  string   `xml:"method,attr,omitempty"`
	Say     string   `xml:",omitempty"`
}

type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Method  string   `xml:"method,attr,omitempty"`
	Action  string   `xml:",innerxml"`
}

type Play struct {
	XMLName xml.Name `xml:"Play"`
	Digits  string   `xml:"digits,attr,omitempty"`
	Url     string   `xml:",innerxml"`
}
