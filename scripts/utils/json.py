import json
from pygments import highlight, lexers, formatters


def format_json(data: object):
    formatted_data = json.dumps(data, indent=4)
    colorized_json = highlight(
        formatted_data, 
        lexers.JsonLexer(), 
        formatters.TerminalFormatter()
    )
    return colorized_json