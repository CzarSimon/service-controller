#service controller
Command line tool for simplifying running services using docker swarm.

The goal of the tool is to not have to manually ssh into hosts and to start/stop
or alter the state of a swarm cluster and keep images on host up to date as well
as statically specify service configutations similarly to docker compose.

##Run as follows:

$ sctl **start service_name**

Locates the specified services configuration file and instructes the swarm
master to run the service acording to the configuration.

$ sctl **update service_name**

Pushed the image of the supplied service to the users image registry and then
instructs all the nodes in the cluster to pull the new version.

###Service config format
```json
{
  "name": "service_name",
  "image": "some_image",
  "keywordArgs": [
    "-t", 
    "-d", 
    "-p 80:80", 
    "--restart always"
  ],
  "envVars": [
    "db_host=some_host",
    "db_pwd=$SOME_PWD"
  ]
}
```
