import requests
import sys
from utils.enviroment import SERVER_ADDR
from utils.log import Log, LogEx
from utils.json import format_json


try:
    if len(sys.argv) > 1:
        order_id = sys.argv[1]
    else:
        order_id = None
except ValueError as err:
    LogEx(
        f"[ValueError]: {err}", 
        "[Usage]: python script.py <integer_value: Page>", 
        code=1
    )

if order_id:
    r = requests.get(f'{SERVER_ADDR}/orders/{order_id}')
    try:
        json = r.json()
        Log(format_json(json), custom_status_text=r.status_code, code=r.status_code)
    except Exception as err:
        if r.status_code == 404:
            LogEx(
                f"[Error]: No Order with ID {order_id}", 
                custom_status_text=r.status_code,
                code=1
            )
        LogEx(
            f"[Error]: {err}", 
            custom_status_text=r.status_code,
            code=1
        )
else:
    LogEx(
        "[Usage]: python script.py <integer_value: Order Id>", 
        code=1,
        custom_status_text="Wrong params passed"
    )