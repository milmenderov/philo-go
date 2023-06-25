package table

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
