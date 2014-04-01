package util

import (
    r "../../github.com/christopherhesse/rethinkgo"
    "encoding/json"
    "math"
)

type SensorDB struct {
    Name string
}

type SensorData struct {
    Id        string
    Timestamp string
    Latitude  float64
    Longitude float64
    Data      string
}

type SensorDataRow []SensorData

var DbPort string = ":28015"
var DbName string = "sense"
var TSpec r.TableSpec = r.TableSpec{Name: "sensors", PrimaryKey: "id"}

func MakeSensorDB(addr string) SensorDB {
    db := SensorDB{addr}
    db.Initialize()
    return db
}

func (sdb SensorDB) Connect() (*r.Session, error) {
    session, err := r.Connect(sdb.Name+DbPort, DbName)
    return session, err
}

func (sdb SensorDB) Create() error {
    session, err := sdb.Connect()
    if err != nil {
        return err
    } else {
        r.DbCreate(DbName).Run(session).Exec()
        r.Db(DbName).TableCreateWithSpec(TSpec).Run(session).Exec()
        return nil
    }
}

func (sdb SensorDB) Initialize() (bool, error) {
    var db_list []string
    var err error

    session, err := sdb.Connect()

    if err != nil {
        return false, err
    }

    err = r.DbList().Run(session).All(&db_list)

    if err != nil {
        return false, err
    }

    for _, db := range db_list {
        if DbName == db {
            return false, nil
        }
    }

    sdb.Create()

    return true, nil
}

func (sdb SensorDB) Insert(sd SensorData) error {
    session, _ := sdb.Connect()
    var response r.WriteResponse
    err := r.Db(DbName).Table(TSpec.Name).Insert(
        sd.ToMap(),
    ).Overwrite(true).Run(session).One(&response)
    return err
}

func (sdb SensorDB) GetNear(lat, lon, rad float64) SensorDataRow {
    var sd_list SensorDataRow

    session, _ := sdb.Connect()
    r.Db(DbName).Table(TSpec.Name).Run(session).All(&sd_list)

    filtered_sd_list := make(SensorDataRow, 0)
    for _, sd := range sd_list {
        if sd.GetDistance(lat, lon) < rad {
            filtered_sd_list = append(filtered_sd_list, sd)
        }
    }

    return filtered_sd_list
}

func (sd SensorData) ToMap() r.Map {
    return r.Map{
        "Id":        sd.Id,
        "Timestamp": sd.Timestamp,
        "Latitude":  sd.Latitude,
        "Longitude": sd.Longitude,
        "Data":      sd.Data,
    }
}

func (sd SensorData) GetDistance(lat, lon float64) float64 {
    R := float64(6371)
    dLat := deg2rad(lat - sd.Latitude)
    dLon := deg2rad(lon - sd.Longitude)
    a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(sd.Latitude))*
        math.Cos(deg2rad(lat))*math.Sin(dLon/2)*math.Sin(dLon/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
    d := R * c
    return d
}

func (sd SensorData) String() (s string) {
    b, err := json.Marshal(sd)
    if err != nil {
        s = ""
        return
    }
    s = string(b)
    return
}

func (sdr SensorDataRow) String() (s string) {
    b, err := json.Marshal(sdr)

    if err != nil {
        s = ""
        return
    }

    s = string(b)
    return
}

func deg2rad(deg float64) float64 {
    return deg * (math.Pi / 180)
}
