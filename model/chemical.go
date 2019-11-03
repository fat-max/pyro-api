package model

import (
    "log"
    _ "fmt"
)

type Chemical struct {
    Base
    Name string `json:"name"`
    DescriptionLink string `json:"description_link"`
    Slug string `json:"slug"`
}

func GetChemical(id string) (*Chemical) {
    chemical := &Chemical{}
    

    if err := GetDatabase().Table("chemicals").Where("id = ?", id).First(chemical).Error; err != nil {
        log.Println(err.Error())
        return nil
    }

    return chemical
}

func GetAllChemicals() ([]*Chemical) {
    chemicals := make([]*Chemical, 0)
    
    if err := GetDatabase().Table("chemicals").Find(&chemicals).Error; err != nil {
        log.Println(err.Error())
        return nil
    }

    return chemicals
}