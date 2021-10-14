export const WILL_MUTE = '1'
export const WILL_UNMUTE = '0'

export const getGameVoiceChannel = (channels: any) => {
  const gameChannel = channels.find(
    (channel: any) => {
      return channel.name === 'game'
    }
  )
  return gameChannel
}
