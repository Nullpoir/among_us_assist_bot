response_list = {
  '/hiyochi': 'に、にゃーん・・・///',
  '/mugitea': 'ぐふっ・・・!',
  '𝓜𝓾𝓰𝓲 𝓣𝓮𝓪': 'ぶはっ！'
}

def create_response(message):
  return response_list.get(message)
