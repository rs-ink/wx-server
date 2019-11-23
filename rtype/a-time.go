package rtype

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

const rsDateFormat string = "2006-01-02"

const rsDateTimeFormat string = "2006-01-02 15:04:05"

type RDate time.Time

func (l RDate) MarshalJSON() ([]byte, error) {
	format := time.Time(l).Format(rsDateFormat)
	return json.Marshal(format)
}

func (l *RDate) UnmarshalJSON(b []byte) error {
	var aa string
	err := json.Unmarshal(b, &aa)
	if err == nil {
		dd, err := time.Parse(rsDateFormat, aa)
		if err == nil {
			*l = RDate(dd)
		}
	}
	return err
}

type RDateTime time.Time

func (t *RDateTime) UnmarshalJSON(b []byte) error {
	var aa string
	err := json.Unmarshal(b, &aa)
	if err == nil {
		dd, err := time.Parse(rsDateTimeFormat, aa)
		if err == nil {
			*t = RDateTime(dd)
		}
	}
	return err
}

func (t RDateTime) MarshalJSON() ([]byte, error) {
	format := time.Time(t).Format(rsDateTimeFormat)
	return json.Marshal(format)
}

type RDateTimeSecond time.Time

func (rt *RDateTimeSecond) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var sec int64
	err := d.DecodeElement(&sec, &start)
	if err != nil {
		return err
	}
	*rt = RDateTimeSecond(time.Unix(sec, 0))
	return nil
}

func (rt *RDateTimeSecond) UnmarshalJSON(data []byte) error {
	var sec int64
	err := json.Unmarshal(data, &sec)
	if err != nil {
		return err
	}
	*rt = RDateTimeSecond(time.Unix(sec, 0))

	return nil
}

func (rt RDateTimeSecond) MarshalJSON() ([]byte, error) {
	t := time.Time(rt)
	return json.Marshal(t.Unix())
}
