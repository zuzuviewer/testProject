package mongo

import (
	//"ecologyServer/handler"
	"gopkg.in/mgo.v2"
)

type oneAlgae struct {
	algaeName    int32
	algaeCount   int32
	algaeDensity float64
}

type Result struct {
	DeviceId              string     `json:"device_id"`
	UserId                int32      `json:"user_id"`
	ExperimentName        string     `json:"experiment_name"`
	Algaes                []oneAlgae `json:"algaes,omitempty"`
	TotalAlgaeDensity     float64    `json:"total_algae_density"`
	TotalAlgaeCount       int32      `json:"total_algae_count"`
	AdvantageAlgaeName    int32      `json:"advantage_algae_name"`
	AdvantageAlgaeDensity float64    `json:"advantage_algae_density"`
	AdvantageAlgaePercent float32    `json:"advantage_algae_percent"`
	ScannerSampleVolume   float32    `json:"scanner_sample_volume"` //取样容量
	SampleVolume          float32    `json:"sample_volume"`         //玻片容量
	TotalVolume           float32    `json:"total_volume"`          //总容量
	DilutionMultiple      int32      `json:"dilution_multiple"`
	ViewCount             int32      `json:"view_count"`
	SamplePlace           string     `json:"sample_place,omitempty"`
	SampleDate            string     `json:"sample_date"`
	ExperimentDateTime    string     `json:"experiment_date_time"`
	OneViewArea           float32    `json:"one_view_area"`
	SlideArea             float32    `json:"slide_area"`
}

var session = &mgo.Session{}

func Init() error {
	initSession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		return err
	}
	session = initSession
	session.SetMode(mgo.Strong, true)
	return err
}

func SaveExperimentResult(result Result) error {
	db := session.DB("ecology").C("scanner_result")
	err := db.Insert(&result)
	return err
}
