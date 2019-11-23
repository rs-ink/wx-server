package condition

import (
	"log"
	"testing"
)

func TestNewCondition(t *testing.T) {
	con := NewCondition()
	con.Add(OpeNumericLessThanEquals, KeySecureTransport, "https")

	log.Println(con.String())
}
