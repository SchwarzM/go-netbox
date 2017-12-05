### Generated go-netbox

1. generated swagger.json by ```curl "http://localhost/api/docs/v2?format=openapi" | jq '.' > swagger.json```
2. added consumes and produces fields to swagger.json to ensure the accept header is right
3. ran ```swagger generate client```


