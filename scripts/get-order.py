import requests
import sys
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
    r = requests.get(f'http://localhost:3000/orders/{order_id}')
    try:
        json = r.json()
        Log(format_json(json), custom_status_text=r.status_code, code=r.status_code)
    except Exception as err:
        LogEx(
            f"[Error]: {err}", 
            code=1
        )
else:
    LogEx(
        f"[ValueError]: {err}", 
        "[Usage]: python script.py <integer_value: Order Id>", 
        code=1
    )