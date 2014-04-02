
import urllib
import urllib2
import json
import time

class Client(object):

    def __init__(self, addr, port):
        self.addr = addr
        self.port = port
        self.url = "http://" + addr + ":" + str(port)

    def get_sensor_data(self, lat, lon, rad):
        start_time = time.time()
        req_url = self.url + "/client/" + "/".join(map(str, [lat, lon, rad]))
        req = urllib2.Request(req_url)

        try:
            resp = urllib2.urlopen(req)
            time_diff = time.time() - start_time
            print time_diff
            resp_str = resp.read()
            resp_dict = json.loads(resp_str)
        except:
            return list()

        return resp_dict


class ClientTester(object):

    def __init__(self, num_runs):
        self.cl = Client("localhost", 8000)
        self.num_runs = num_runs


    def run(self):

        for i in range(self.num_runs):
            self.cl.get_sensor_data(56.339892299999995, -2.8094739, 3)


if __name__ == "__main__":
    ct = ClientTester(10)
    ct.run()

