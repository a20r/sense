
import urllib2
import json

class Benchmarker(object):

    def __init__(self, addr, port):
        self.url = "http://" + addr + ":" + str(port)


    def get_stats(self):
        resp = urllib2.urlopen(self.url + "/ajax/stat")
        resp_dict = json.loads(resp.read())

        return resp_dict


if __name__ == "__main__":
    bm = Benchmarker("localhost", 8080)
    print bm.get_stats()
