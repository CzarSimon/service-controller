import db
import paramiko
import parse


def _run_commands(commands, server_ip):
    ssh = _connect_to_server(server_ip)
    print(_check_dependecy(ssh, "volume-counter"))
    for command in commands:
        output = _get_output(ssh, command)
        _print_result(command, output)


def _check_dependecy(ssh, dependecy):
    running_services = _get_output(ssh, "docker ps")
    dep_fulfilled = False
    for service_str in running_services:
        if dependecy in service_str:
            dep_fulfilled = True
    return dep_fulfilled


def _print_result(command, outputs):
    print("Running command: $ {}".format(command))
    print("---- Response ----")
    for output in outputs:
        print(output.replace("\n", ""))
    print(" ")


def _get_output(ssh, command):
    stdin, stdout, stderr = ssh.exec_command(command)
    return stdout.readlines()


def service_start(service_name, server_name):
    server_ip = db.get_value(server_name)["ip"]
    commands = parse.service_cmd(service_name)
    _run_commands(commands, server_ip)
    db.store_service_ip(service_name, server_ip)


def _connect_to_server(server_ip):
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(server_ip, username='simon')
    return ssh
