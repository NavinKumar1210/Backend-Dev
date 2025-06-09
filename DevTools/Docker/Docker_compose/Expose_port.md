# Expose Port:
* Docker by default doesn't allow access or communicate with outside of container
* so expose is used expose a port that running inside the container to outside ports
* In docker expose we can avoid the port collision by mapping different ports from outside the container
* for example if there are 2 services that try to expose port 80 we can map one to port 80 and another to 81 (eg : 80:80 81:80)
