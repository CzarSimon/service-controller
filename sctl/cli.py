from sctl import config, db
from sctl import add_option, helper, interactive, remote
import sys


def main():
    instruction, arguments = _format_arguments(sys.argv)
    menu = {
        "start": lambda args: remote.start_service(args),
        "get": lambda args: interactive.get_values(args),
        "set": lambda args: interactive.set_values(args),
        "add": lambda args: add_option.menu(args),
        "--help": lambda args: helper.main_help()
    }
    if instruction in menu:
        try:
            menu[instruction](arguments)
            db.save_db()
        except KeyboardInterrupt:
            pass
    else:
        print("Command: {} is invalid, type sctl --help for help".format(instruction))


def _format_arguments(args):
    clean_args = filter(lambda arg: arg not in config.stopwords, args[1:])
    main_instruction = clean_args[0]
    return main_instruction, clean_args[1:]


if __name__ == '__main__':
    main()
