package main

import (
    "gopkg.in/mgo.v2/bson"
    "time"
)

var timeSpents TimeSpents


func FindTimeSpent(date string) (error, TimeSpents) {
    err := a.Session.DB(a.Database).C(a.Collection).Find(bson.M{"date": date}).All(&timeSpents)

    if timeSpents == nil{
        return err, TimeSpents{}
    }

    return err, timeSpents
}

func FindTimeSpentByDateAndUser(date string, user string) (error, TimeSpents) {
    err := a.Session.DB(a.Database).C(a.Collection).Find(bson.M{"date": date, "user": user}).All(&timeSpents)

    if timeSpents == nil{
        return err, TimeSpents{}
    }

    return err, timeSpents
}


func CreateTimeSpent(t TimeSpent) (error, TimeSpent) {
    t.ID = bson.NewObjectId()
    t.CreatedAt = time.Now()
    if len(t.Date) == 0 {
        t.Date = time.Now().Format("2006-01-02")
    }
    err := a.Session.DB(a.Database).C(a.Collection).Insert(t)

    return err, t
}

func DestroyTimeSpent(id string) (error, bool) {
    err := a.Session.DB(a.Database).C(a.Collection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})

    return err, true
}

func ClearUserTimeSpentByDay(date string, user string) (error) {
    _, err := a.Session.DB(a.Database).C(a.Collection).RemoveAll(bson.M{"date": date, "user": user})
    
    return err
}