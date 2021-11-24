# GoLang Lambda

![Deploy](https://github.com/WsCandy/go-lambda/actions/workflows/deploy.yml/badge.svg?branch=master)

A starting point for an AWS Lambda function in Go with Terraform and GitHub Actions.

## Prerequisites

- Uses Terraform Cloud, requires `TF_API_TOKEN` to be set either in your organisation or individual repository for `apply` and `plan` to run.
- Requires you to set the `backend` block in `main.tf`. For Terraform Cloud you will need to set an organisation and workspace. See below:

```
backend "remote" {
    organization = "organisation_name"

    workspaces {
      name = "workspace_name"
    }
}
```

## Workflow
The template is set up to work with a PR team workflow. The `master` branch should be protected and only **merge commits** should be allowed.

On PR the code will be build, run automated testing and present a Terraform plan for review. The plan will be automatically posted to the PR for review. Once the code has been reviewed the PR can be merged if it is up-to-date with `master`. Once merged the plan will be applied and the deployment will take place.