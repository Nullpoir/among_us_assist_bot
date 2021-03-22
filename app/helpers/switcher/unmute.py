async def unmute(members):
  for member in members:
    await member.edit(mute=False)
  return {"status": True}
