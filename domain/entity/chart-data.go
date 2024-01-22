package entity

type CandlestickElement struct {
	From  int64 //시작 타임스탬프
	Start int32 //시작 값
	End   int32 //종료 값
	Min   int32 //최댓값
	Max   int32 //최솟값
}

type AverageElement struct {
	From int32 `json:"from"`
	Val  int32 `json:"val"`
}
