package main

import "fmt"

type A struct {
	UID int64
}

func (a *A) GetUID() int64 {
	return a.UID
}

type B struct {
	UserID int64
}

func (b *B) GetUserID() int64 {
	return b.UserID
}

func GetUserID(something interface{}) int64 {
	if st, _ := something.(interface{ GetUID() int64 }); st != nil {
		return st.GetUID()
	}

	if st, _ := something.(interface{ GetUserID() int64 }); st != nil {
		return st.GetUserID()
	}

	return int64(0)
}

func main() {
	exampleA := &A{UID: 123}
	fmt.Printf("exampleA.GetUID()=%v\n", exampleA.GetUID())

	exampleB := &B{UserID: 234}
	fmt.Printf("exampleB.GetUserID()=%v\n", exampleB.GetUserID())

	fmt.Printf("exampleA.GetUID()=%v with interface assertion\n", GetUserID(exampleA))
	fmt.Printf("exampleB.GetUserID()=%v with interface assertion\n", GetUserID(exampleB))
}
