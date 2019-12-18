package model

import (
	"encoding/json"
	"log"
)

// RaceResult is race result
type RaceResult struct {
	Rank         string `json:"rank"`          // 着順
	FrameNumber  string `json:"frame_number"`  // 枠番
	HorseNumber  string `json:"horse_number"`  // 馬番
	HorseName    string `json:"horse_name"`    // 馬名
	Age          string `json:"age"`           // 性齢
	Weight       string `json:"weight"`        // 斤量
	JockeyName   string `json:"jockey_name"`   // 騎手
	Time         string `json:"time"`          // タイム
	Gap          string `json:"gap"`           // 着差
	TimeIndex    string `json:"time_index"`    // ﾀｲﾑ指数
	Passing      string `json:"passing"`       // 通過
	Rise         string `json:"rise"`          // 上り
	Odds         string `json:"odds"`          // 単勝
	Popular      string `json:"popular"`       // 人気
	HorseWeight  string `json:"horse_weight"`  // 馬体重
	TrainingTime string `json:"training_time"` // 調教ﾀｲﾑ
	Comment      string `json:"comment"`       // 厩舎ｺﾒﾝﾄ
	Remarks      string `json:"remarks"`       // 備考
	TrainerName  string `json:"trainer_name"`  // 調教師
	OwnerName    string `json:"owner_name"`    // 馬主
	Prize        string `json:"prize"`         // 賞金(万円)
}

// Race is race details
type Race struct {
	RaceResults []*RaceResult `json:"race_results"`
}

// ToJSON is Race struct to json
func (r *Race) ToJSON() string {
	buf, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}

	return string(buf)
}
