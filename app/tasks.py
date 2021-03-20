from celery import Celery

celery_app = Celery(broker='redis://localhost')

@celery_app.task(serializer='json')
async def mute(members):
  for member in members:
    await member.edit(mute=True)
  return {"status": True}

@celery_app.task(serializer='json')
async def unmute(members):
  for member in members:
    await member.edit(mute=False)
  return {"status": True}
