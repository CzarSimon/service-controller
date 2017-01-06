import json


def alter_object(old_obj, new_obj):
    for key, val in new_obj.iteritems():
        old_obj[key] = val
    return old_obj


def confirm(key, value):
    print("Saving {} as: ".format(key))
    print(json.dumps(value, sort_keys=True, indent=4))
    return "y" == raw_input("Look ok? [y/n]: ").lower()
