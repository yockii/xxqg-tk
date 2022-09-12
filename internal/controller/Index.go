package controller

import (
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	logger "github.com/sirupsen/logrus"

	"xxqg-tk/internal/model"
	"xxqg-tk/pkg/database"
	"xxqg-tk/pkg/domain"
	"xxqg-tk/pkg/server"
	"xxqg-tk/pkg/util"
)

type Bank struct {
	Question    string `json:"q,omitempty" form:"q"`
	Options     string `json:"o,omitempty" form:"o"`
	Answer      string `json:"a" form:"a"`
	WrongAnswer string `json:"wa" form:"wa"`
}

func (qb *Bank) SortOptions() {
	options := strings.Split(qb.Options, "|")
	sort.Strings(options)
	qb.Options = strings.Join(options, "|")
}

func InitRouter() {
	qa := server.Group("/api/v1/bank")
	qa.Post("/query", func(ctx *fiber.Ctx) error {
		bank := new(Bank)
		if err := ctx.BodyParser(bank); err != nil {
			return fiber.ErrBadRequest
		}
		//questionWithOptions := strings.Split(bank.Question, "|")
		bank.SortOptions()
		qb := &model.QuestionBank{Options: bank.Options}
		//qb := &model.QuestionBank{Options: strings.Join(questionWithOptions[1:], "|")}
		if exist, err := database.DB.Where("question like ?", bank.Question+"%").Get(qb); err != nil {
			//if exist, err := database.DB.Where("question like ?", questionWithOptions[0]+"%").Get(qb); err != nil {
			logger.Error(err)
			return ctx.JSON(&domain.CommonResponse{})
		} else if exist {
			return ctx.JSON(&domain.CommonResponse{
				Data: &Bank{
					Answer:      qb.Answer,
					WrongAnswer: qb.WrongAnswer,
				},
			})
		}
		database.DB.Insert(&model.QuestionBank{
			Question: bank.Question,
			//Question:    questionWithOptions[0],
			Options: bank.Options,
			//WrongAnswer: bank.WrongAnswer,
		})
		return ctx.JSON(&domain.CommonResponse{})
	})

	qa.Post("/queryLike", func(ctx *fiber.Ctx) error {
		bank := new(Bank)
		if err := ctx.BodyParser(bank); err != nil {
			return fiber.ErrBadRequest
		}
		//questionWithOptions := strings.Split(bank.Question, "|")
		bank.SortOptions()

		qb := &model.QuestionBank{Options: bank.Options}
		//qb := &model.QuestionBank{Question: questionWithOptions[0], Options: strings.Join(questionWithOptions[1:], "|")}
		if exist, err := database.DB.Where("question like ?", bank.Question).Get(qb); err != nil {
			logger.Error(err)
			return ctx.JSON(&domain.CommonResponse{})
		} else if exist {
			return ctx.JSON(&domain.CommonResponse{
				Data: &Bank{
					Answer:      qb.Answer,
					WrongAnswer: qb.WrongAnswer,
				},
			})
		}

		return ctx.JSON(&domain.CommonResponse{})
	})

	qa.Post("/add", func(ctx *fiber.Ctx) error {
		bank := new(Bank)
		if err := ctx.BodyParser(bank); err != nil {
			logger.Error(err)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
		bank.SortOptions()
		dbBank := &model.QuestionBank{
			Question: bank.Question,
			Options:  bank.Options,
		}
		if exists, err := database.DB.Get(dbBank); err != nil {
			logger.Error(err)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else if exists {
			if strings.Contains(dbBank.WrongAnswer, bank.WrongAnswer) {
				bank.WrongAnswer = ""
			} else {
				bank.WrongAnswer = dbBank.WrongAnswer + "|" + bank.WrongAnswer
			}
			database.DB.ID(dbBank.Id).Update(
				&model.QuestionBank{
					Answer:      bank.Answer,
					WrongAnswer: bank.WrongAnswer,
				})
			return ctx.JSON(&domain.CommonResponse{})
		}
		if _, err := database.DB.Insert(&model.QuestionBank{
			Id:          util.GenerateDatabaseID(),
			Question:    bank.Question,
			Options:     bank.Question,
			Answer:      bank.Answer,
			WrongAnswer: bank.WrongAnswer,
		}); err != nil {
			logger.Error(err)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
		return ctx.JSON(&domain.CommonResponse{})
	})
}
