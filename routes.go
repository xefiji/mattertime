package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "GetTimeToday",
        "GET",
        "/time",
        GetTimeToday,
    },
    Route{
        "GetTimeDay",
        "GET",
        "/time/{date}",
        GetTimeDay,
    },
    Route{
        "RecordTime",
        "POST",
        "/time",
        RecordTime,
    },
    Route{
        "DeleteTime",
        "DELETE",
        "/time/{id}",
        DeleteTime,
    },
}