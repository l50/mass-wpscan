install:
	go build

setup:
	@echo "===> Installing deps"
	go get -u github.com/fatih/color

test:
	docker-compose up -d

destroy:
	docker-compose stop && docker-compose rm -f

dynamictest:
	python test_lab.py build

dynamicdestroy:
	python test_lab.py destroy
