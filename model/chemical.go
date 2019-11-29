package model

import (
    "log"

    "github.com/jinzhu/gorm"
)

type Chemical struct {
    gorm.Model
    Name            string `json:"name"`
    DescriptionLink string `json:"description_link"`
    Slug            string `json:"slug"`
}

func GetChemical(slug string) (*Chemical) {
    chemical := &Chemical{}

    if err := GetDatabase().Table("chemicals").Where("slug = ?", slug).First(chemical).Error; err != nil {
        log.Println(err.Error())
        return nil
    }

    return chemical
}

func GetAllChemicals() ([]*Chemical) {
    chemicals := make([]*Chemical, 0)
    
    if err := GetDatabase().Table("chemicals").Order("Name").Find(&chemicals).Error; err != nil {
        log.Println(err.Error())
        return nil
    }

    return chemicals
}