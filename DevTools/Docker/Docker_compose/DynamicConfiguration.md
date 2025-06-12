# Dynamic configurations :

### Named subset of services :
* if there are 2 service teams that wants to run the their own service only in the single compose file
* we can use the -profiles keyword to run the specified service only
* example : -profiles:
                upi_services
* if a container with no profiles will be included in default with every other service profiles.
* example command for compose - docker-compose --profile upi_services up

### Multiple compose files :
* we can use the multiple compose files for different environments
* by default docker compose will read 2 configurations file docker-compose.yaml and docker-compose.override.yaml
* docker compose will merge these file by merge rules
* you can name the primary and override file by your reference  
example : docker-compose.local.yaml, docker-compose.staging.yaml
* docker compose command : docker-compose -f [primary-file] -f [override-file] command
* if there are fields like array docker compose will add all the data from primary and override file
* if there is any single field compose will give preference override file over primary

#### Environment Variables :
* if docker-compose have different behaviour in different environment you don't wants to support override
you can use environment Variables.
* common Use cases : image tags,version,ports
* example: image:"mongodb:${TAG}"
* you can use set environment Variables in inline docker compose configurations or add ENV file
* docker-compose --env-file [path]
* a Variables access by the host environment will always override the default value
* you can throw err if the value is not present
example : "mongodb:?mongodb value is not found"