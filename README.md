# Distributed-Token-Manager-System

<pre>    <img src="/Readme PNG's/Readme-01.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-02.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-03.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-04.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-05.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-06.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-07.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-08.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-09.png" >     </pre>
<pre>    <img src="/Readme PNG's/Readme-10.png" >     </pre>
Explanation of the Token Manager Project
What is the Project?
The project is a fail-silent replicated token manager with atomic semantics, implemented using the read-impose-write-all protocol. It manages tokens as a (1, N) register, where a single writer interacts with multiple readers to handle read and write operations in a distributed, fault-tolerant system.

Core Concepts:
1. Fail-Silent Model:
A fail-silent system ensures that nodes stop functioning (fail silently) instead of producing incorrect results when they encounter a failure.
In this project:
A node simulates failure by going to "sleep" (stops responding to requests), but the system continues to function by achieving a majority quorum among the remaining nodes.
2. Atomic Semantics:
Guarantees that read/write operations are all-or-nothing:
A read request always retrieves the latest value.
A write operation is committed only when it’s acknowledged by a majority of nodes.
3. Read-Impose-Write-All Protocol:
Ensures atomicity by:
Broadcasting the read request to all nodes to collect token values and timestamps.
Imposing the most recent value with the latest timestamp and broadcasting it to all nodes.
For write operations, the new value is broadcasted, and acknowledgments from a quorum of nodes ensure the write's success.
Project Architecture:
Server:
Multiple reader nodes and a single writer node serve client requests.
Each node is distributed across unique IP addresses and ports.
Client:
Sends read or write requests to the token manager system:
Read Request: Retrieves the latest token value.
Write Request: Updates a token's value, broadcasting it to all nodes.
Key Components:
RBroadcast & ReadOne: Handles broadcasting and acknowledgment for read operations.
WBroadcast & WriteOne: Manages broadcasting and acknowledgment for write operations.
Fail-Silent Simulation: Introduces failure by simulating unresponsive nodes during read or write requests.
Execution Flow:
Read Operation:
The client sends a read request to a server.
The server broadcasts the request to all nodes, collecting token values and timestamps.
The server waits for acknowledgments from a quorum of nodes (majority of N/2).
The latest value (based on timestamps) is:
Updated at the current server.
Broadcasted to all nodes for consistency.
The server responds to the client with the latest token value.
Write Operation:
The client sends a write request with new token data.
The server updates its local token value with a new timestamp.
The server broadcasts the new value to all nodes.
After receiving a quorum of acknowledgments (N/2), the server confirms the write operation to the client.
Protocols in Action:
Read-Impose-Write-All Protocol:
For Read Requests:
Ensures the client always gets the latest value by gathering and imposing the most recent token data across nodes.
For Write Requests:
Guarantees consistency by ensuring that the new token value is written to a quorum of nodes before being acknowledged.
Fail-Silent Model Simulation:
Failure is simulated by introducing a "sleep" in the ReadOne or WriteOne function for a specific node.
The system achieves quorum by relying on the remaining responsive nodes, ensuring fault tolerance.
Key Tools and Implementation Details:
Go Language:
Server and client logic is implemented in Go.
gRPC and Protocol Buffers:
Enables communication between nodes and clients.
Bash Script for Automation:
Simplifies starting servers and clients for testing.
Quorum-Based Consensus:
Ensures system reliability and consistency by requiring a majority of nodes to agree.
How to Run the Project:
Servers:
Start multiple servers on different ports:
bash
Copy code
go run server.go -ip 127.0.0.20 -port 8000
go run server.go -ip 127.0.0.30 -port 8001
Clients:
Read Operation:
bash
Copy code
go run client.go -read -id 101 -host 127.0.0.20 -port 8000
Write Operation:
bash
Copy code
go run client.go -write -id 101 -name abc -low 0 -mid 10 -high 100 -host 127.0.0.20 -port 8000
Why This Project is Significant:
Fault Tolerance:
The fail-silent model ensures system reliability despite node failures.
Atomicity:
Guarantees consistent and reliable read/write operations in distributed systems.
Scalable Design:
Supports adding more reader/writer nodes without impacting core functionality.
Would you like further clarification on any specific aspect of the project or its protocols?






You said:
So what is an rpc call? gRPC and are these requests in the project asynchronous?
ChatGPT said:
ChatGPT
What is an RPC Call?
An RPC (Remote Procedure Call) allows a program to execute a function or procedure on a remote server as if it were a local call. The client makes a request to invoke the remote function, and the server executes it and returns the result.

Key Characteristics:
Transparency:
The client doesn’t need to know the underlying details of the remote function, such as where it’s hosted or how it’s implemented.
Client-Server Model:
The client sends requests, and the server handles them by executing the requested function and returning results.
Interoperability:
RPCs work across different platforms and programming languages using standard protocols.
What is gRPC?
gRPC is a modern, high-performance RPC framework developed by Google. It’s widely used for building distributed systems due to its speed, efficiency, and interoperability.

Key Features:
Protocol Buffers (Protobuf):
gRPC uses Protobuf to define the structure of requests and responses, making data serialization fast and efficient.
Multi-Language Support:
Supports multiple languages, such as Go, Python, Java, and C++, enabling cross-platform communication.
HTTP/2-Based Communication:
Leverages HTTP/2 for multiplexing, bidirectional streaming, and reduced latency.
Streaming Capabilities:
Supports unary RPCs (one request, one response) and streaming RPCs (continuous data streams).
gRPC in Your Project:
The project uses gRPC to enable communication between the client and server, as well as between server nodes.
Example:
A client sends a Read or Write request to a server using gRPC.
The server processes the request and responds via gRPC, ensuring efficient and reliable communication.
Are the Requests Asynchronous?
In gRPC:
Requests in gRPC can be synchronous or asynchronous, depending on how the client and server are configured.
In Your Project:
Based on typical implementations for distributed systems:

Asynchronous Server-to-Server Communication:

When a server broadcasts a Read or Write request to other nodes, these operations are usually asynchronous.
This allows the server to continue processing other tasks while waiting for acknowledgments from a quorum of nodes.
Client-to-Server Communication:

Client requests to the server (Read or Write) are typically synchronous in your project to ensure that the client receives a definitive response (e.g., a successful write or the latest token value).
Why Asynchronous Requests?
Server-to-Server Asynchronous Calls:

Improves efficiency by allowing the server to handle multiple tasks concurrently.
Reduces waiting time for acknowledgments, especially in fail-silent scenarios where some nodes might be unresponsive.
Client-to-Server Synchronous Calls:

Ensures reliability and consistency, as the client must wait for the operation to complete and return a valid result.
How gRPC Handles Asynchronous Requests:
Server-Side:
gRPC provides APIs to handle incoming requests asynchronously, allowing the server to process multiple calls simultaneously.
Client-Side:
Asynchronous clients can initiate non-blocking calls, receiving results through callbacks or futures.
Summary:
RPC Call: A mechanism to invoke functions on remote servers transparently.
gRPC: A modern, efficient RPC framework using Protobuf and HTTP/2.
Request Types in the Project:
Client-to-Server: Likely synchronous for reliability.
Server-to-Server: Likely asynchronous for scalability and fault tolerance.
Would you like more technical details or examples of asynchronous gRPC calls in your project?
