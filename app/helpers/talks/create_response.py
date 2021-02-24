def create_response(message):
  response_list = {
    '/hiyochi': 'に、にゃーん・・・///',
    '/mugitea': 'ぐふっ・・・!',
    '𝓜𝓾𝓰𝓲 𝓣𝓮𝓪': 'ぶはっ！'
  }

  return response_list.get(message)
