import redis
from settings.settings import *

rc = redis.Redis(host=REDIS_HOST, port=REDIS_PORT)
