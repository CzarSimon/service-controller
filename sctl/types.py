def get_type(t):
    return _types[t](t)


def get_types():
    return _types.keys()


_types = {
    "service": lambda t: service(),
    "server": lambda t: server(),
    "network": lambda t: network()
}


# service
def service():
    return {
        "name": "str",
        "keyword_args": "list",
        "env_vars": "list",
        "image": "str",
        "dependencies": "list",
        "command": "str"
    }


# server
def server():
    return {
        "name": "str",
        "ip": "str"
    }


# network
def network():
    return {
        "name": "str",
        "subnet": "str",
        "gateway": "str",
        "type": "str"
    }
