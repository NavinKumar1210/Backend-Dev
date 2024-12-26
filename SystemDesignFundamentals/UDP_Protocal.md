# UDP Protocal :

* User Datagram Protocal is a fundamental protocal used to send thge data over the network
  like TCP

#### Why UDP :

* Unlike TCP UDP prefer speed over reliablity
  
* So it is useful for applications where speed is crucial and data loss or 
  duplicatoion is acceptable

* For example video streaming and Online gaming

#### Structure of UDP Packet :

##### header : 
* Source Port : the port number of sender
  
* Destination Port :port number of receiver
   
* Length : the length of UDP header and data
   
* Checksum : used for error checking and header

##### Data : 
* the Data for the Request


#### Key charecteristics :

##### Connectionless :
* UDP is connectionless meaning it does not establish a connection before sending the data
* Each packets send independently

##### No Guarantee of Delivery:
* UDP doesn't guarantee of delivery of packets.
* There is no acknowledgement or receipt.
* The packets may not arrive or out of order or duplicates in it.

##### Low overhead :
* It has a lower overhead compared to TCP because it doesn't maintain a connection or 
  handshaking this makes it faster and efficient.