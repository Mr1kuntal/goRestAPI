package database

import (
	"sync"
	"crudAPI/internal/model"
)

type Usersdbstr struct{
	Usersdb map[int]model.User
	NextID int
	Mu sync.RWMutex
}

func (u *Usersdbstr) Dbinit() {

	u.Usersdb = make(map[int]model.User)
	u.NextID=1

	u.Usersdb[u.NextID] = model.User{
		ID:       u.NextID,
		Username: "rupesh",
		Age:      25,
		Hobbies:  []string{"exercise", "movies"},
	}
	u.NextID++
	
	u.Usersdb[u.NextID] = model.User{
		ID:       u.NextID,
		Username: "amit",
		Age:      20,
		Hobbies:  []string{"cricket"},
	}
	u.NextID++
}