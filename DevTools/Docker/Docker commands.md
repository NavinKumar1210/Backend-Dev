# Docker commands

* docker container create hello-world - create the container based on the image hello-world
* docker ps - list the actively running containers 
* docker ps -a -list all the containers 
* docker ps -s - list running conatainers with SIZE
* docker start <containerId> - cb2 
* docker start cb2 --attach - start the container and attach the container output in the terminal
* docker logs <containerId> - to get the logs of the container
* docker logs --tail 1000 <containerId> -to get the tail of the logs contains 1000 lines
* docker logs --details <containerId> - used to get the extra details of the container
* docker logs --since <UTC> <containerId>- to get the log since the given time
* docker logs --until <UTC> <containerId>- to get the log until the given time
* docker stop <containerId>
* docker stop -t <seconds> <containerId> - to stop the container in mentioned seconds
* docker kill <containerId> - terminate the container Immediately
* docker ps -aq | xargs docker rm - to remove all the containers
* docker rmi <imageId> - to remove the image
* docker rmi -f <imageId> - to remove the image forcefully
* docker exec -it <containerId> bash - to enter the bash terminal of the container
* docker rm <containerId>
* docker build --file <filename> --tag timeserver . -tell the docker to run the given as dockerfile and add the image name using --tag.
* docker run -d --name <containername> <imageId or name> - to run the conatiner with the name
* docker run -d --name webserver -p 5001:5000 <containerId> - used to bind the ports in the Docker 5001 is sytem port 5000 is container port
* docker rmi -f $(docker images -q) - to remove all the docker images
* docker images prune - to remove all dangling images you can use -a to remove non used images
* docker container prune - same like image prune removes dangling containers
* docker system prune - remove all the stopped containers unused images dangling images and build cache
* docker search --filter stars=100 <ImageName> - docker search is used to search the docker images from
Docker hub you can add filter for each fields
* docker search --limit <ImageName> - limit the search options
* docker image pull <Imagename>:<version> - to pull the image from Dockerhub
* docker tag <oldImageName>:<version> <newImagename>:<version>- to add the new tag to the Image

### easy way to start the container :
* docker run hello-world - it will create and start and attach the logs of the container


Example:

