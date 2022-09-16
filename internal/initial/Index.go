package initial

import (
	"regexp"
	"strings"

	logger "github.com/sirupsen/logrus"

	"xxqg-tk/internal/model"
	"xxqg-tk/pkg/database"
)

func InitData() {
	//changeQuestionForm()
}

func changeQuestionForm() {
	var blankReg = regexp.MustCompile("[ \\s]+")
	var all []*model.QuestionBank
	if err := database.DB.Find(&all); err != nil {
		logger.Fatal(err)
	}
	for _, qb := range all {
		question := strings.TrimSpace(qb.Question)
		question = blankReg.ReplaceAllString(question, "____")

		if c, err := database.DB.Count(&model.QuestionBank{
			Question: question,
			Options:  qb.Options,
		}); err != nil {
			logger.Error(err)
			continue
		} else if c > 0 {
			database.DB.ID(qb.Id).Delete(&model.QuestionBank{})
			continue
		} else {
			database.DB.ID(qb.Id).Update(&model.QuestionBank{Question: question})
		}
	}
}
