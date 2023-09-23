CODES_TO_COLOR: dict[int, str] = {
    0: "\033[0m",
    200: "\033[92m",
    400: "\033[91m",
}

def color_code(code: int) -> str:
    success_term = 200 if code in range(200, 299) else 0
    error_term = 400 if code in range(400, 599) else 0
    
    if success_term > 0:
        return CODES_TO_COLOR[success_term]
    
    if error_term > 0:
        return CODES_TO_COLOR[error_term]
    
    return CODES_TO_COLOR[0]