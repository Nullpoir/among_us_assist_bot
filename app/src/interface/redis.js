const redis = require('async-redis');
const env = require('../settings.js');

exports.client = redis.createClient(env.REDIS_PORT, env.REDIS_HOST);
