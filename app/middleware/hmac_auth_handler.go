package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const MaxRequestAgeSeconds = 300
const HmacAuthScheme = "APIAuth-HMAC-SHA256"

// For the sake of this test and for simplicity
// api credentaials are defined here
// otherwise they would be in a database
const (
    user = "test"
    secret = "test"
)

func HandleAuthenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.Request.Header["Authorization"]
        requestTimeStamp := c.GetHeader("Date")

        if len(authHeader) == 0 || requestTimeStamp == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        authorization := strings.Fields(authHeader[0])
        if authorization[0] != HmacAuthScheme {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        signature := authorization[1]
        if signature == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        authValues := strings.Split(authorization[1], ":")
        if len(authValues) < 2 {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        appId := authValues[0]
        authBase64Signature := authValues[1]

        if appId != user || !isValidRequest(c.Request, authBase64Signature, requestTimeStamp, secret) || !isValidDate(requestTimeStamp) {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        c.Next()
    }
}

func isValidDate(dateInHeader string) bool {
    loc, _ := time.LoadLocation("GMT")
    currentDate := time.Now().In(loc)

    parsedRequestDate, parseError := time.Parse(time.RFC1123, dateInHeader)
    if parseError != nil {
        return false
    }

    if parsedRequestDate.Before(currentDate) && parsedRequestDate.After(currentDate.Add(-MaxRequestAgeSeconds * time.Second)) {
        return true
    }

    return false
        
}

func isValidRequest(request *http.Request, incomingBase64Signature string, requestTimeStamp string, secretKey string) bool {
    if secretKey == "" {
        return false
    }
    uri := request.URL.EscapedPath()

    body := request.Body
    bodyBytes, _ := ioutil.ReadAll(body)

    requestContentBase64String := ""
    if len(bodyBytes) > 0 {
        // Hack -> can't read the request body more than once
        // if middleware reads it first it sets it empty
        // and later using .ShouldBind will give me EOF error
        // this fixes this
        request.Body = io.NopCloser(bytes.NewReader(bodyBytes))

        bodyHash := sha256.New()
        bodyHash.Write(bodyBytes)

        bodyHashString := hex.EncodeToString(bodyHash.Sum(nil))
        requestContentBase64String = bodyHashString
    }

    rawData := fmt.Sprintf("application/json,%s,%s,%s", requestContentBase64String, uri, requestTimeStamp)
    hash := hmac.New(sha256.New, []byte(secretKey))
    hash.Write([]byte(rawData))

    expectedHmac := base64.StdEncoding.EncodeToString(hash.Sum(nil))
    if expectedHmac == incomingBase64Signature {
        return true
    }

    return false
}

