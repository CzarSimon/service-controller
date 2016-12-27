import db
import interactive
import remote
import sys


def main():
    args = sys.argv[1:]
    if args[0] == "start":
        remote.service_start(args[1], args[3])
    elif args[1] == "values":
        if args[0] == "set":
            interactive.set_values()
        elif args[0] == "get":
            interactive.get_values()
    db.save_db()


if __name__ == '__main__':
    main()
