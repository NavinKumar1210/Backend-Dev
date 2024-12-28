# SMTP Protocal :

* SMTP (Simple Mail Transfer Protocol) is a fundamental and foundation protocal for sending and receiving emails 
  over the network
* It defines rules for communication and allows email clients and servers to
  exchange the messages<br>

### Key Components:

##### SMTP Client: 
* Initiates a connection with the SMTP Server to send emails.
  
##### SMTP Server: 
* Receives emails from clients, processes them, and forwards them to the intended recipient’s server.
  
##### Commands and Responses: 
* SMTP clients and servers communicate using commands (e.g., HELO, MAIL FROM, RCPT TO) 
  and responses (e.g., 250 OK, 550 Error).

### How It Works:

* The SMTP client initiates a connection with the SMTP server using the HELO (Hello) command.
  
* The client transfers email headers, recipients, body messages (including attachments), and data to the server step-by-step.
  
* After transmission is complete, the connection is closed.

### Limitations :

* SMTP can only send plain text emails
  
* SMTP have charecter limits for message body
  
* If you want to surpass this limitations you can use MIME (Multipurpose Internet Mail Extensions) with SMTP
  or You can use ESMTP ( Extended Simple Mail Transfer Protocol).

### Real Life Examples :
* Gmail and Microsoft Outlook use the SMTP to send and receive emails

