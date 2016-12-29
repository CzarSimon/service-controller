import db
import json


def set_values():
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


def get_values():
    while True:
        key = raw_input("Get value for key (type exit to exit): ")
        if key == "exit":
            break
        print("Value of {} is: ".format(key))
        value = json.dumps(db.get_value(key), indent=4)
        print(value)
