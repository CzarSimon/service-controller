from os.path import expanduser

service_name = "Service controler"

description = service_name + ": a simple command line tool for running docker containers on servers"

tool_name = "sctl"

basepath = "{}/{}".format(expanduser("~"), tool_name)

database = "/data/service-config.db"

value_types = set(["service", "network", "server"])

stopwords = set(['on'])
