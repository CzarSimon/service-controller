from . import config, db
from . import add_option, helper, start_option, get_option, set_option
from . import version
import sys


def main():
    instruction, arguments = _format_arguments(sys.argv)
    menu = {
        "start": lambda args: start_option.start_service(args),
        "get": lambda args: get_option.get_values(args),
        "set": lambda args: set_option.set_values(args),
        "add": lambda args: add_option.menu(args),
        "alter": lambda args: _not_implemented(),
        "stop": lambda args: _not_implemented(),
        "remove": lambda args: _not_implemented(),
        "--version": lambda args: version.print_verison(),
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


def _not_implemented():
    print("Not implemented yet, stay tuned :)")


def _format_arguments(args):
    clean_args = filter(lambda arg: arg not in config.stopwords, args[1:])
    main_instruction = clean_args[0]
    return main_instruction, clean_args[1:]


if __name__ == '__main__':
    main()
