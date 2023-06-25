package stRuct

type TableStruct struct {
	PhiloCount int
	TimeDie    int
	TimeEat    int
	TimeSleep  int
	MustEat    int
	FullPhilo  int
	StartTime  int64
	Philos     []PhiloStruct
	Forks      []bool
}

type PhiloStruct struct {
	Number    int
	State     int
	EatCount  int
	LeftFork  *bool
	RightFork *bool
	Table     *TableStruct
}
