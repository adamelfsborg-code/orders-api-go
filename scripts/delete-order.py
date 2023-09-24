import requests
import sys
from utils.enviroment import SERVER_ADDR
from utils.log import LogEx

try:
    if len(sys.argv) > 1:
        order_id = sys.argv[1]
    else:
        order_id = None
except ValueError as err:
    LogEx(
        f"[ValueError]: {err}", 
        "[Usage]: python script.py <integer_value: Order Id>", 
        code=1
    )

if order_id:
    r = requests.delete(f'{SERVER_ADDR}/orders/{order_id}')
    LogEx(
        custom_status_text=r.status_code,
        code=r.status_code
    )
else:
    LogEx(
        "[Usage]: python script.py <integer_value: Order Id>", 
        code=1,
        custom_status_text="Wrong params passed"
    )