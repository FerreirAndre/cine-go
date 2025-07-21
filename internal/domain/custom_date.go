package domain

import (
	"strings"
	"time"
)

type CustomDate struct {
	time.Time
}

const dateLayout = "02/01/2006"

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + cd.Time.Format(dateLayout) + `"`), nil
}
