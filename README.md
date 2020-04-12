# Exercise: Jobs Portal Workflow using Cadence

This repository contains an exercise I'm doing to learn how code-based workflow engines work.

# Pre-Requisites

To run stuff in here, confirm you have the following installed:

| Software | Command to check | Install |
| --- | --- | --- |
| Docker | `docker -v` | [Install Docker](https://docs.docker.com/get-docker/) |
| Docker Compose | `docker-compose -v` | [Install Docker Compose](https://docs.docker.com/compose/install/) |
| Go | `go version` | [Install Go](https://golang.org/doc/install) |
| Make | `make -v` | [Install Make (for Windows)](http://gnuwin32.sourceforge.net/packages/make.htm) |

# Setting Up

1. Run `make deploy_cadence` to setup a local instance of Cadence
2. Run `make register_domain` to setup a testing domain named `"eg-cadence-jobportal"`
3. Run `make deps` to pull in all dependencies

# Workflows

## `helloworld`

> Note: this was done mainly for practice to get a POC workflow up

A POC workflow that interacts with the Cadence workflow service.

To run this service in worker mode (runs as a long-running service that handles workflow requests), run `go ./workflows/helloworld -m worker`

To run this service in trigger mode (triggers the workflow), run `go ./workflows/helloworld -m trigger`

## `apply-job`

A job application workflow for a typical jobseeker:

1. Jobseeker is on jobs page
2. Jobseeker clicks on **Apply for Job**
3. Jobseeker enters in contact information and clicks **Next**
4. Jobseeker enters in professional credentials and clicks **Next**
5. Jobseeker reviews the job application and clicks **Submit**
6. Jobseeker sees a page indicating job has successfully been applied for

## `accept-applicant`

A 

# Services

