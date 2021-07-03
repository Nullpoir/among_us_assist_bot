from interface import rc

def update_state(category_id):
  key = f'mHiyori:gaming state { category_id }'

  if rc.exists(key):
    state = rc.get(key)
    if state == b'1':
      rc.set(key, b'0')
      return True
    else:
      rc.set(key, b'1')
      return False
  else:
    rc.set(key, b'0')
    state = b'1'
    return True
