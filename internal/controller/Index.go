package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	logger "github.com/sirupsen/logrus"

	"xxqg-tk/internal/model"
	"xxqg-tk/pkg/database"
	"xxqg-tk/pkg/server"
	"xxqg-tk/pkg/util"
)

type Bank struct {
	Question string `json:"q,omitempty" form:"q"`
	Answer   string `json:"a,omitempty" form:"a"`
}

func InitRouter() {
	qa := server.Group("/api/v1/bank")
	qa.Post("/query", func(ctx *fiber.Ctx) error {
		bank := new(Bank)
		if err := ctx.BodyParser(bank); err != nil {
			return fiber.ErrBadRequest
		}
		questionWithOptions := strings.Split(bank.Question, "|")

		qb := &model.QuestionBank{Question: questionWithOptions[0], Options: strings.Join(questionWithOptions[1:], "|")}
		if exist, err := database.DB.Get(qb); err != nil {
			logger.Error(err)
			return ctx.SendString("")
		} else if exist {
			return ctx.SendString(qb.Answer)
		}
		return ctx.SendString("")
	})

	qa.Post("/queryLike", func(ctx *fiber.Ctx) error {
		bank := new(Bank)
		if err := ctx.BodyParser(bank); err != nil {
			return fiber.ErrBadRequest
		}
		questionWithOptions := strings.Split(bank.Question, "|")

		qb := &model.QuestionBank{Question: questionWithOptions[0], Options: strings.Join(questionWithOptions[1:], "|")}
		if exist, err := database.DB.Where("question like ?", questionWithOptions[0]).Get(qb); err != nil {
			logger.Error(err)
			return ctx.SendString("")
		} else if exist {
			return ctx.SendString(qb.Answer)
		}
		return ctx.SendString("")
	})

	qa.Post("/add", func(ctx *fiber.Ctx) error {
		bank := new(Bank)
		if err := ctx.BodyParser(bank); err != nil {
			logger.Error(err)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
		questionWithOptions := strings.Split(bank.Question, "|")
		if c, err := database.DB.Count(&model.QuestionBank{Question: questionWithOptions[0], Options: strings.Join(questionWithOptions[1:], "|")}); err != nil {
			logger.Error(err)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else if c > 0 {
			database.DB.Update(&model.QuestionBank{Answer: bank.Answer}, &model.QuestionBank{Question: questionWithOptions[0], Options: strings.Join(questionWithOptions[1:], "|")})
			return ctx.SendString("ok")
		}
		if _, err := database.DB.Insert(&model.QuestionBank{
			Id:       util.GenerateDatabaseID(),
			Question: questionWithOptions[0],
			Options:  strings.Join(questionWithOptions[1:], "|"),
			Answer:   bank.Answer,
		}); err != nil {
			logger.Error(err)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
		return ctx.SendString("ok")
	})
}
