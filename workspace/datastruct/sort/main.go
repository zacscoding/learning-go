package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Member struct {
	Id   uint
	Name string
}

func (m Member) String() string {
	return fmt.Sprintf("Member{Id:%d, Name:%s}", m.Id, m.Name)
}

func main() {
	members := []Member{
		createMember(5),
		createMember(10),
		createMember(6),
		createMember(1),
	}

	sort.Slice(members, func(i, j int) bool {
		//// id asc
		//return members[i].Id < members[j].Id
		// id desc
		return members[i].Id > members[j].Id
	})

	for i, member := range members {
		fmt.Printf("[%d] %s\n", i, member.String())
	}
}

func createMember(id uint) Member {
	return Member{Id: id, Name: "Member" + strconv.Itoa(int(id))}
}
