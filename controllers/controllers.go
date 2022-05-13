package controllers

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/teamanict/marks/excel"
)

func File(c *fiber.Ctx) error {
	var freader io.Reader
	freader.Read(c.Context().FormValue("sheet"))
	marks := excel.Parsesheet(freader)
	return c.JSON(marks)
}
