import datetime
import sys

def Log(*data: str):
    current_datetime = datetime.datetime.now()
    formatted_datetime = current_datetime.strftime("%Y-%m-%d %H:%M:%S")
    print(f"[DATE]: {formatted_datetime}", *data, sep=" \n", end=".\n", flush=True)

def LogEx(*data: str, exit_code=1):
    Log(*data)
    sys.exit(exit_code)