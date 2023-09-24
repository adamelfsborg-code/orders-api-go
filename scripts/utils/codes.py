CODES_TO_COLOR: dict[int, str] = {
    0: "\033[0m",
    200: "\033[92m",
    400: "\033[91m",
}

CODES_TO_STATUS_TEXT: dict[int, str] = {
    0: "INFO",
    200: "SUCCESS",
    400: "ERROR",
}

def extrat_from_code_dict(dict: dict, code: int) -> any:
    success_term = 200 if code in range(200, 299) else 0
    error_term = 400 if code in range(400, 599) else 0
    
    if success_term > 0:
        return dict[success_term]
    
    if error_term > 0:
        return dict[error_term]
    
    return dict[0]

def code_to_color(code: int) -> str:
    return extrat_from_code_dict(dict=CODES_TO_COLOR, code=code)

def code_to_status(code: int) -> str:
    return extrat_from_code_dict(dict=CODES_TO_STATUS_TEXT, code=code)