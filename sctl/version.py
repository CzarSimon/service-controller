import json
from . import config

_version = [0, 0, 20]

__version__ = ".".join(map(lambda num: str(num), _version))
__pkg_name__ = config.tool_name

def print_verison():
    print("{} v{}".format(__pkg_name__, __version__))
