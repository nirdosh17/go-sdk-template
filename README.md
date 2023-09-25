[![Go Report Card](https://goreportcard.com/badge/github.com/nirdosh17/go-sdk-template)](https://goreportcard.com/report/github.com/nirdosh17/go-sdk-template)

# Go SDK Example
Sample client SDK in Go with most useful features for production

## Requirements
Let's say we own a SAAS company which specializes on AI and we want to create a Golang SDK for developers so that they can interact with our services programmatically.
We offer a service called `ChatAI` at the moment but our company can grow in future and offer multiple services. So, we need to design the structure of SDK accordingly.

The sample SDK contains most common features which are important for production.
### Configuration Options
Check [documentation](https://pkg.go.dev/github.com/nirdosh17/go-sdk-template/config) for all available options.
- **Authentication**

  Just injects headers while making API requests, rest is server's responsibility.

- **Option to pass Context**

  We can pass context to define timeouts and cancellations.


- **Custom HTTP Client**

  We can pass our own client for fine grain control (e.g. proxy settings)


- **Custom Errors**

  Custom error type allows to check type of error via code instead of string match.

- **Retry mechanism**
  - Implements fixed interval based retry
  - Max retries can be configured
  - Custom retry function can be passed if we want to implement our own retry strategy

- **Logging**
  - Option to enabled verbose logging (http dumps)
  - Use own custom logger

## Usage
**Install**

`go get github.com/nirdosh17/go-sdk-template`


### Using the API

We have one API at this point called ChatAI inside chatai package.

**With default configs**
```go
c := config.NewConfig("enter-api-key")
ai := chatai.NewService(c)
// without using context
answer, err := ai.Ask("how does Go scheduler work?")
```

**With Custom HTTP Client**
```go
// example: custom http client with proxy settings
proxyURL, err := url.Parse("http://asia.backendapi.com")
if err != nil {
  log.Println("error parsing proxy url", err)
}
transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
customClient := &http.Client{Transport: transport}

// see documentation for more options
c := config.NewConfig("enter-api-key").
      WithHTTPClient(customClient).
      WithMaxRetry(5)


ai := chatai.NewService(c)
ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
answer, err := ai.AskAIWithContext(ctx, "how does Go scheduler work?")
```


## Folder Structure
All files starting with `example_` contains tests and documentation which is displayed in [pkg.go.dev](https://pkg.go.dev/github.com/nirdosh17/go-sdk-template). `doc.go` file has also been used for same purpose.
```
├── config
│   ├── config.go
│   ├── config_test.go
│   ├── example_retryer_test.go
│   └── example_test.go
├── api                           // each folder represents a service
│   └── chatai                    // one of the services offered by our dummy company
│       ├── doc.go                // it is displayed as overview in pkg.dev.go
│       ├── error.go              // errors related to this service
│       ├── examples_test.go      // test + documentation
│       ├── interface.go          // interfaces for DI and mocking
│       └── service.go            // contains APIs offered by the service
├── apierror
│   └── error.go                  // error interface, custom error types and common errors codes
├── client
│   ├── httpClient.go             // http requester interface
│   ├── requester.go              // requester implementation
│   ├── requester_test.go
│   ├── retryer.go                // retry interface and default retry function
│   └── retryer_test.go
├── logger
│   └── logger.go                 // logger interface and default logger
├── model
│   └── model.go
├── test
│   └── helper.go                 // helper methods for tests
├── LICENSE
├── Makefile
├── README.md
├── doc.go                        // high level doc
└── go.mod
```
