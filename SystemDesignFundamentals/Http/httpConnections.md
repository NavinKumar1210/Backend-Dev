# HTTP Connections

### HTTP 1x :
##### Short-lived connections (http 1.0) : 
* In short Lived connections Each http request have it's own connection.
you can't send multiple requests over a single connection. 

* It also means Tcp hanshake happens for every request which is time consuming
##### Persistent connections :
* It allows multiple requests and responses to be sent over a single connection 
without reopening it for each request. This reduces latency and improves performance.

* The Connection and Keep-Alive headers are used to manage connection persistence and 
control how long a connection should be kept open
##### Pipelining :
* In pipelining we can send multiple requests without waiting for the corresponding
Responses

### HTTP 2 :
#####  Multiplexing : 
* HTTP/2 Update persistent connections by eliminating the head-of-line 
blocking issue present in HTTP/1.1.

#####  Header Compression : 
* HTTP/2 uses HPACK compression to reduce the size of HTTP 
headers, improving performance.

#####  Stream Prioritization :
* HTTP/2 supports prioritization of streams, allowing more 
important resources to be delivered faster.

#####  Server Push : 
* This feature allows servers to send resources to the client proactively, reducing 
latency by eliminating the need for the client to request these resources.

### HTTP/3 :
#####  QUIC Protocol :
* HTTP/3 is built on the QUIC transport protocol, which provides improved connection
setup times and better handling of packet loss compared to TCP.

##### Improved Security : 
* QUIC integrates TLS 1.3, providing enhanced security and faster handshakes.

##### Connection Migration : 
* QUIC supports connection migration, allowing connections to continue seamlessly 
even if the client's IP address changes.