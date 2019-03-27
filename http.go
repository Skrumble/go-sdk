package skrumble 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const Version = "3.0"
const userAgent = "skrumble-go/" + Version

// Base URL 
var BaseURL = "https://api.skrumble.com"

// @TODO: implement error logic & structure 
type skrumbleError struct {
	Code 		int	`json: "code"`
	Message 	string	`json:"message"`
	MoreInfo 	string	`json:"more_info"`
	Status 		int 	`json: "status"`
}

func newClient(clientName string) *Client {
	// @TODO: implement new client logic here 
}
