import os
from dotenv import load_dotenv

current_directory = os.path.dirname(os.path.abspath(__file__))
path=os.path.join(current_directory, '../..', '.env')
load_dotenv(dotenv_path=path)

SERVER_ADDR = os.getenv('SERVER_ADDR')
