# Docker Best practices
* use only verified Images they are more secure
* Container image scanner - to check the intermediary images and warn you about the malicious layers
of the docker image exmple : clair ,trivy,Dagda
* Avoid Latest tag - if you don't provide tag name for the conainer it will be tagged as latest
It is not ok if you don't know the version you are downloading and latest can be overridden making rollback difficult
* Docker uses root as the default user in the container so use non root users when running the containers for more secure within the container and ouside of it.
if some of the cmd or files need root access you can use root user and the switch it into non user using USER command
* if your container relying on multiple services inside a server you can use docker compose to run and connect multiple containers on a single machine
* if your container relying on multiple services in different servers you can use orchestration tools like kubernates to create move and scale containers automatically. it's also make grouping and scalling very easily and load balancing and securing container traffic is much easier with kubernates