import db
import config
import json
import paramiko
import parse
import sys


def main():
    args = sys.argv[1:]
    if args[0] == "start":
        _service_start(args[1], args[3])
    elif args[1] == "values":
        if args[0] == "set":
            _set_values()
        elif args[0] == "get":
            _get_values()
    db.save_db()


def _service_start(service_name, server_name):
    server_ip = db.get_value(server_name)["ip"]
    print("On server {} with ip {} run:".format(server_name, server_ip))
    commands = parse.service_cmd(service_name)
    for command in commands:
        print(command)
    db.store_service_ip(service_name, server_ip)


def _set_values():
    while True:
        key = raw_input("Set value for key (type exit to exit): ")
        if key == "exit":
            break
        data = raw_input("Value of {}: ".format(key))
        try:
            value = json.loads(data)
        except ValueError as e:
            value = data
        db.set_value(key, value)


def _get_values():
    while True:
        key = raw_input("Get value for key (type exit to exit): ")
        if key == "exit":
            break
        print("Value of {} is: ".format(key))
        print(db.get_value(key))


def _check_value_type(value_type):
    return value_type in config.value_types


if __name__ == '__main__':
    main()
