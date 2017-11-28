package main

import "github.com/jinzhu/gorm"

type Talk struct {
	gorm.Model
	Msg string
}
