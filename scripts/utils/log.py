import datetime
import sys
from .codes import color_code

def Log(*data: str, code: int):
    current_datetime = datetime.datetime.now()
    formatted_datetime = current_datetime.strftime("%Y-%m-%d %H:%M:%S")
    color = color_code(code=code)
    divider=f"{color}----------------------------------------------------"
    print(divider)
    print(
        f"[DATE]: {formatted_datetime}", 
        *data, 
        sep=".\n", 
        end=f".\n{divider}\n", 
        flush=True
    )

def LogEx(*data: str, exit_code=1):
    if exit_code >= 1:
        exit_code = 1
        code = 500
    else:
        exit_code = 0
        code = 200
    Log(*data, code=code)
    sys.exit(exit_code)

