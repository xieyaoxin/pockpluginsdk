package service

import (
	"context"
	"encoding/json"
	"fmt"
)

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func buildResponse(code int, response interface{}) string {
	Response := make(map[string]interface{})
	Response["code"] = code
	Response["data"] = response
	marshal, err := json.Marshal(Response)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func buildSuccessResponse(response interface{}) string {
	return buildResponse(200, response)
}

func buildFailedResponse(response interface{}) string {
	return buildResponse(400, response)
}

func ToDo() string {
	return buildResponse(200, nil)
}
