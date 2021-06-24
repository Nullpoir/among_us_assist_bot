import { WILL_MUTE } from './common.js'

const clearState = async (rc: any, key: string) => {
  return new Promise(async (resolve, reject) => {
    let result = await rc.get(`mHiyori:gaming state ${key}`)
    if (result !== false) {
      await rc.set(`mHiyori:gaming state ${key}`, WILL_MUTE)
      resolve(WILL_MUTE)
    } else {
      reject(null)
    }
  })
}

exports.clear = async (msg: any, rc: any) => {
  let categoryName = msg.channel.parent.name
  await clearState(rc, categoryName)
}
