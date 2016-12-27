import db
import paramiko
import parse


def service_start(service_name, server_name, run_if_missing):
    commands, dependecies = parse.service_cmd(service_name)
    server_ip = _run_commands(commands, server_name, dependecies, run_if_missing)
    db.store_service_ip(service_name, server_ip)


def _run_commands(commands, server_name, dependecies, run_if_missing):
    server_ip = db.get_value(server_name)["ip"]
    ssh = _connect_to_server(server_ip)
    _check_dependecies(ssh, dependecies, run_if_missing, server_name)
    for command in commands:
        output = _get_output(ssh, command)
        _print_result(command, output)
    return server_ip


def _check_dependecies(ssh, dependecies, run_if_missing, server_name):
    running = _get_output(ssh, "docker ps")  # Gets running docker containers
    missing = filter(lambda dep: not _check_dependecy(running, dep), dependecies)
    if len(missing) > 0:
        if run_if_missing:
            print("running missing services: {}".format(" ".join(missing)))
            _run_missing_dependecies(missing, server_name)
        else:
            print("these services are missing: {}".format(" ".join(missing)))


def _run_missing_dependecies(missing_deps, server_name):
    for missing_dependency in missing_deps:
        service_start(missing_dependency, server_name, True)


def _check_dependecy(running_services, dependecy):
    for service_str in running_services:
        if dependecy in service_str:
            return True
    return False


def _print_result(command, outputs):
    print("Running command: $ {}".format(command))
    print("---- Response ----")
    for output in outputs:
        print(output.replace("\n", ""))
    print(" ")


def _get_output(ssh, command):
    stdin, stdout, stderr = ssh.exec_command(command)
    return stdout.readlines()


def _connect_to_server(server_ip):
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(server_ip, username='simon')
    return ssh
