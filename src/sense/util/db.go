package util

import (
    r "../../github.com/christopherhesse/rethinkgo"
)

type SensorDB struct {
    Name string
}

type SensorData struct {
    Id        string
    Timestamp string
    Latitude  float64
    Longitude float64
    Data      interface{}
}

var DbPort string = ":28015"
var DbName string = "sense"
var TSpec r.TableSpec = r.TableSpec{Name: "sensors", PrimaryKey: "id"}

func MakeSensorDB(addr string) SensorDB {
    return SensorDB{addr}
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
    return nil
}
