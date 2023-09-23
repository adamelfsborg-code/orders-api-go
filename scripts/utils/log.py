import datetime
import json
import sys
from .codes import code_to_color, code_to_status
from pygments import highlight, lexers, formatters


def Log(*data: str, code: int, custom_status_text=None):
    current_datetime = datetime.datetime.now()
    formatted_datetime = current_datetime.strftime("%Y-%m-%d %H:%M:%S")
    color = code_to_color(code=code)
    status_text = code_to_status(code=code) if not custom_status_text else custom_status_text

    divider=f"----------------------------------------- {status_text} -----------------------------------------"
    colored_divider=f"{color}{divider}\x1b[0m"
    print(colored_divider)
    print(
        f"[Date]: {formatted_datetime}", 
        *data, 
        sep="\n", 
        end=f"\n", 
        flush=True
    )
    print(colored_divider)

def LogEx(*data: str, code: int):
    code = 500 if code == 1 else code
    Log(*data, code=code)
    sys.exit(code)


def LogJson(data: object, code=0):
    formatted_data = json.dumps(data, indent=4)
    colorized_json = highlight(
        formatted_data, 
        lexers.JsonLexer(), 
        formatters.TerminalFormatter()
    )
    Log(colorized_json, code=code)