package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// AirQuality represent an air quality entry.
type AirQuality struct {
	Timestamp *DateTime `gorm:"PRIMARY_KEY;NOT NULL;type:timestamp" json:"timestamp" csv:"TimeInstant"`
	EntityID  string    `gorm:"PRIMARY_KEY;NOT NULL" json:"id_entity" csv:"id_entity"`
	SO2       float64   `gorm:"NOT NULL" json:"so2" csv:"so2"`
	NO2       float64   `gorm:"NOT NULL" json:"no2" csv:"no2"`
	CO        float64   `gorm:"NOT NULL" json:"co" csv:"co"`
	O3        float64   `gorm:"NOT NULL" json:"o3" csv:"o3"`
	PM10      float64   `gorm:"NOT NULL" json:"pm10" csv:"pm10"`
	PM25      float64   `gorm:"NOT NULL" json:"pm2_5" csv:"pm2_5"`
}

// DateTime custom date type
type DateTime struct {
	time.Time
}

// UnmarshalCSV Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05.000", csv)
	return err
}

// MarshalJSON converts the internal date as json string
func (date *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", date.Time.Format("2006-01-02 15:04:05.000"))), nil
}

// Value return csv value, implement driver.Valuer interface
func (date *DateTime) Value() (driver.Value, error) {
	return date.Time, nil
}

// Scan scan value into csv, implements sql.Scanner interface
func (date *DateTime) Scan(b interface{}) (err error) {
	switch x := b.(type) {
	case time.Time:
		date.Time = x
	default:
		err = fmt.Errorf("unsupported scan type %T, %v", b, b)
	}
	return err
}
