package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
    baseUrl = "http://localhost:5123"

    GetAllCustomersEndpoint = "/api/v1/customer/getAll"
    GetCustomerByIdEndpoint = "/api/v1/customer/get/%s"
    CreateNewCustomerEndpoint = "/api/v1/customer/create"
)

type clientImpl struct {
    request http.Request
}

func (c clientImpl) hmacAuthDelegation() {
    clientId := "test"
    secret := "test"

    requestUri := c.request.URL.RequestURI()

    loc, _ := time.LoadLocation("GMT")
    date := time.Now().In(loc)
    requestDate := date.Format(http.TimeFormat)

    requestContentBase64String := ""

    body := c.request.Body
    if body != nil {
        bodyBytes, _ := ioutil.ReadAll(body)
        if len(bodyBytes) > 0 {
            bodyHash := sha256.New()
            bodyHash.Write(bodyBytes)

            bodyHashString := hex.EncodeToString(bodyHash.Sum(nil))
            requestContentBase64String = bodyHashString
        }
    }

    rawSignatureData := fmt.Sprintf("application/json,%s,%s,%v", requestContentBase64String, requestUri, requestDate)
    hash := hmac.New(sha256.New, []byte(secret))
    hash.Write([]byte(rawSignatureData))

    sha256 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
    hmacKey := fmt.Sprintf("APIAuth-HMAC-SHA256 %s:%s", clientId, sha256)

    c.request.Header.Add("Authorization", hmacKey)
    c.request.Header.Add("Date", requestDate)
}

func (c clientImpl) newRequest(method string, endpoint string, body map[string]string) string {

    bodyJson, jsonErr := json.Marshal(body)
    if jsonErr != nil {
        return jsonErr.Error()
    }

    req, reqErr := http.NewRequest(method, baseUrl + endpoint, bytes.NewBuffer(bodyJson))
    if reqErr != nil {
        return reqErr.Error()
    }

    req.Header.Set("Content-Type","application/json")
    req.Header.Add("Accept","/")

    c.request = *req
    c.hmacAuthDelegation()

    // Hack -> can't read the request body more than once
    // After auth body is set to empty for some reason
    // this fixes this
    c.request.Body = io.NopCloser(bytes.NewReader(bodyJson))
    
    client := &http.Client{Timeout: 1000 * time.Second}
    response, err := client.Do(&c.request)
    if err != nil {
        return fmt.Sprintf("Error: %s", err.Error())
    }
    defer response.Body.Close()

    responseBody, _ := ioutil.ReadAll(response.Body)
    return string(responseBody)
}

func main() {

    response := clientImpl{}.newRequest(http.MethodGet, GetAllCustomersEndpoint, nil)
    //response := clientImpl{}.newRequest(http.MethodGet, fmt.Sprintf(GetCustomerByIdEndpoint, "19"), nil)

    /*requestBody := map[string]string {
        "companyName": "test company",
        "firstName": "test example",
        
    }*/

    //response := clientImpl{}.newRequest(http.MethodPost, CreateNewCustomerEndpoint, requestBody)
    fmt.Println("GetAllCustomers response: ", response)
}


