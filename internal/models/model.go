package models

type TableStruct struct {
	PhiloCount int
	TimeDie    int
	TimeEat    int
	TimeSleep  int
	MustEat    int
	FullPhilo  int
	StartTime  int64
	Philos     []*PhiloStruct
	Forks      []*bool
}

type PhiloStruct struct {
	number    int
	state     int
	eatCount  int
	leftFork  *bool
	rightFork *bool
	table     *TableStruct
}
