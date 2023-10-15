import os
from os.path import join, dirname
from dotenv import load_dotenv

load_dotenv(verbose=True)

dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path)

ACCESS_TOKEN = os.environ.get("ACCESS_TOKEN")

KVS_HOST = os.environ.get("KVS_HOST")
KVS_PORT = int(os.environ.get("KVS_PORT"))
