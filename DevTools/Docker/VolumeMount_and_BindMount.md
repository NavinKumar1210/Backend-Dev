# VolumeMount and BindMount:

### VolumeMount :
* docker volumes are used to persist the data generated and used by the Docker containers in a specific path.
* They are independent of the container lifecylcle means the data stored in volume remain even after 
container has been stopped or removed.
* use this command to add the volume to your container
* docker run -d -v <volumename>:<directorypath> <ImageName or Id>


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