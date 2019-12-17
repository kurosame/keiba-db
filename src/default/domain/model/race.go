package model

// RaceResult is race result
type RaceResult struct {
	Rank         uint    // 着順
	FrameNumber  uint    // 枠番
	HorseNumber  uint    // 馬番
	HorseName    string  // 馬名
	Age          string  // 性齢
	Weight       uint    // 斤量
	JockeyName   string  // 騎手
	Time         string  // タイム
	Gap          string  // 着差
	TimeIndex    string  // ﾀｲﾑ指数
	Passing      string  // 通過
	Rise         float32 // 上り
	Odds         float32 // 単勝
	Popular      uint    // 人気
	HorseWeight  string  // 馬体重
	TrainingTime string  // 調教ﾀｲﾑ
	Comment      string  // 厩舎ｺﾒﾝﾄ
	Remarks      string  // 備考
	TrainerName  string  // 調教師
	OwnerName    string  // 馬主
	Prize        float32 // 賞金(万円)
}

// Race is race details
type Race struct {
	RaceResults []*RaceResult
}
