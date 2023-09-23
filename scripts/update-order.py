import requests
import sys
from utils.enviroment import SERVER_ADDR
from utils.log import LogEx

try:
    order_id = sys.argv[1]
    order_status = sys.argv[2]
    if order_status not in ["shipped", "completed"]:
        LogEx(f"[ValueError]: Wrong Status passed", "[Usage]: python script.py <integer_value: Order Id> <string: Status>", code=1)
except ValueError as err:
    LogEx(f"[ValueError]: {err}", "[Usage]: python script.py <integer_value: Order Id> <string: Status>", code=1)
except IndexError as err:
    LogEx(f"[IndexError]: {err}", "[Usage]: python script.py <integer_value: Order Id> <string: Status>", code=1)


order = {
    "status": order_status,
}

r = requests.put(f"{SERVER_ADDR}/orders/{order_id}", json=order)
LogEx(f"[Code]: {r.status_code}", code=r.status_code)