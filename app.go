package main

import "github.com/gofiber/fiber/v2"

type App struct {
	*fiber.App
}

func Default() *App {

	return &App{}
}
