import requests
import sys

try:
    order_id = sys.argv[1]
    order_status = sys.argv[2]
except ValueError as err:
    print("Usage: python script.py <integer_value> <string>")
    sys.exit(1)

order = {
    "status": order_status,
}

r = requests.put(f"http://localhost:3000/orders/{order_id}", json=order)
print("Response", r.status_code)