package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/loopcontext/go-graphql-orm/templates"

	"gopkg.in/yaml.v2"

	"github.com/loopcontext/go-graphql-orm/model"
	"github.com/urfave/cli"
)

var initCmd = cli.Command{
	Name:  "init",
	Usage: "initialize new project",
	Action: func(ctx *cli.Context) error {
		p := ctx.Args().First()
		if p == "" {
			p = "."
		}

		fmt.Printf("Initializing project in %s ...\n", p)

		if !fileExists(path.Join(p, "graphql-orm.yml")) {
			if err := createConfigFile(p); err != nil {
				return cli.NewExitError(err, 1)
			}
		}

		if !fileExists(path.Join(p, "model.graphql")) {
			if err := createDummyModelFile(p); err != nil {
				return cli.NewExitError(err, 1)
			}
		}

		if err := createMainFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}
		if err := createLambdaMainFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if !fileExists(path.Join(p, "src/resolver.go")) {
			if err := createResolverFile(p); err != nil {
				return cli.NewExitError(err, 1)
			}
		}

		if err := createMakeFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := createDockerFiles(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := runGenerate(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	},
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return true
	}
	return false
}

func createConfigFile(p string) error {
	packagePath := ""
	modFilename := "go.mod"

	_, err := os.Stat(modFilename)
	if os.IsNotExist(err) {
		return fmt.Errorf("Go modules required (no go.mod file found). Use `go mod init MODULE_NAME` to initialize go modules")
	}

	data, err := ioutil.ReadFile(modFilename)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(bytes.NewReader(data))
	line, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	packagePath = strings.ReplaceAll(string(line), "module ", "")

	c := model.Config{Package: packagePath}

	content, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(p, "graphql-orm.yml"), content, 0o644)
	return err
}

func createMainFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	return templates.WriteTemplate(templates.Main, path.Join(p, "main.go"), templates.TemplateData{Config: &c})
}

func createLambdaMainFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	ensureDir(path.Join(p, "lambda"))
	return templates.WriteTemplate(templates.Lambda, path.Join(p, "lambda/main.go"), templates.TemplateData{Config: &c})
}

func createDummyModelFile(p string) error {
	data := templates.TemplateData{Model: nil, Config: nil}
	return templates.WriteTemplate(templates.DummyModel, path.Join(p, "model.graphql"), data)
}

func createMakeFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	data := templates.TemplateData{Model: nil, Config: &c}
	return templates.WriteTemplate(templates.Makefile, path.Join(p, "makefile"), data)
}

func createDockerFiles(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	data := templates.TemplateData{Model: nil, Config: &c}

	ensureDir(path.Join(p, "docker"))
	ensureDir(path.Join(p, "scripts"))
	err = templates.WriteTemplate(templates.GitIgnore, path.Join(p, ".gitignore"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.AirConf, path.Join(p, ".air.conf"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.DotenvExample, path.Join(p, ".env.example"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.DotenvExample, path.Join(p, ".env.dev"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.RunDevSh, path.Join(p, "scripts/run-dev.sh"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.RunSh, path.Join(p, "scripts/run.sh"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.DockerfileDev, path.Join(p, "docker/dev.dockerfile"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.DockerfileProd, path.Join(p, "docker/prod.dockerfile"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.Dockerfile, path.Join(p, "Dockerfile"), data)
	if err != nil {
		return err
	}
	err = templates.WriteTemplate(templates.DockerComposeYml, path.Join(p, "docker-compose.yml"), data)
	return err
}

func createResolverFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	data := templates.TemplateData{Model: nil, Config: &c}
	ensureDir(path.Join(p, "src"))
	return templates.WriteTemplate(templates.ResolverSrc, path.Join(p, "src/resolver.go"), data)
}

func runGenerate(p string) error {
	return generate("model*.graphql", p)
}
