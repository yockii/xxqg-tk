package model

type QuestionBank struct {
	Id       string `json:"id,omitempty" xorm:"pk varchar(50)"`
	Question string `json:"question,omitempty" xorm:"varchar(2000)"`
	Options  string `json:"options,omitempty" xorm:"varchar(1000)"`
	Answer   string `json:"answer,omitempty" xorm:"varchar(500)"`
}

func init() {
	SyncModel = append(SyncModel, QuestionBank{})
}
