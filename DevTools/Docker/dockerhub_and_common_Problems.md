# Docker Hub and common Docker problems:

### DockerHub :
* docker login - to login via CLI
* enter the username 
* after successfull login use the following command
* docker push <UserName>/<containerName>:<version> - you can provide version if you want

### common Problems :

##### unable to create more containers :
* It is usually happen because docker images are compressed image of intermediary images and containers also have significant storage compared to images.
* In windows and MacOs it is also running inside the tiny virtual machine
* you can use these following commands to remove the images and containers 
* docker rmi <imageId> - to remove the image
* docker rmi -f <imageId> - to remove the image forcefully
* docker rm <containerId>
* docker system prune - remove all the stopped containers unused images dangling images and build cache

##### My container is very slow :
* docker stats <containerId or name> - you can debug by the statistics of the container by using this command
* example Data:
CONTAINER ID   NAME             CPU %     MEM USAGE / LIMIT     MEM %     NET I/O        BLOCK I/O     PIDS 
eab6617c1544   hopeful_turing   0.00%     5.016MiB / 7.653GiB   0.06%     1.7kB / 126B   4.58MB / 0B   2 
* docker top  <containerId or name> - you can use this to check number of running processes inside the containers
* docker inspect <containerId or name> - it will give advance information about the container in the json format