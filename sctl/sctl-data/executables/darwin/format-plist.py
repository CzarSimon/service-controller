import os
import sys


# get_file_name Parsers the user supplied filename
def get_file_name():
    service = sys.argv[1]
    reverse_domain = "com." + ".".join(service.split("-")[::-1])
    return service + os.sep + reverse_domain + ".plist"


# get_plist Reads the content of the minion plist
def get_plist(file_name):
    with open(file_name, "r") as plist:
        return plist.read()


# write_plist Substitues gopath in the plist definition and writes to the plist file
def write_plist(plist, file_name):
    if reverse():
        plist = plist.replace(get_gopath(), "$GOPATH")
        plist = plist.replace(os.getlogin(), "$USER")
    else:
        plist = plist.replace("$GOPATH", get_gopath())
        plist = plist.replace("$USER", os.getlogin())
    with open(file_name, "w") as f:
        f.write(plist)


# get_gopath Returns the gopath value if it exists, calls sys exit otherwise
def get_gopath():
    if not "GOPATH" in os.environ:
        print("No gopath installed, which is reqired, exiting")
        sys.exit(1)
    return os.environ["GOPATH"]


# reverse checks if the call is a revierse request
def reverse():
    if (len(sys.argv) > 2):
        return "reverse" == sys.argv[2]
    else:
        return False


def main():
    file_name = get_file_name()
    plist = get_plist(file_name)
    write_plist(plist, file_name)


if __name__ == '__main__':
    main()
