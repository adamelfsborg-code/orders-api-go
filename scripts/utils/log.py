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

def LogEx(*data: str, code: int):
    code = 500 if code == 1 else code
    Log(*data, code=code)
    sys.exit(code)

