import datetime
import sys
from .codes import code_to_color, code_to_status

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

def LogEx(*data: str, custom_status_text = None, code: int):
    code = 500 if code == 1 else code
    Log(*data, custom_status_text=custom_status_text, code=code)
    sys.exit(code)