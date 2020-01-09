
## Generate Thycotic SOAP client

1. Get gowsdl `go get github.com/hooklift/gowsdl/...`
2. Get Thycotic wdsl `curl https://secret.example.com/SecretServer/webservices/sswebservice.asmx?wsdl >thycotic.wsdl`
3. Generate code `gowsdl -o thycotic-generated.go -p thycotic thycotic.wsdl`

### Fixes in generated code
Add `type String string` to package to define missing type. 

Multiple fields have been commented out, search for `//REMOVED(`

Logging `log.Printf` have been changed to `glog.V(3).Infof` 



## Using the API
For documentation refer to the [Thycotic RestAPIGuide](./ThycoticSS_RestAPIGuide.pdf) or browse to https://secret.example.com/SecretServer and login then change to url to https://secret.example.com/SecretServer/webservices/sswebservice.asmx