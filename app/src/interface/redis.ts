const redis = require('async-redis');
import { REDIS_PORT, REDIS_HOST } from '../settings'

exports.client = redis.createClient(REDIS_PORT, REDIS_HOST);
