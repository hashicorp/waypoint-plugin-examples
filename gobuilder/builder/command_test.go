package builder

import (
	"testing"
)

func TestLambdaBuilder(t *testing.T) {
	var b Builder
	b.config.OutputName = "lambda"
	b.config.Source = "./"
	var lambda Lambda
	lambda.Amd64 = true
	lambda.Linux = true
	b.config.Lambda = &lambda
	c := BuildCommand(&b)
	output := c.String()
	if string(output) != "GOARCH=amd64 GOOS=linux go build -o lambda ./" {
		t.Fatal("command build sent unexpected output")
	}

}

func TestCommandBuilder(t *testing.T) {
	var b Builder
	b.config.OutputName = "main"
	b.config.Source = "./"
	c := BuildCommand(&b)
	output := c.String()
	if string(output) != " build -o lambda ./" {
		t.Fatal("command build sent unexpected output")
	}
	c.Run()

}
