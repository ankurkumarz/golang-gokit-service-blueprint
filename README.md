# A Microservice Reference Implementation in Golang
- To demonstrate a Web Service built using Golang with [Go kit](https://gokit.io/) library, which is a programming toolkit with circuit breaker, tracing, pluggable serialization, and more.
- A Go kit implementation have following layers:
  - Transport layer
  - Endpoint layer
  - Service layer
- Requests enter the service at layer 1, flow down to layer 3, and responses take the reverse course. Read more [here](https://gokit.io/faq/).
## Key Decisions
- [Go Kit](https://gokit.io/) being used a toolkit to avoid reinventing the wheel when it comes to 
RPC safety, system observability, infrastructure integration, and overall program design.
- Initial implementation is using REST as an interface (future support for gRPC)

## To RUN
- Run the application:
```
go run cmd/main.go
```
- Check the endpoint:
```
 curl http://localhost:8080/entitlement/01
```

## Sequence Diagram
```mermaid
sequenceDiagram
    autonumber
    transport-->>endpoint: entitlementsvc.MakeGetEntitlementEndpoint()
    activate endpoint

    activate middleware
    endpoint-->>middleware: entitlementsvc.GetEntitlement()
    Note over endpoint,middleware: Handling logging and intermediate processing

    activate service
    middleware-->>service: entitlementsvc.GetEntitlement()
    Note over middleware,service: Service can fetch data from backend system
    Note right of service: Service has business logic + JSON payload mapping

    service-->>middleware: JSON response
    deactivate service

    middleware-->>endpoint: JSON response
    deactivate middleware

    endpoint-->>transport: JSON response
    deactivate endpoint
```