import config
import os
import pickledb


db = pickledb.load(os.getcwd() + config.database, False)


def store_service_ip(service_name, new_ip):
    ip_key = "server_ip"
    service_info = get_value(service_name)
    if isinstance(service_info, dict):
        service_info[ip_key] = new_ip
        set_value(service_name, service_info)
    elif service_info is None:
        set_value(service_name, {ip_key: new_ip})


def save_db():
    db.dump()


def set_value(key, value):
    db.set(key, value)


def get_value(key):
    return db.get(key)
