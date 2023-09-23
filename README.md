# Go SDK Example
Sample client SDK in Go with most useful features for production

## Requirements
Let's say we own a SAAS company which specializes on AI and we want to create a Golang SDK for developers so that they can interact with our services programmatically.
We offer a service called `ChatAI` at the moment but our company can grow in future and offer multiple services. So, we need to design the structure of SDK accordingly.

The sample SDK contains most common features which are important for production readiness.
### Features
- Authorization
  - Auth is server's responsibility. SDK only injects auth header in the request.
- Able pass `context` while making request
- Configurable HTTP Timeout
- Able to use custom HTTP client for fine grain control (e.g. proxy settings)
- Custom errors/codes instead of plain text
- Retry mechanism
  - Implements fixed interval based retry
  - Configurable max retries
  - Able to overwrite default retry behavior with custom retry function. Can be useful to implement different strategies like exponential backoff
- Input Validator [Not implemented]
- Versioning
  - Git tags are good for a single service but if your library has multiple services then package level segregation can be done as shown below:
    - `/api/chatai/v1`
    - `/api/texttoimage/v1`
- Tests [In-progress]

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
