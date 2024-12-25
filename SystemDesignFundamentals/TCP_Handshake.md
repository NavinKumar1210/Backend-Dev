# TCP Handshake :

* TCP handshake is a process establishing a communication between 2 devices
* It have three steps :
    * SYN (Synchronize)
    * SYN-ACK (Synchronize-Acknowledge)
    * ACK (Acknowledge)

### SYN :

* In this step the client (browser) wants to establish connection to the server
* so it sends a SYN Packet and synchronized sequence number to server to start communication

### SYN-ACK :

* The Server responds with SYN-ACK Packets.
* this packet acknowledge SYN Packet and sequence number.

### ACK :

* The client sends the ACK packet back to the server by acknowledging SYN-ACK packet.
* After this point the connection is established the data transfer can begin.


