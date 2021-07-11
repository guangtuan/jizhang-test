package main

import (
	"github.com/cucumber/godog"
	"jizhang-test-go/Implemention"
)

func InitializeScenario(ctx *godog.ScenarioContext) {

	api := &Implemention.Api{}

	ctx.BeforeScenario(api.ResetResponse)

	ctx.Step(`^login with email "([^"]*)" and password "([^"]*)"$`, api.LoginWithEmailAndPassword)
	ctx.Step(`^the response code should be (\d+)$`, api.TheResponseCodeShouldBe)
	ctx.Step(`^the response body should be a json string with a (\d+)-bit token`, api.TheResponseShouldHasTokenWithLength)

}