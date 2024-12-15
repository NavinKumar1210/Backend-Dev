# Reverse Proxy :
* As the name implies, a reverse proxy does the opposite of what a forward proxy does.
  Forward proxies can hide the identities of clients whereas reverse proxies can hide 
  the identities of servers. eg : Refer below image. image credit - DesignGuru


## Why Reverse Proxy :

### Load Balancing:
* Distributes incoming traffic across multiple backend servers, ensuring no single server is overwhelmed.
  This improves the application's performance and reliability.

 
### Security:
* Hides the identity and structure of backend servers from clients, adding an extra layer of security.
  It can also filter out malicious requests, protecting backend servers from attacks.

### Content Delivery:
* Reverse proxy to cache frequently accessed content, 
  reducing the load on its origin servers and improving response times for users.

### Anonymity and Privacy :
* A reverse proxy hides the server's IP address, providing an extra layer of privacy and security.

### Handle SSL/TSL  :
* It will offload encryption and decryption tasks from backend servers.it will boosts web accelaration.

### Microservice  :
* It provides a single entry point for incoming traffic.

![alt text](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fwqces5nwe4hb4bydyd13.png)
