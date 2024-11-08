package main

import "fmt"

// Biểu thức cơ bản với iota

type Weekday int

const (
	Sunday    Weekday = iota //Weekday = 1
	Monday                   //Weekday = 2
	Tuesday                  //Weekday = 3
	Wednesday                //Weekday = 4
	Thursday                 //Weekday = 5
	Friday                   //Weekday = 6
	Saturnday                //Weekday = 7
)

func (day Weekday) String() string {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturnday"}
	if day < Sunday || day > Saturnday {
		return "Unknown"
	}
	return names[day]
}

func (day Weekday) Weekend() bool {
	switch day {
	case Saturnday, Sunday:
		return true
	default:
		return false
	}
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

// Iota reset

const (
	zero = iota
	one
)

// iota sẽ reset về 0
const (
	two = iota
)

// iota sẽ reset về 0 lần nữa => mỗi lần const xuất hiện, iota sẽ đc reset về 0
const three = iota

// skip value

type Timezone int

const (
	EST Timezone = -(5 + iota)
	_            //Blank => bị skip nhưng iota vẫn tăng
	MST
	PST
)

func main() {
	fmt.Printf("EST=%d, MST=%d, PST=%d\n", EST, MST, PST)

	fmt.Println("Which day is it?", Weekday(6))
	fmt.Println("Is it a weekend?", Monday.Weekend())
	fmt.Printf("Sunday=%[1]d (type= %[1]T)\n"+
		"Monday=%[2]d (type= %[2]T)\n"+
		"Tuesday=%[3]d (type= %[3]T)\n"+
		"Wednesday=%[4]d (type= %[4]T)\n"+
		"Friday=%[5]d (type= %[5]T)\n"+
		"Saturnday=%[6]d (type= %[6]T)\n",
		Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturnday)

	ns := transition(StateIdle)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)

	fmt.Println(zero, one, two, three)
}
