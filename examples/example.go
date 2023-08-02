package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
    baseUrl = "http://localhost:5123"

    GetAllCustomersEndpoint = "/api/v1/customer/getAll"
    GetCustomerByIdEndpoint = "/api/v1/customer/%s"
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

func (c clientImpl) newRequest(method string, endpoint string, body io.Reader) string {
    req, _ := http.NewRequest(method, baseUrl + endpoint, body)

    c.request = *req
    c.hmacAuthDelegation()
    
    client := &http.Client{Timeout: 10 * time.Second}
    response, err := client.Do(&c.request)
    if err != nil {
        return err.Error()
    }

    responseBody, _ := ioutil.ReadAll(response.Body)
    return string(responseBody)
}

func main() {

    response := clientImpl{}.newRequest(http.MethodGet, GetAllCustomersEndpoint, nil)
    fmt.Println("GetAllCustomers response: ", response)

}


