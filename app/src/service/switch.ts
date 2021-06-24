const WILL_MUTE = '1'
const WILL_UNMUTE = '0'

const getGameVoiceChannel = (channels: any) => {
  const gameChannel = channels.find(
    (channel: any) => {
      return channel.name === 'game'
    }
  )
  return gameChannel
}
const mute = async (members: any) => {
  members.each(
    (member: any) => {
      member.edit({mute:true})
    }
  )
}
const unmute = async (members: any) => {
  members.each(
    (member: any) => {
      member.edit({ mute: false })
    }
  )
}
const updateState = async (rc:any, key: string) => {
  return new Promise(async (resolve, reject) => {
    let result = await rc.get(`mHiyori:gaming state ${key}`)
    console.log(result)
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
  let gameVC = getGameVoiceChannel(msg.channel.parent.children)
  let categoryName = msg.channel.parent.name
  let state = await updateState(rc, categoryName)
  console.log(state)
  if(state === WILL_MUTE) {
    await mute(gameVC.members)
  } else {
    await unmute(gameVC.members)
  }
}
