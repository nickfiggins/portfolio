.PHONY: clean build package deploy

default: test

clean:
	

test:

build:
	@cd ./ui && ng build
	docker-compose build ui

deploy:
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 347362381652.dkr.ecr.us-east-1.amazonaws.com
	docker tag portfolio:prod 347362381652.dkr.ecr.us-east-1.amazonaws.com/portfolio:latest
	docker push 347362381652.dkr.ecr.us-east-1.amazonaws.com/portfolio:latest
	aws ecs update-service --force-new-deployment --service portfolio-ecs-task --cluster portfolio-spot-v1