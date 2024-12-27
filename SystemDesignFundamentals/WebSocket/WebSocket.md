# WebSocket :

* WebSocket is a modern technology that enables a persistent and bidirection communication channel
  between client (web browser) and server.

* This allows real time data exchange without makin new http requests or reload the page.

* Websocket connections can be secured using WSS (Websocket Secure) it encrypts the data that has been transmitted. 

* Example file in golang Websocket.go

### Why WebSocket :

##### Bidirectional Communication :
* Websocket enables Bidirectional Communication so client and server can send request and receive data simultaneously.

##### Low Latency :
* Websocket maintains persistent connection so it reduce the delay of making connection and transfer the data quickly.

##### Reduced Network Overhead :
* It uses a single persistent connection instead of multiple http requests.
  So it reduces the amount of Data sent over the network

##### Supports Binary Data :
* Websocket supports Binary Data so you can send multimedia files like images,videos or complex data

##### Efficiency :
* Websocket is more efficient than http polling or long polling methods which requires
  repeated request to the server

##### Upgradable :
* WebSocket can be used over HTTP ports 443 and 80, making it compatible with HTTP.
<br>

### How WebSocket Works :

##### Connection Establishment :
* The communication begins with a standard HTTP request from the client to the server.If the server supports WebSocket,
   it responds with an HTTP 101 status code, indicating that the protocol is switching to WebSocket.


##### Persistent Connection: 
* Once the connection is established, it remains open, allowing both the client and server to send and receive messages at any time.
<br>

### Use Cases :

* Real Time Applications like chat applications, live scores, Online Gaming,Live Streaming 
* Push Notifications : It allows server to push notifications without request from the client.
  Example : mobile apps notifications