import { getGameVoiceChannel, WILL_MUTE, WILL_UNMUTE } from './common'

const mute = async (members: any, mute: boolean) => {
  return Promise.all(
    members.map(
      (member: any) => {
        member.edit({ mute: mute })
      }
    )
  )
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
    await mute(gameVC.members, true)
    let duration = Date.now() - start
    await msg.channel.send(`ミュートしました。(${duration}ms)`)
  } else {
    await mute(gameVC.members, false)
    let duration = Date.now() - start
    await msg.channel.send(`議論してください！(${duration}ms)`)
  }
}
