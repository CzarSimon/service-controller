# service controller
Command line tool for simplifying running services using docker swarm.

The goal of the tool is to not have to manually ssh into hosts and to start/stop
or alter the state of a swarm cluster and keep images on host up to date as well
as statically specify service configutations similarly to docker compose.

## Run as follows:

`$ sctl start <service-name>`

Locates the specified services configuration file and instructes the swarm
master to run the service acording to the configuration.

`$ sctl stop <service-name>`

Instructs the swarm master to stop and remove the specified service

`$ sctl update <service-name>`

Pushed the image of the supplied service to the users image registry and then
instructs all the nodes in the cluster to pull the new version.

`$ sctl init`

Starts a new project, the user will be promted to provide project details such as project name, master node ip / os and specify a local folder in which to place service configurations. (see *Service config format* heading)

`$ sctl project ls`

List all project, will specify what project is active

`$ sctl project <project-name>`

Switch active project to the specified one

`$ sctl lock`

Locks the master and minion nodes in the cluster, which results in them not accepting commands to run on their host node. This does not affect the running of the swarm, only the responsiveness on the sctl deamon on each node.

`$ sctl unlock`

Unlocks the master and minion nodes, allowing them to accept commands and run against their host nodes

## Service config format
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
