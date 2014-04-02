
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
        if self.count % 1 == 0:
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

        time_start = time.time()
        try:
            resp = urllib2.urlopen(req)
        except:
            for i in range(100):
                try:
                    self.worker_url = self.get_url()
                except:
                    exit(0)
                try:
                    req = urllib2.Request(
                        self.worker_url + "/sensors", post_dictionary_encode
                    )
                    resp = urllib2.urlopen(req)
                    break
                except Exception as e:
                    pass
        time_diff = time.time() - time_start
        return time_diff


class ProducerTester(object):

    def __init__(self, num_runs):
        self.id_str = str(random.random())
        self.prod = Producer("localhost", 8000)
        self.num_runs = num_runs


    def run(self):

        for i in range(self.num_runs):
            print self.prod.send_sensor_data(
                self.id_str, 56.339892299999995,
                -2.8094739, {"test": 0}
            )

if __name__ == "__main__":
    pt = ProducerTester(10)
    pt.run()
