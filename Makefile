install:
	go build

buildtest:
	python test_lab.py build

destroytest:
	python test_lab.py destroy
