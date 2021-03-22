async def mute(members):
  for member in members:
    await member.edit(mute=True)
  return {"status": True}
