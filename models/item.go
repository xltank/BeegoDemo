package models

import "time"

type Item struct {
    ID int
    Name string
    Price float32
    ProcudeDate time.Time
}

func FindOne() Item{
    t, _ := time.Parse(time.RFC3339, "2017-05-10T12:34:56Z08:00")
    return Item{111, "Test", 3020.13, t}
}