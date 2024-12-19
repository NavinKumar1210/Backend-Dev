# OSI MODEL (Open Systems Interconnection) :

* OSI provides a framework for how different network protocols interact and how 
  data is transmitted across a network.

* It is a conceptual framework that divides network communications functions into 
  seven layers.

### Why OSI matters :

* The modern Internet does not follow OSI model strictly but still it will be useful 
  for troubleshooting network Problems. 

* It helps you to breakdown the problem and find the source of the issue.

### OSI Model Layers : 

* Application Layer
* Presentation layer
* Session layer
* Transport layer
* Network layer
* Data link layer
* Physical layer

#### Application Layer :
* In application Layer user can Request a content from it and it returns
  the content or data in Required Format. eg : image,pdf,json...

* client software application are not a part of application layer
  rather than It is responsible for protocals and data manipulation to 
  present meaningful data to the user

* Protocals : https,smtp,dns...

* it will add protocal headers like HTTP headers for browsing and SMTP header for mail
  

#### Presentaion Layer :

* The presentaion layer is responsible for Translation,encryption and compression of data

* The Encryption and compression for Secure and reduce the data to the next Layer

* From the receiver side it Translate and Decrypt and uncompress
  for deliver it to the Application Layer for another User

* Presentaion header add Translation,encryption and compression details in it
  
  
#### Session layer :

* This layer opening and closing connections and maintaing session between Connections

* the Session Ensures the data Exchange is Properly synchronized 

* for Example if 100mb file is Transfered it will add checkpoint of every 5mb

* this ensures if a connection is Crashed or disconnected.it restarts from last 
  checkpoint so we don't need to download from start. eg : movie or music files download 
  from web   

* Session header add session and connection management details such as sessionID and
  control Information

* It maintains connection between both user both sender and Receiver


#### Transport layer :

* This layer Ensure complete data Transfer from end to end.

* It breaks the data from session Layer into chunks called segments 
  
* From receiver side it is responsible for reassemble the segments for session Layer

* It Performs error control by ensuring the Data Transfer is complete from the reciver side 
  if it isn't it request a retransmission.  

* protocals : TCP,UDP

* Segment header (TCP) or Datagram header (UDP) 


#### The network layer :

* It is responsible for Transfer data between two networks.if both are in same 
  network the network layer is unnecessary for routing.

* It breaks up the segments from Transport Layer and make it as smaller units
  called packets

* From Reciver side it reassembles the Data and send it to Transport Layer

* It also find best Physical path to reach the Destination called routing

* Protocal : Internet Control Message Protocol (ICMP), the Internet Group Message Protocol (IGMP),IPsec suite. 

* A packet header. This header contains the source and destination IP addresses, as well as other routing information 
  like TTL (Time to Live).


#### The data link layer :

* The data link layer similar like network layer but it Transfer data between
  Same network

* It converts packets breaks into smaller unit called Frames

* It also responsible intra network connection flow and error control

* It includes two sublayers LLC ( Logical Link Control) and MAC (Media Access Control)

* Header includer source and destion like MAC address of system  


#### Physical layer :

* this layer includes physical equipment like cables and switches

* this layer converts data into bit stream like 0s and 1s

* both device agreed on signal convention so that 1s can be distinguished from 0s

* It has no header the data sent As Raw 



## Real Life Example (Mail from one to another) :

Person A tries to send mail for Person B

* writes the mail and click send it happens at application Layer

* It enters Presentaion layer and encrypt the mail and send the data

* the session was created between 2 users in Session Layer

* The data broken into segments in transport Layer  

* The data brocken into packets and find the best route in Network layer
  and Ip address added

* Data packets converted into Frames and add MAC address and send it to Physical
  Layer

* At physical layer data Transfered using cables and switches in form
  of 0s and 1s

