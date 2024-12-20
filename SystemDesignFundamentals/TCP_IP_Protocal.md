# TCP (Transmission Control Protocal) Model :

  Note : Please Check OSI model before this It wil help you easily undestand this

* TCP/IP model is the backbone of modern internet communication  

* Tcp is a Standard protocal that enables application to send the data over network
  
* Tcp prefers reliablity over speed while UDP prefers speed over reliablity

* It ensures reliablity by error free and flow control and retransmission if connection issues. 


### Layers of TCP :

* Application Layer

* Transport Layer

* Internet Layer

* Network Access Layer

#### Application Layer :

* This Layer interface with user application like web browser and email
  
* Protocals : HTTP,FTP,SMTP,DNS...
  
* It ensures the data format suitable for the application. eg :pdf,txt...
  
* Directly interact with the end users 

#### Transport Layer :

* Transport layer provides end-to-end communication by managing data
  Transfer between devices

* Protocals :TCP for reliable communication and UDP for unreliable communications

* It uses error checking and flow control

* It breaks data into small segmentaions

#### Internet Layer :

* Internet layer identify the device on the internet using IP address
  
* Uses routing and forwarding to find the bast path across the network
  
* Protocals : IP,ICMP,ARP
  
* Break down the segemented data into packets


#### Network Access Layer :

* Network Access Layer breaks the packets into frames 
  
* It have error detection and physical transmission
  
* Provides Local network communication by identifying the devices
  using MAC address