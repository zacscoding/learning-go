package main

import (
	"fmt"
	"strconv"
)

type MemberA struct {
	Id   uint
	Name string
	Age  int
}

func (m MemberA) String() string {
	return fmt.Sprintf("Id : %d, Name : %s, Ages : %d", m.Id, m.Name, m.Age)
}

type MemberB struct {
	Id    uint
	Name  string
	Hobby string
}

func (m MemberB) String() string {
	return fmt.Sprintf("Id : %d, Name : %s, Hobby : %s", m.Id, m.Name, m.Hobby)
}

const (
	emptyId = uint(0)
)

func main() {
	memberAList := createMemberA(1)
	memberBList := createMemberB(2)

	members, memberANext, memberBNext := aggregateMembers(memberAList, memberBList, 4)
	fmt.Println("result :", len(members))
	for _, member := range members {
		if v, ok := member.(MemberA); ok {
			fmt.Println("MemberA ==> " + v.String())
		} else if v, ok := member.(MemberB); ok {
			fmt.Println("MemberB ==> " + v.String())
		}
	}
	fmt.Println("memberA next :", memberANext)
	fmt.Println("memberB next :", memberBNext)
}

func createMemberA(ids ...uint) []MemberA {
	var members []MemberA
	for _, id := range ids {
		members = append(members, MemberA{Id: id, Name: "MemberA" + strconv.Itoa(int(id)), Age: int(id)})
	}
	return members
}

func createMemberB(ids ...uint) []MemberB {
	var members []MemberB
	for _, id := range ids {
		members = append(members, MemberB{Id: id, Name: "MemberB" + strconv.Itoa(int(id)), Hobby: "Hobby" + strconv.Itoa(int(id))})
	}
	return members
}

//func aggregateMembers2(memberAList []MemberA, memberBList []MemberB, size int) ([]interface{}, uint, uint) {
//	sort.Slice(memberAList, func(i, j int) bool {
//		return memberAList[i].Id > memberAList[j].Id
//	})
//	sort.Slice(memberBList, func(i, j int) bool {
//		return memberBList[i].Id > memberBList[j].Id
//	})
//
//	memberALen := len(memberAList)
//	if memberALen > size {
//		memberALen = size
//	}
//	memberBLen := len(memberBList)
//	if memberBList > size {
//		memberBLen = size
//	}
//
//	var results []interface{}
//	if memberALen > 0 {
//		results = append(results, memberAList)
//	}
//	if memberBLen > 0 {
//		results = append(results, memberBList)
//	}
//
//	sort.Slice(results, func(i, j int) bool {
//		return false
//	})
//}

func aggregateMembers(memberAList []MemberA, memberBList []MemberB, size int) ([]interface{}, uint, uint) {
	// check less than size?
	var results []interface{}
	memberALastId := emptyId
	memberBLastId := emptyId

	if len(memberAList) > 0 {
		memberALastId = memberAList[0].Id + 1
	}

	if len(memberBList) > 0 {
		memberBLastId = memberBList[0].Id + 1
	}

	memberAIdx := 0
	memberBIdx := 0

	for i := 0; i < size; i++ {
		// not exists both
		if memberAIdx >= len(memberAList) && memberBIdx >= len(memberBList) {
			memberALastId = emptyId
			memberBLastId = emptyId
			break
		}

		if memberAIdx >= len(memberAList) {
			results = append(results, memberBList[memberBIdx])
			memberBLastId = memberBList[memberBIdx].Id
			memberALastId = emptyId
			memberBIdx++
			continue
		}

		if memberBIdx >= len(memberBList) {
			results = append(results, memberAList[memberAIdx])
			memberALastId = memberAList[memberAIdx].Id
			memberBLastId = emptyId
			memberAIdx++
			continue
		}

		if memberAList[memberAIdx].Id >= memberBList[memberBIdx].Id {
			results = append(results, memberAList[memberAIdx])
			memberALastId = memberAList[memberAIdx].Id
			memberAIdx++
		} else {
			results = append(results, memberBList[memberBIdx])
			memberBLastId = memberBList[memberBIdx].Id
			memberBIdx++
		}
	}

	return results, memberALastId, memberBLastId
}
