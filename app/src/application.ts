const Discord = require('discord.js')
import { ACCESS_TOKEN } from './settings'
const redisClient = require('./interface/redis.js').client;
const service = require('./service/service');

const client = new Discord.Client()

client.on('ready', () => {
  console.log(`ready with ${client.user.tag} `)
})

client.on('message', async (message: any) => {
  if (message.channel.name === 'bot操作' || message.channel.name === 'チャット'){
    if (message.content === 'm') {
      service.switch(message, redisClient)
    } else if (message.content === 'c') {
      service.clear(message, redisClient)
    }
  }
})

client.on('disconnect', () => {
  console.log('bye')
  redisClient.quit()
})

client.login(ACCESS_TOKEN)
