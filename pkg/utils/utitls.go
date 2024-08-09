package utils

import "time"

type Time time.Time

func (m *Time) UnmarshalCSV(csvValue string) error {
	t, err := time.Parse("01/02/2006 15:04", csvValue)
	if err != nil {
		return err
	}

	*m = Time(t)
	
	return nil
}

func (m Time) MarshalCSV() (string, error) {
	return time.Time(m).Format("01/02/2006 15:04"), nil
}

type Date interface {
	nsec() int32
	sec() int64
	unixSec() int64
	addSec(d int64)
	setLoc(loc *time.Location)
	stripMono()
	setMono(m int64)
	mono() int64
	UnmarshalBinary(data []byte) error
	GobDecode(data []byte) error
	UnmarshalJSON(data []byte) error
	UnmarshalText(data []byte) error
}
