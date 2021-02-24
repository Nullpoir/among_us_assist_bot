from helpers.talks import *

class TestTalks:
  def test_create_response(self):
    response_list = {
        '/hiyochi': 'に、にゃーん・・・///',
        '/mugitea': 'ぐふっ・・・!',
        '𝓜𝓾𝓰𝓲 𝓣𝓮𝓪': 'ぶはっ！'
    }
    for t, a in response_list.items():
        assert a == create_response(t)
