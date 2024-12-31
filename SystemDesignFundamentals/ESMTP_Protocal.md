# ESMTP Protocal (Extended SMTP) :

*  ESMTP is an enhancement of the Simple Mail Transfer Protocol (SMTP), designed to provide additional 
   functionality and flexibility in email communication.
*  Kindly Read SMTP protocal before this for a better understanding.
  

#### Why To Use ESMTP :

##### Attachment :
* SMTP only support text only emails .
* Use ESMTP If you need to send emails with Attachment like img,video..etc.

##### Enhanced Security :
* It has Authentication and Encryption to secure Emails.

##### Advanced Feautures :
* Specify message limits
* Various charecter sets for text message
  
##### High Volume Mails :
* Use ESMTP when you want to send the large number of mails.


#### How ESMTP Works :

* EHLO Command: When an email client connects to an ESMTP server, it uses the EHLO command instead of the HELO command used in SMTP. 
* If the server also supports ESMTP the connection will be established.
* After the connection made the client send the mail to the server
* The server verifies the IP address and check Authentication of the client
* Once it verified the server will route the email to the receiptent server
* Then the receiptent will receive the mail and store it in inbox.