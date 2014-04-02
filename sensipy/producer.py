
import time
import urllib
import urllib2
import json
import random
import threading

class Producer(object):

    def __init__(self, addr, port):
        self.addr = addr
        self.port = port

        self.url = "http://" + addr + ":" + str(port)
        self.worker_url = self.get_url()
        self.count = 0


    def get_url(self):
        resp = urllib2.urlopen(self.url + "/url")
        data = json.loads(resp.read())
        return data["address"]


    def send_sensor_data(self, id_str, latitude, longitude, data):
        self.count += 1
        if self.count % 10 == 0:
            self.worker_url = self.get_url()

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


class ProducerTester(object):

    def __init__(self, num_runs):
        self.id_str = str(random.random())
        self.prod = Producer("localhost", 8000)
        self.num_runs = num_runs


    def run(self):

        for i in range(self.num_runs):
            self.prod.send_sensor_data(
                self.id_str, 56.339892299999995,
                -2.8094739, {"test": 0}
            )

if __name__ == "__main__":
    pt = ProducerTester(1000)
    pt.run()
