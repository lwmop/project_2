package models

type TextInput struct {
	Text string `json:"text"`
}

type WordFrequencyStruct struct {
	Key   string `json:"word"`
	Count int    `json:"count"`
}
type WordFrequencyStructData struct {
	Dat []WordFrequencyStruct `json:"dat"`
}
