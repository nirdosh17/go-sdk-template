# Go SDK Example
Sample client SDK in Go with most useful features for production

## Requirements
Let's say we own a SAAS company which specializes on AI and we want to create a Golang SDK for developers so that they can interact with our services programmatically.
We offer a service called `ChatAI` at the moment but our company can grow in future and offer multiple services. So, we need to design the structure of SDK accordingly.

The sample SDK contains most common features which are important for production readiness.
### Features
- Authentication:
  - Injects credentials while making API requests, rest is server's responsibility.
- Allows to use `context` while making request. timeouts and cancellations can be handled via this.
- Custom HTTP client can be used for fine grain control (e.g. proxy settings)
- Return custom error instead of plain text
- Retry mechanism
  - Implements fixed interval based retry
  - Max retries can be configured
  - Custom retry function can be used to implement different strategies
- Input Validator [Not implemented]
- Logging:
  - option to enabled verbose logging via `config.Debug` flag
  - Able to use custom logger
- Tests

## SDK Structure
```
.
├── config
│   ├── config.go         // exposes main configs to initialize SDK
│   └── config_test.go
├── apierror
│   └── error.go          // common SDK level errors
├── client                // client to interact with external services
│   ├── client.go         // HTTP requester
│   ├── retryer.go        // retry interface and retry function implementation
│   └── retryer_test.go
├── api                   // APIs to talk to services
│   ├── chatai            // example service 1
│   │   ├── error.go      // errors related to this service
│   │   ├── interface.go  // to enable mocking
│   │   └── service.go    // exposes APIs for SDK consumers
│   │
│   └── anotherservice    // example service 2
│       ├── error.go
│       ├── interface.go
│       └── service.go
├── test
│   └── helper.go
├── Makefile
├── README.md
├── doc.go                // root level Go doc
└── go.mod
```
