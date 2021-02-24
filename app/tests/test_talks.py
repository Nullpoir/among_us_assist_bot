from helpers.talks import *

class TestTalks:
  def test_create_response(self):
    for t, a in response_list.items():
        assert a == create_response(t)
