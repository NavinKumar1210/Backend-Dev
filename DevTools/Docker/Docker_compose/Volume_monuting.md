# Volume Mounting:
* Volumes are persistent container storage in the system even when the container is removed or stopped the important data is not deleted
* volume has target and source target is the directory in the container and source is the directory of the system

### Bash directory syntax :
* ./ - curr directory
* ../ - parent directory
* . - root directory
     
### Access modes :
* you can mention the access mode in the compose file
* rw - read/write , ro -read-only , wo-write-only
* aws region

### Named Volumes :
* Named volumes are used to persist the data even after compose is down.
* Nameless volumes will be created each time when docker compose up
* It will copy the data from the old container to the new one