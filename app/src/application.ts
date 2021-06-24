const Discord = require('discord.js')
import { ACCESS_TOKEN } from './settings'
const redisClient = require('./interface/redis.js').client;
const service = require('./service/service');

const client = new Discord.Client()

client.on('ready', () => {
  console.log(`${client.user.tag} でログインしています。`)
})

client.on('message', async (message: any) => {
  if (message.content === 'm') {
    service.switch(message,redisClient)
  }
})

client.on('disconnect', () => {
  console.log('bye')
  redisClient.quit()
})

client.login(ACCESS_TOKEN)
