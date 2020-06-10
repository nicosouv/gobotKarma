[![Go Report Card](https://goreportcard.com/badge/github.com/PoulpePoulpePoulpe/gobotKarma?style=flat-square)](https://goreportcard.com/report/github.com/PoulpePoulpePoulpe/gobotKarma)

# GobotKarma

Free karma bot for Slack.

Handle live reloading containers for development. It also contains a production ready container.

## Before anything 
You need an `.env` file in `internal` folder.
Here is an example:
```
SLACK_TOKEN=xoxb-12345678912346468774135453
SLACK_CHANNEL=general
WEBHOOK=https://hooks.slack.com/services/WHATEVER/TOKEN/YOU/SHOULD/HAVE/HERE
```

## Docker install for development
You can just do:

```bash
whatever$ cd deployments
deployments$ docker-compose up --build gobot-development
```
That's all.

## Docker install for production
Soon.

## TIL
You can use tool like `ngrok` to test your localhost app on the web! Check it out: https://ngrok.com/
