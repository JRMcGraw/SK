package OAuth2

import (

	"net/http"
	"fmt"
	"net/url"
	"os"
	"bytes"
	"strconv"

	"github.com/codegangsta/negroni"
		
)

/*
SERVICE_CREDENTIAL is found by the following command:
	echo -n {CLIENT_ID}:{CLIENT_SECRET} | base64
  example:
	echo -n flydubai-service-client:iM2NrWVd | base64
*/

const(
	DEBUG = false
	CLIENT_ID = "flydubai-service-client"
	CLIENT_SECRET = "iM2NrWVd"
	SERVICE_CREDENTIAL = "Zmx5ZHViYWktc2VydmljZS1jbGllbnQ6aU0yTnJXVmQ="
)

var (
	SCOPES = []string{""}
)

func ParseTokenRequest(r *http.Request) string{
	token := r.Header.Get("Authorization")[7:]
	return token
}

func GetAuthenticationToken(r *http.Request) (token string, exists bool) {

	exists = true
	if len(r.Header.Get("Authorization")) == 0  {
		exists = false
		return 
	}

	token = ParseTokenRequest(r)
	return
}

func CheckAuthentication(r *http.Request) (statusCode int) {

	host_name := os.Getenv("UAA_HOST_NAME")
	host_url := "https://"+host_name+"/oauth/check_token"
	if (host_name == "") {
		statusCode = http.StatusNetworkAuthenticationRequired		//  511
		fmt.Println("Error: No UAA_HOST_NAME specified in Environment")
		return
	}

	token, exists := GetAuthenticationToken(r)
	if (!exists) {
//		statusCode = http.StatusUnauthorized						//  401
		statusCode = http.StatusNetworkAuthenticationRequired		//  511
		return
	}
	
	data := url.Values{}
	data.Set("token", token)

	client := &http.Client{}
	req, err := http.NewRequest("POST", host_url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Basic " + SERVICE_CREDENTIAL)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
//	fmt.Println("DBG-> resp: ", resp)
	statusCode = resp.StatusCode
    defer resp.Body.Close()

	if (statusCode == http.StatusOK) {
		fmt.Println("OAuth2_DBG-> Verified Request Authorization")
	} else {
		fmt.Println("OAuth2_DBG-> Verify Request Authorization Failed; Code: ", statusCode)
	}

	return
}

func IsAuthenticated() negroni.Handler  {
	au := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		statusCode := CheckAuthentication(r)
		if statusCode != 200{
			//Handle the different response codes appropriately
			w.WriteHeader(statusCode)
			fmt.Fprintf(w, "Authentication Status Code: " + strconv.Itoa(statusCode))
			return
		}

		next(w,r)
	}
	return negroni.HandlerFunc(au)
}
