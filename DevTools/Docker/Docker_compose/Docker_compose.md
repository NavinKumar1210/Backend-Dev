# DockerCompose:
* docker compose is a tool for running and defining multi container applications
* docker compose is a yaml file that defines the services,networks,volume for a docker application
* docker compose is designed for a single hosted server it is well suited for local,staging or CI testing env
* It is not designed for distributed systems so it can't run containers across multiple hosts
* it is markdown tool for developer to specify the docker configurations as a code

### docker compose file contains :
* version: specifies the version of the Compose file format
* services: defines the services (containers) in the application
* networks: defines the networks for the application
* volumes: defines the volumes for the application

### Commands :
* docker-compose up: starts the containers defined in the Compose file, it will build the image create the container and start the application
* docker-compose down: stops and removes the containers defined in the Compose file, stop all the containers delete containers and images remove all artifacts.
* docker-compose restart :restart the container
* docker-compose ps: lists the containers defined in the Compose file
* docker-compose logs: displays the logs for the containers defined in the Compose file
* docker-compose exec: executes a command in a running container
* docker-compose --help : list all the options of docker compose