import json
from . import db, types, arguments


_help_message = "Was this helpfull?"
_error_message = "Do it right instead"


def menu(full_args):
    add_host, run_help, args = _resolve_flags(full_args)
    args_length = len(args)
    if run_help:
        print(_help_message)
    else:
        if args_length == 1:
            _add_type(args[0], add_host)
        elif args_length == 2:
            _add_type(args[0], add_host, args[1])
        elif args_length == 0:
            obj_type = raw_input("What type?, options {}: ".format(types.get_types()))
            _add_type(obj_type, add_host)
        else:
            print(_error_message)


def _add_type(obj_type, add_host, name=None):
    if name is None:
        name = raw_input("Name of {}: ".format(obj_type))
    type_def = types.get_type(obj_type)
    obj = {}
    for attr, attr_type in type_def.iteritems():
        if attr == "name":
            obj[attr] = name
            print("name: {}".format(obj[attr]))
        else:
            obj[attr] = raw_input("{} (type: {}): ".format(attr, attr_type))
    key = "{}/{}".format(obj_type, name)
    if _confirm(key, obj):
        _save_value(key, obj)


def _confirm(key, value):
    print("Saving {} as: ".format(key))
    print(json.dumps(value, sort_keys=True, indent=4))
    return "y" == raw_input("Look ok? [y/n]: ").lower()


def _save_value(key, value):
    try:
        db.set_value(key, value)
    except TypeError as e:
        print(e)


def _resolve_flags(args):
    flags = {
        "help": "--help",
        "dont_add_host": "-n"
    }
    add_host = False if (flags["dont_add_host"] in args) else True
    run_help = True if (flags["help"] in args) else False
    return add_host, run_help, arguments.remove_flags(args, flags)
