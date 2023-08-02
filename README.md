# Start
Run `docker compose -f docker-compose-redis.yml up` to start redis cache locally. If cache can not be reached the errors will be overlooked here and request will be make against ERPLY API.

Once redis is running start the api by running `go run .`

## Api authorization
For the sake of simplicity user and user secret are hardcoded

### Postman scipt
```
var AppId = "test";
var APIKey = "test";
var currentDate = new Date();
var requestURI = "/" + pm.environment.values.substitute(pm.request.url, null, false).toString().split('/').splice(3).join('/');
requestURI = requestURI.split("?")[0];
console.log(requestURI)
var requestContentBase64String = "";
if (pm.request.body.raw) {
    var md5 = CryptoJS.SHA256(pm.request.body.toString());
    requestContentBase64String = md5.toString();
}

console.log(requestContentBase64String)
var signatureRawData  = `application/json,${requestContentBase64String},${requestURI},${currentDate.toUTCString()}`; //check
var signature = CryptoJS.enc.Utf8.parse(signatureRawData);
var secretByteArray = CryptoJS.enc.Base64.parse(APIKey);
var signatureBytes = CryptoJS.HmacSHA256(signature,APIKey);
var requestSignatureBase64String = CryptoJS.enc.Base64.stringify(signatureBytes);
var hmacKey = "APIAuth-HMAC-SHA256 " + AppId + ":" + requestSignatureBase64String;
pm.variables.set("hmacDate", currentDate.toUTCString());
pm.variables.set("hmacKey", hmacKey)
```
**Header values**

**Authorization**: {{hmacKey}}

**Date**: {{hmacDate}}

# Swagger doc
Swagger URL is localhost:5123/swagger/index.html

# Unit tests
To run unit tests run `go test ./...`
