package storage

import (
	_ "github.com/JoseGuillen160722/GoApi"
)

type Memory struct {
	CurrentId int
	Persons   map[int]model.Person
}
