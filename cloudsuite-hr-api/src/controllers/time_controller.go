package controllers

import (
	"cloudsuite-hr-api/models"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TimeController struct {
	service services.TimeService
}

func NewTimeController(service services.TimeService) *TimeController {
	return &TimeController{
		service: service,
	}
}

func (c *TimeController) SetupRoutes(app *fiber.App) {
	app.Post("/times", c.CreateTime)
	app.Get("/times", c.GetAllTimes)
	app.Get("/times/date/:date", c.GetTimesByDate)
	app.Get("/times/year/:year", c.GetTimesByYear)
	app.Get("/times/month/:month", c.GetTimesByMonth)
	app.Get("/times/day/:day", c.GetTimesByDay)
}

// @Summary Creates a new time entry
// @Description Creates a time entry for an employee
// @Tags Times
// @Accept  json
// @Produce  json
// @Param time body models.Time true "Time entry details"
// @Success 201 {object} models.Time
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /times [post]
func (c *TimeController) CreateTime(ctx *fiber.Ctx) error {
	var time models.Time
	if err := ctx.BodyParser(&time); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := c.service.CreateTime(time); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create time entry",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(time)
}

// @Summary Retrieves all time entries
// @Description Gets a list of all time entries
// @Tags Times
// @Produce  json
// @Success 200 {array} models.Time
// @Failure 500 {object} map[string]interface{}
// @Router /times [get]
func (c *TimeController) GetAllTimes(ctx *fiber.Ctx) error {
	times, err := c.service.GetAllTimes()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times",
		})
	}
	return ctx.JSON(times)
}

// @Summary Retrieves time entries by date
// @Description Gets a list of time entries for a specific date
// @Tags Times
// @Produce  json
// @Param date path string true "Date in the format YYYY-MM-DD"
// @Success 200 {array} models.Time
// @Failure 500 {object} map[string]interface{}
// @Router /times/date/{date} [get]
func (c *TimeController) GetTimesByDate(ctx *fiber.Ctx) error {
	date := ctx.Params("date")
	times, err := c.service.GetTimesByDate(date)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by date",
		})
	}
	return ctx.JSON(times)
}

// @Summary Retrieves time entries by year
// @Description Gets a list of time entries for a specific year
// @Tags Times
// @Produce  json
// @Param year path int true "Year in the format YYYY"
// @Success 200 {array} models.Time
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /times/year/{year} [get]
func (c *TimeController) GetTimesByYear(ctx *fiber.Ctx) error {
	year := ctx.Params("year")
	parsedYear, err := strconv.Atoi(year)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid year",
		})
	}
	times, err := c.service.GetTimesByYear(parsedYear)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by year",
		})
	}
	return ctx.JSON(times)
}

// @Summary Retrieves time entries by month
// @Description Gets a list of time entries for a specific month
// @Tags Times
// @Produce  json
// @Param month path int true "Month in the format MM"
// @Success 200 {array} models.Time
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /times/month/{month} [get]
func (c *TimeController) GetTimesByMonth(ctx *fiber.Ctx) error {
	month := ctx.Params("month")
	parsedMonth, err := strconv.Atoi(month)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid month",
		})
	}
	times, err := c.service.GetTimesByMonth(parsedMonth)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by month",
		})
	}
	return ctx.JSON(times)
}

// @Summary Retrieves time entries by day
// @Description Gets a list of time entries for a specific day
// @Tags Times
// @Produce  json
// @Param day path int true "Day in the format DD"
// @Success 200 {array} models.Time
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /times/day/{day} [get]
func (c *TimeController) GetTimesByDay(ctx *fiber.Ctx) error {
	day := ctx.Params("day")
	parsedDay, err := strconv.Atoi(day)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid day",
		})
	}
	times, err := c.service.GetTimesByDay(parsedDay)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by day",
		})
	}
	return ctx.JSON(times)
}
