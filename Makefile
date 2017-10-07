build:
	go build

test:
	docker-compose up -d

destroy:
	docker-compose stop && docker-compose rm -f

dynamictest:
	python test_lab.py build

dynamicdestroy:
	python test_lab.py destroy
