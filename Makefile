CADENCE_DOMAIN=eg-cadence-jobportal

deployments/cadence/docker-compose.yml:
	curl -vv -o ./deployments/cadence/docker-compose.yml 'https://raw.githubusercontent.com/uber/cadence/master/docker/docker-compose.yml'
deploy_cadence:
	docker-compose -f ./deployments/cadence/docker-compose.yml up
deploy_cadence_d:
	docker-compose -f ./deployments/cadence/docker-compose.yml up -d

register_domain:
	docker run --network=host --rm ubercadence/cli:master \
		--do $(CADENCE_DOMAIN) domain register -rd 1
check_domains:
	docker run --network=host --rm ubercadence/cli:master \
		--do $(CADENCE_DOMAIN) domain describe

deps:
	go mod vendor -v
deps_thrift_hack:
	# thanks to https://github.com/uber-go/cadence-client/issues/523
	go mod edit -replace "github.com/apache/thrift=github.com/apache/thrift@a9b748bb0e02"
	go mod tidy -v
	go mod vendor -v