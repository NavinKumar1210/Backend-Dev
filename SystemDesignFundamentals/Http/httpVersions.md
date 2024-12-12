# HTTP Versions

## HTTP 1.1 :
            * Transport Protocal : TCP
            * Persistent connections : Allows multiple requests and responses over a single connection one by one.
            * Chunked Transfer : server can send the response by chunks

## HTTP 2  :
            * Transport Protocal : TCP
            * Multiplexing : Multiple requests and responses can be sent over a single connection simultaneously, reducing latency.
            * Header compression :Reduce the header size using HPACK compression
            * Stream Prioritization : Allow client to prioritize requests
            * Server Push: Enables servers to send resources to the client proactively.
            * Head-of-line blocking issue still in Transport layer at TCP

## HTTP 3 :    
            * Transport Protocol: QUIC (It use UDP instead of TCP)
            * Built-in Encryption: All connections are encrypted by default, enhancing security.
            * Eliminates Head-of-Line Blocking: QUIC's Stream are delivered independently
              so packet loss affecting one Stream doesn't affect others
            
## Summary
            *  HTTP/1.1: Introduced persistent connections and chunked transfer encoding but suffered from head-of-line blocking.
            *  HTTP/2: Improved performance with multiplexing, header compression, and stream prioritization, but still relied on TCP.
            *  HTTP/3: Uses QUIC for faster, more reliable connections, eliminating head-of-line blocking and enhancing security.