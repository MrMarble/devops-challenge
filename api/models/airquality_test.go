package models

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/gocarina/gocsv"
)

var airQualityCsv string = `TimeInstant,id_entity,so2,no2,co,o3,pm10,pm2_5
2016-10-01 00:00:00.004,aq_salvia,6.80117094260474,48.398337879833704,0.657363926741451,48.49706558445371,20.1015302324903,9.137353903174679`

func setUp() (*AirQuality, error) {
	t, err := time.Parse(time.RFC3339, "2016-10-01T00:00:00.004Z")
	if err != nil {
		return nil, err
	}
	var airQuality AirQuality = AirQuality{
		Timestamp: &DateTime{t},
		EntityID:  "aq_salvia",
		SO2:       6.80117094260474,
		NO2:       48.398337879833704,
		CO:        0.657363926741451,
		O3:        48.49706558445371,
		PM10:      20.1015302324903,
		PM25:      9.137353903174679,
	}
	return &airQuality, nil
}

func TestUnmarshalCSV(t *testing.T) {
	got := []*AirQuality{}
	want, _ := setUp()

	err := gocsv.UnmarshalString(airQualityCsv, &got)
	if err != nil {
		t.Fatalf("Got error when none expected: %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("Got %v, wanted %v", len(got), 1)
	}

	if !reflect.DeepEqual(got[0], want) {
		t.Errorf("Got %v, wanted %v", got[0], want)
	}
}

func TestMarshalJSON(t *testing.T) {
	airStruct, _ := setUp()
	got, err := json.Marshal(airStruct)
	want := `{"timestamp":"2016-10-01 00:00:00.004","id_entity":"aq_salvia","so2":6.80117094260474,"no2":48.398337879833704,"co":0.657363926741451,"o3":48.49706558445371,"pm10":20.1015302324903,"pm2_5":9.137353903174679}`
	if err != nil {
		t.Fatalf("Got error when none expected: %v", err)
	}
	if string(got) != want {
		t.Errorf("Got %v, wanted %v", string(got), want)
	}
}
