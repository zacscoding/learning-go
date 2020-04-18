> Dependencies for json rpc server  

```cmd
$ go get github.com/gorilla/rpc
$ go get -u github.com/gorilla/mux
```  

> SMS service

```cmd
curl -X POST \
  http://localhost:3000/rpc \  
  -H 'content-type: application/json' \  
  -d '{
	"method" : "sms.Send",
	"params" : [
		{
			"number" : "010-123-1234",
			"content" : "Hello"
		}
	],
	"jsonrpc":"2.0",
	"id" : 2
}'
```  

```cmd
curl -X POST \
  http://localhost:3000/rpc \  
  -H 'content-type: application/json' \  
  -d '{
	"method" : "sms.IsSended",
	"params" : [
		"010-123-1234"
	],
	"jsonrpc":"2.0",
	"id" : 1
}'
```  

> EMAIL service  

```cmd
curl -X POST \
  http://localhost:3000/rpc \
  -H 'content-type: application/json' \
  -d '{
	"method" : "email.Send",
	"params" : [
		{
			"email" : "zaccoding@github.com",
			"content" : "Hello"
		}
	],
	"jsonrpc":"2.0",
	"id" : 2
}'
```
