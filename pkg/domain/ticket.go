package domain

import (
	"fmt"
	"math/rand"
)

type Ticket struct {
	Id          string
	Issue       string
	Description string
}

func (t *Ticket) GenID() {
	t.Id = fmt.Sprint(rand.Int())
}

func (t *Ticket) IsValid() bool {
	return rand.Intn(11) < 10
}
