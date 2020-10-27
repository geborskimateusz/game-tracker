package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlertFunc func(duration time.Duration, amount int)

func (b BlindAlertFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	b(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprint(os.Stdout, "Blimd is now %d\n", amount)
	})
}
