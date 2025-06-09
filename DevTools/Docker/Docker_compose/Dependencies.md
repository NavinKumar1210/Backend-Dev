# Dependencies:
* Many applications have service dependencies like a service unable to start without another service.
* example : database container has to be start first before application container
* depends_on field is used to specify the dependencies containers in the yaml file
* example : depends_on :
                - database
* when we use depends on the docker compose up first start the dependency service the our application
and docker compose down will stop our service first then the dependency services
* docker compose up doesn't guarentee the Uptime of dependency services it only guarentee it has started
* 