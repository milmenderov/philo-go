package philo

ipmort (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/table"
)

type PhiloStruct struct {
	Number    int
	State     int
	EatCount  int
	LeftFork  *bool
	RightFork *bool
	Table     *TableStruct
}
