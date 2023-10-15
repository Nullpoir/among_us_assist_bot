import redis
from settings.settings import *

rc = redis.Redis(host=KVS_HOST, port=KVS_PORT)
