# Spring Cloud DataFlow - Go SDK
This is a go SDK for [Spring Cloud DataFlow's](https://spring.io/projects/spring-cloud-dataflow/) [Open API](https://docs.spring.io/spring-cloud-dataflow/docs/current/reference/htmlsingle/#api-guide) generated by [Kiota](https://github.com/microsoft/kiota)

## Use
```
go get github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2@v2.11.2-prerelease1
```

Example
```
// API requires no authentication, so use the anonymous
// authentication provider
authProvider := auth.AnonymousAuthenticationProvider{}

// Create request adapter using the net/http-based implementation
adapter, err := http.NewNetHttpRequestAdapter(&authProvider)
if err != nil {
    return nil, err
}

// Create the API client
client := client.NewDataFlowClient(adapter)


result, err := client.Apps().ByType(app.Type).ByName(app.Name).ByVersion(app.Version).Get(ctx, nil)
```

## (Re-) Generate
1. Extract OpenAPI Specification from Spring Cloud DataFlow (`http://localhost:9393/v3/api-docs`) by starting [docker-compose setup](https://dataflow.spring.io/docs/installation/local/docker/#installing-by-using-docker-compose) with following extra environment variables in `data-flow-server`:
```
SPRINGDOC_API_DOCS_ENABLED=True
SPRINGDOC_SWAGGER_UI_ENABLED=True
```

2. Insert extracted OpenAPI Specification into file `./dataflow-api.yml`

3. Install kiota binaries:
https://learn.microsoft.com/en-us/openapi/kiota/install?tabs=bash#download-binaries

4. Run Code generation
```
kiota generate -l go -c DataFlowClient -n github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client -d ./dataflow-api.yml -o ./client
```

5. tidy, compile and test
```
go mod tidy
go test ./...
```

6. Commit, Tag and Push


7. Make the module available (see [here](https://go.dev/doc/modules/publishing))
```
GOPROXY=proxy.golang.org go list -m github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2@<version>
```


## Other
Created by following [this](https://learn.microsoft.com/en-us/openapi/kiota/quickstarts/go) Guide 