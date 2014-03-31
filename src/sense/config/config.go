package config

import (
    "time"
)

/* Broker URLs */
var BrokerUrl string = "http://localhost:8000"
var HeartbeatRoute string = "/heartbeat"
var BrokerHeartbeatUrl string = BrokerUrl + HeartbeatRoute

/* Time needed to remove worker */
var WorkerTimeRemove int64 = 5
var TimeDelayRemoveCheck time.Duration = 5000 * time.Millisecond

/* Worker Routes */
var WorkerSensorRoute string = "/sensors"
