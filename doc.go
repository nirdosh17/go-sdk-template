// go-sdk-template is an sample client sdk in Go with the most common features for a production ready client.
//
// Features offered by the SDK:
//
// # HTTP Client
//
// Default HTTP client can be overridden by own client from config. This enables us to have fine grained control over our requests. For example, using a proxy server.
//
// # Retry Logic
//
// The sdk has a default retry algorithm which is based on a fixed window delay. We can overwrite this behavior by passing our own retry function which satisfies the given Retryer interface.
//
// Default max retry count can also be overridden if needed.
//
// # Use of context
//
// Timeouts and cancellations can be handled by passing `context`. e.g. service.AskAIWithContext(ctx, ...)
//
// # Logging
//
// We can enable debug mode for verbose logging. When debug is enabled, it prints out http requets and response objects.
//
// We can also override default logger with our own. `config.WithLogger(logger)`
//
// # Override default service endpoint
//
// In some cases, we might want to a different API endpoint offered by the service. For example AWS has region based endpoints. We can override the endpoint using
//
// Example:
//
// config.WithEndpoint()
package go_sdk_template
