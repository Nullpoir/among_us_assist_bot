const Discord = require('discord.js')
const env = require('./settings.js')
const redisClient = require('./interface/redis.js').client;
const service = require('./service/service.js');

const client = new Discord.Client()

client.on('ready', () => {
  console.log(`${client.user.tag} でログインしています。`)
})

client.on('message', async message => {
  if (message.content === 'm') {
    service.switch(message,redisClient)
  }
})

client.on('disconnect', () => {
  console.log('bye')
  redisClient.quit()
})

client.login(env.ACCESS_TOKEN)
