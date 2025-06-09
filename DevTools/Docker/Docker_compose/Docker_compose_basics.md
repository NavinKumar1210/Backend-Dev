# DockerCompose basics :
* Docker compose is declarative rather than procedural
* procedural - series of ordered steps, next step will execute based on previous step
* declarative - specify the end result , system will determine which steps to execute and get the end result

### docker compose benefits :
* version control
* self documenting
* easy management

### when to use docker compose :
* Local Development
* Staging Server
* Continuous Integration Testing Environment


##### commands :
* docker volume create <VolumeName> - create the volume
* docker volume ls - list all the volumes
* docker volume inspect - Display detailed information on one or more volumes
* docker volume rm - Remove one or more volumes
* docker volume prune - Remove unused local volumes

### BindMount :
* BindMount are used to mounts a host directory into the container
* They are dependent of the container lifecylcle means the data stored in volume will be deleted after 
container has been stopped or removed.
* you can mount either a single file or Directory
* use this command to add the volume to your container
* docker run -d --mount type=bind,source=<directorypath> <ImageName or Id>