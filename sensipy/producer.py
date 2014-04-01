
import time
import urllib
import urllib2
import json

class Producer(object):

    def __init__(self, addr, port):
        self.addr = addr
        self.port = port

        self.url = "http://" + addr + ":" + str(port)
        self.worker_url = self.get_url()


    def get_url(self):
        resp = urllib2.urlopen(self.url + "/url")
        data = json.loads(resp.read())
        return data["address"]


    def send_sensor_data(self, id_str, latitude, longitude, data):
        post_dictionary = {
            "id": id_str,
            "timestamp": time.time(),
            "latitude": latitude,
            "longitude": longitude,
            "data": data
        }

        post_dictionary_encode = urllib.urlencode(post_dictionary)

        req = urllib2.Request(
            self.worker_url + "/sensors", post_dictionary_encode
        )

        resp = urllib2.urlopen(req)


if __name__ == "__main__":
    pr = Producer("localhost", 8000)
    pr.send_sensor_data(
        "python", 56.339892299999995, -2.8094739, {"test": 0}
    )

