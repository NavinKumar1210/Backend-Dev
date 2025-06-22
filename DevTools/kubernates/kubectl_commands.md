# Kubectl commands :
* kubectl run mynginx --image=docker.io/library/nginx - used to run the container
* kubectl api-resources - tell all the resources in the kubernates
* kubectl explain <resource name> - explain all the details about the resource
* kubectl -h - to get all the available options and commands on kubectl
* kubectl describe pod <podname> - shows the properties of currently running pod
* kubectl delete pod <podname> -stop the pod if you want immediate you can use --force option
* kubectl delete deployment <podname> - delete the Deployment and all its Pods
* kubectl create -f <yaml_file> -create resource from the yaml file
* kubectl apply -f <yaml_file>: applies a YAML file to a resource, creating or updating it as needed
* kubectl delete -f <yaml_file>: deletes a resource specified in a YAML file
* kubectl explain -f <yaml_file>: explains the fields and values in a YAML file
* kubectl run <podname> --image=<imgname> --dry-run=client -o <format> - use this command to generate the yaml or the json or go template format
* 