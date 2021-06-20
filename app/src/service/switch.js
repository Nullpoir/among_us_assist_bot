const WILL_MUTE = '1'
const WILL_UNMUTE = '0'

const getGameVoiceChannel = (channels) => {
  const gameChannel = channels.find(
    (channel) => {
      return channel.name === 'game'
    }
  )
  return gameChannel
}
const mute = async (members) => {
  members.each(
    (member) => {
      member.edit({mute:true})
    }
  )
}
const unmute = async (members) => {
  members.each(
    (member) => {
      member.edit({ mute: false })
    }
  )
}
const updateState = async (rc, key) => {
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
  })
}

exports.switch = async(msg, rc) => {
  gameVC = getGameVoiceChannel(msg.channel.parent.children)
  categoryName = msg.channel.parent.name
  let state = await updateState(rc, categoryName)
  console.log(state)
  if(state === WILL_MUTE) {
    await mute(gameVC.members)
  } else {
    await unmute(gameVC.members)
  }
}
