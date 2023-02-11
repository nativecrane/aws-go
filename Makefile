############################
# TEST
############################

test:
	@cp .env.test api/.env
	@cd api; go test
	@rm api/.env

############################
# INCREMENTAL DEPLOYMENT
############################

dev: build archive-dev
	@sls deploy function -f api --stage=dev

uat: build archive-uat
	@sls deploy function -f api --stage=uat

prod: build archive-prod
	@sls deploy function -f api --stage=prod

############################
# FULL DEPLOYMENT
############################

deploy-dev: build archive-dev
	@sls deploy --stage=dev
	@rm api/lambda.zip

deploy-uat: build archive-uat
	@sls deploy --stage=uat
	@rm api/lambda.zip

deploy-prod: build archive-prod
	@sls deploy --stage=prod
	@rm api/lambda.zip

############################
# PRIVATE
############################

build:
	@cd api; env GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

archive-dev:
	@cp .env.dev api/.env
	@cd api; zip lambda.zip bootstrap .env
	@cd api; rm bootstrap
	@rm api/.env

archive-uat:
	@cp .env.uat api/.env
	@cd api; zip lambda.zip bootstrap .env
	@cd api; rm bootstrap
	@rm api/.env

archive-prod:
	@cp .env.prod api/.env
	@cd api; zip lambda.zip bootstrap .env
	@cd api; rm bootstrap
	@rm api/.env