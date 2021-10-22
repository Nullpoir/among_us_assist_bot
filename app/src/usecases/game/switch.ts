import { getGameVoiceChannel, WILL_MUTE, WILL_UNMUTE } from './common.js'

const mute = async (members: any) => {
  await Promise.all(members.map((member: any) => member.edit({ mute: true })))
}
const unmute = async (members: any) => {
  await Promise.all(members.map((member: any) => member.edit({ mute: false })))
}
const updateState = async (rc:any, key: string) => {
  return new Promise(async (resolve, reject) => {
    let result = await rc.get(`mHiyori:gaming state ${key}`)
    if (result === null) {
      await rc.set(`mHiyori:gaming state ${key}`, WILL_MUTE)
      resolve(WILL_MUTE);
    } else {
      await rc.set(`mHiyori:gaming state ${key}`, result === WILL_MUTE ? WILL_UNMUTE : WILL_MUTE)
      resolve(result)
    }
    reject(null)
  })
}

exports.switch = async(msg: any, rc: any) => {
  let start: any = Date.now()
  let gameVC = getGameVoiceChannel(msg.channel.parent.children)
  let categoryName = msg.channel.parent.name
  let state = await updateState(rc, categoryName)
  if(state === WILL_MUTE) {
    await mute(gameVC.members)
    let duration = Date.now() - start
    await msg.channel.send(`ミュートしました。(${duration}ms)`)
  } else {
    await unmute(gameVC.members)
    let duration = Date.now() - start
    await msg.channel.send(`議論してください！(${duration}ms)`)
  }
}
