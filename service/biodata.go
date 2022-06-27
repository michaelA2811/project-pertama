package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {
	var i int
	checkEmployee(i)
}

func checkEmployee(i int) {
	var employee = []struct {
		person Biodata
	}{
		{person: Biodata{name: "Budi", address: "Jl.Semangka", job: "Software Developer", reason: "Butuh uang"}},
		{person: Biodata{name: "Truno", address: "Jl.Pepaya", job: "Backend Developer", reason: "Butuh uang"}},
		{person: Biodata{name: "Aldo", address: "Jl.Pisang", job: "Human Resource", reason: "Butuh uang"}},
		{person: Biodata{name: "Sodikin", address: "Jl.Kejujuran", job: "Frontend Developer", reason: "Butuh uang"}},
	}
	tempArgs := os.Args
	if len(tempArgs) == 2 {
		i, err := strconv.Atoi(tempArgs[1])
		if err != nil {
			fmt.Println("angkanya error")
			return
		}
		fmt.Printf("%+v\n", employee[i])
	}
	// fmt.Println(os.Args)
	// fmt.Println(employee[len(os.Args)-1])
}
