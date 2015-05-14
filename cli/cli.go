//Copyright 2015 Bestupefy.Inc

  // Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

  //     http://www.apache.org/licenses/LICENSE-2.0

   //Unless required by applicable law or agreed to in writing, software
   //distributed under the License is distributed on an "AS IS" BASIS,
   //WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   //See the License for the specific language governing permissions and
   //limitations under the License.
   
package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/bestupefy/canister"
)

var (
	canisterHost string
	logger       = logrus.New()
)

func main() {
	cfg, err := loadConfig(nil)
	if err != nil {
		if err != ErrConfigDoesNotExist {
			logger.Fatal(err)
		}
	}
	if cfg != nil {
		sUrl := os.Getenv("CANISTER_URL")
		if sUrl == "" {
			cfg.Url = sUrl
		}
	}
	app := cli.NewApp()
	app.Name = "canister"
	app.Usage = "manage your canister"
	app.Version = canister.VERSION
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "allow-insecure",
			Usage: "allow insecure certificates if using TLS",
		},
	}
	app.Commands = []cli.Command{
		loginCommand,
		changePasswordCommand,
		accountsCommand,
		addAccountCommand,
		deleteAccountCommand,
		containersCommand,
		containerInspectCommand,
		runCommand,
		stopCommand,
		restartCommand,
		scaleCommand,
		logsCommand,
		destroyCommand,
		engineListCommand,
		engineAddCommand,
		engineRemoveCommand,
		engineInspectCommand,
		serviceKeysListCommand,
		serviceKeyCreateCommand,
		serviceKeyRemoveCommand,
		extensionsCommand,
		addExtensionCommand,
		removeExtensionCommand,
		webhookKeysListCommand,
		webhookKeyCreateCommand,
		webhookKeyRemoveCommand,
		infoCommand,
		eventsCommand,
	}
	app.Run(os.Args)
}
