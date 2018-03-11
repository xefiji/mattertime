package main

import (
    "gopkg.in/mgo.v2/bson"
)

var timeSpents TimeSpents


func RepoFindTimeSpent(date string) (error, TimeSpents) {
    err := a.Session.DB(a.Database).C(a.Collection).Find(bson.M{"date": date}).All(&timeSpents)

    if timeSpents == nil{
        return err, TimeSpents{}
    }

    return err, timeSpents
}

func RepoCreateTimeSpent(t TimeSpent) (error, TimeSpent) {
    t.ID = bson.NewObjectId()
    err := a.Session.DB(a.Database).C(a.Collection).Insert(t)

    return err, t
}

func RepoDestroyTimeSpent(id string) (error, bool) {
    err := a.Session.DB(a.Database).C(a.Collection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})

    return err, true
}