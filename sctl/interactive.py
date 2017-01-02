from . import db
import json


def set_values(args):
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


def get_values(args):
    if len(args) == 1:
        key = args[0]
        _print_value(key)
    elif len(args) == 0:
        while True:
            key = raw_input("Get value for key (type exit to exit): ")
            if key == "exit":
                break
            _print_value(key)
    else:
        print("To many arguments, get can either be called with a key or without a key, in wihch case promt will appear.")


def _print_value(key):
    print("Value of {} is: ".format(key))
    print(json.dumps(db.get_value(key), indent=4))
