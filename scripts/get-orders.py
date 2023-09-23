import requests
import sys
from utils.log import Log, LogEx
from utils.json import format_json


try:
    if len(sys.argv) > 1:
        page = sys.argv[1]
    else:
        page = None
except ValueError as err:
    LogEx(
        f"[ValueError]: {err}", 
        "[Usage]: python script.py <integer_value: Page>", 
        code=1
    )

url = f"http://localhost:3000/orders{f'?page={page}' if page else '/'}"
r = requests.get(url=url)
try:
    json = r.json()
    Log(format_json(json), custom_status_text=r.status_code, code=r.status_code)
except Exception as err:
    LogEx(
        f"[Error]: {err}", 
        code=1
    )