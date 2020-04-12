package main

const (
	ApplicationName = "helloworld"
	BindAddress     = "127.0.0.1:7933"
	DomainName      = "eg-cadence-jobportal"
	TaskListName    = DomainName + "-" + ApplicationName + "-tasks"
	ClientName      = DomainName + "-" + ApplicationName
	ServiceName     = "cadence-frontend"
)
