#service controller
Simple command line tool for running docker containers on servers.
Uses service specific configurations (such as container and image name) specified
in JSON format and config data necessary for multiple services stored in a simple
pickledb key-value JSON store.

##Run as follows:
(all commands preceded by python sctl.py)

**start *service_name* on *server_name***
Starts a service defined with the specified name on the specified server.
Gets service data from JSON file stored under a services folder as well as
the JSON key-value store for shared config data. Creates docker commands for
starting a container (pull image, kill and remove exiting service with same name
and run new container). Enters specified server via ssh and runs these commands.

On successful startup the ip address of the server where the service is started
is stored and like so <service_name>: { "server_ip": <ip_of_server> ... }

**set values**
Displays a prompt which allows the user to set config values to be used.
E.g.
Set value for key: postgres
Value of postgres: {"ip": <ip>, "password": <pwd>}

**get values**
Displays a prompt which allows the user view set config values.
E.g.
Get value for key: postgres
{"ip": <ip>, "password": <pwd>}

###Service config format
```json
{
  "name": "<some_name>",
  "command": "docker run",
  "keyword_args": [
    "-t", "-d", "--network={<some_network>.name}", "--restart always"
  ],
  "env_variables": [
    "db_host={postgres.ip}",
    "db_pwd={postgres.pwd}"
  ],
  "image": "<some_image>"
}
```
