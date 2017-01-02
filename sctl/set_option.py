from . import db, arguments
import json


_help_message = "Set option is used to set config values"
_error_message = "Wrong number of arguments, set takes either a single key or nothing in which case a prompt will appear"

def set_values(full_args):
    run_help, args = _resolve_flags(full_args)
    if run_help:
        print(_help_message)
    else:
        args_length = len(args)
        if args_length == 0:
            _set_value_prompt()
        elif args_length == 1:
            key = args[0]
            _set_values(key)
        else:
            print(_error_message)


def _set_value_prompt():
    while True:
        key = raw_input("Set value for key (type exit to exit): ")
        if key == "exit":
            break
        _set_values(key)


def _set_values(key):
    data = raw_input("Value of {}: ".format(key))
    try:
        value = json.loads(data)
        config_key = "config/" + key
        db.set_value(config_key, value)
    except ValueError as e:
        print("Values must be in json formats")


def _resolve_flags(args):
    flags = {
        "help": "--help"
    }
    run_help = False if (flags["help"] not in args) else True
    return run_help, arguments.remove_flags(args, flags)
