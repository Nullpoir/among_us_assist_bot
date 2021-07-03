import asyncio

async def set_mute(members, state):
  tasks = [member.edit(mute=state) for member in members]
  results = await asyncio.gather(*tasks)
  return results
