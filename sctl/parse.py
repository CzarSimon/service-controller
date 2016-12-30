import json
import os
import re
from sctl import db


def service_cmd(name):
    service_config = _get_config(_create_path("services", name))
    commands = ["docker pull " + service_config["image"]]
    commands.append("docker kill " + name)
    commands.append("docker rm " + name)
    commands.append(run_cmd(service_config, name))
    return commands, service_config["dependencies"]


def run_cmd(service_config, name):
    cmd_list = [service_config["command"]]
    cmd_list.append("--name {}".format(name))
    cmd_list.append(_get_kwags(service_config))
    cmd_list.append(_get_environment_vars(service_config))
    cmd_list.append(service_config["image"])
    return " ".join(cmd_list)


def _get_kwags(config):
    kwarg_list = config["keyword_args"]
    return " ".join(_parse_vars(kwarg_list))


def _get_environment_vars(config):
    env_variables = config["env_variables"]
    env_vars = map(lambda var: "-e " + var, env_variables)
    return " ".join(_parse_vars(env_vars))


def _parse_vars(param_list):
    return map(lambda param: _add_values(param), param_list)


def _add_values(param):
    res = re.findall('\{(.*?)\}', param)
    if len(res) == 0:
        return param
    else:
        param_cmd = res[0]
        value = _get_value(param_cmd)
        return param.replace("{" + param_cmd + "}", value)


def _get_value(param_cmd):
    p = param_cmd.split(".")
    name = p[0]
    values = db.get_value(name)
    if len(p) == 2:
        return values[p[1]]
    else:
        return values


def _get_config(path):
    with open(path, 'r') as config_file:
        config = json.loads(config_file.read())
    return config


def _create_path(obj_type, name):
    return "{}/{}/{}.json".format(os.getcwd(), obj_type, name)
