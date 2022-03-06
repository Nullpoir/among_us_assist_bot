const Discord = require('discord.js')
import { ACCESS_TOKEN } from './settings'
const redisClient = require('./infra/redis.js').client;
const usecases = require('./usecases/usecases');

const client = new Discord.Client()

client.on('ready', () => {
  console.log(`ready with ${client.user.tag} `)
})

client.on('message', async (message: any) => {
  if (message.content === 'm') {
    await usecases.switch(message,redisClient)
  } else if (message.content === 'c') {
    await usecases.clear(message, redisClient)
  }
})

client.on('disconnect', () => {
  console.log('bye')
  redisClient.quit()
})

client.login(ACCESS_TOKEN)
