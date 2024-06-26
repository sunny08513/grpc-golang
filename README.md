# grpc-golang

gRPC is an open-source framework developed by Google that enables efficient communication between distributed systems

Rest API(Representational State Transfer Application Programming Interface) vs GRPC(gRPC Remote Procedure Calls))

Rest API  
1. Protocol
   - Use HTTP1.1 or HTTP 2
   - STandard HTTP method GET, POST, PUT , DELETE
2. Data Format 
   - Usually use JSON format.
   - They can support other formats as well XML, HTML or Plain text.
3.  Statelessness:
   - Each HTTP request from a client to server must contain all the information 
     the server needs to fulfill the request (stateless).
4. Performance:
   - JSON is text-based and can be slower to parse and larger in size compared to binary formats.
     REST can have higher latency due to the overhead of HTTP/1.1

GRPC 
1. Protocol
   - HTTP 2
   - which allows for multiplexing multiple requests on a single TCP connection, reducing latency and improving performance.
2. Data Format:
   - gRPC uses Protocol Buffers (protobuf) for data serialization(Data serialization is the process of converting an object or data structure into a format that can be easily stored, transmitted, and reconstructed later.), which is a compact binary format.
3. Efficiency:
   - Protobuf is more efficient in terms of size and speed compared to JSON.
4. Streaming:
   - gRPC natively supports four types of APIs: unary (single request-response), 
     server streaming, client streaming, and bidirectional streaming, 
     making it more versatile for real-time communication.
5. Strong Typing:
   - Protobuf provides a strongly typed schema, 
     which helps in ensuring data integrity and provides better contract definitions.
6. Performance:
   - gRPC generally has lower latency and higher performance due to the binary format 
   and HTTP/2 features like multiplexing and header compression.
7. Use Cases:
   - gRPC is well-suited for microservices communication, high-performance internal APIs,
     real-time data streaming, and environments where low latency is critical.
8. Error Handling:
   - Uses its own status codes and metadata to provide more detailed error information.

```
Feature	            REST API	            gRPC
Protocol	         HTTP/1.1 or HTTP/2	        HTTP/2
Data Format	         JSON, XML, HTML, etc.	    Protocol Buffers (protobuf)
Human Readability	 Human-readable	            Not human-readable (binary)
Stateless	         Yes	                    Yes
Performance	         Slower(text-based,         Faster (binary, lower latency)
                     higher latency)
Streaming	Limited (e.g., WebSockets)	        Native support for streaming
Strong Typing	     No	                        Yes
Tooling	             Wide support, easy to use	Requires specific tools and libraries
Error Handling	     HTTP status codes	        gRPC status codes and metadata
Use Cases	         Web APIs, public APIs	    Microservices, real-time communication
```

Use REST if:
    You need wide accessibility and ease of use, especially for public APIs.
    Human readability and ease of debugging are important.
    You are building traditional web applications with browsers as clients.
Use gRPC if:
    You need high performance and low latency.
    You are working on microservices architecture with internal services.
    You require real-time data streaming.
    You prefer strong typing and a well-defined contract between client and server.

######For an e-commerce website with multiple microservices, deciding where to use REST APIs and where to use gRPC depends on the specific requirements of each interaction between services. Here’s a practical guide to help you make these decisions:

##When to Use REST API
 1. Public APIs:
    Use REST for APIs that are exposed to external clients, such as mobile apps, third-party services, or web clients.
    REST's widespread use and JSON format make it easy to consume by a variety of clients.
 2. Client-Facing Services:
    Use REST for services that interact directly with the frontend, such as the product catalog, user authentication, and order placement.
    The stateless nature of REST and human-readable JSON format is beneficial for these services.

##When to Use gRPC
1. Internal Microservices Communication:
    Use gRPC for internal communication between microservices where performance is critical.
    Examples include interactions between the inventory management, order processing, and payment services.
2. High Performance and Low Latency:
    Use gRPC for services requiring low latency and high throughput.
    Examples include real-time recommendation engines, fraud detection systems, and inventory updates.



Key Aspects of Data Integrity
    1. Accuracy: Data should be correct and free from errors.
    2. Consistency: Data should be uniform and consistent across different systems and instances.
    3. Reliability: Data should be reliable and usable for its intended purpose without being compromised or altered unintentionally.

#####################

# Setting Up gRPC in Golang
1. Install gRPC and Protocol Buffers
```
# Install Protocol Buffers compiler (protoc)
brew install protobuf

# Install the Go plugins for the Protocol Buffers compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Make sure the binaries are in your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```
2. Define the Protobuf File
Create a .proto file that defines your gRPC service and messages. For example, let's create a simple calculator service.

```
syntax = "proto3";

package calculator;

option go_package = "github.com/yourusername/yourproject/calculator";

service Calculator {
  rpc Sum (SumRequest) returns (SumResponse);
}

message SumRequest {
  int32 a = 1;
  int32 b = 2;
}

message SumResponse {
  int32 result = 1;
}
```

3. Generate Go Code from the Protobuf File
- Use the protoc compiler to generate Go code from your .proto file.
```
protoc --go_out=. --go-grpc_out=. calculator.proto
```

This will generate two files: calculator.pb.go and calculator_grpc.pb.go.

4. Implement the Server
   Create a main.go file inside the calculator/server directory: 

5. Implement the Client
   Create a main.go file inside the client directory:



# Error Handling in grpc

Error handling in gRPC is an important aspect of building robust and reliable services. gRPC uses standard HTTP/2 error codes and provides mechanisms to send detailed error information from the server to the client. This allows clients to handle errors appropriately based on the context.

1. Status Codes: gRPC uses a set of predefined status codes that are similar to HTTP status codes. Some common ones include:
```
OK (0): No error.
CANCELLED (1): The operation was cancelled.
UNKNOWN (2): Unknown error.
INVALID_ARGUMENT (3): Client specified an invalid argument.
DEADLINE_EXCEEDED (4): Deadline expired before operation could complete.
NOT_FOUND (5): Some requested entity was not found.
ALREADY_EXISTS (6): Some entity that we attempted to create already exists.
PERMISSION_DENIED (7): The caller does not have permission to execute the specified operation.
```

2. Error Details: gRPC errors can carry additional details beyond the status code using the google.golang.org/grpc/status package. This allows for more informative error messages and richer error handling.

# Interview Questions
https://techwasti.com/25-grpc-interview-questions-and-answers-for-developers