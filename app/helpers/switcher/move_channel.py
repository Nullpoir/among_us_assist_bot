async def move_channel(from_channle, to_channel):
  for member in from_channle.members:
    print(member)
    await member.move_to(to_channel)
