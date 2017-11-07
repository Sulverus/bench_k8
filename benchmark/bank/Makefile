RELEASE_NAME=bank-bench
IMAGE=$(RELEASE_NAME)

build:
	rm -f *.pyc
	rm -f ammo.txt
	python datagen.py
	docker build -t ${RELEASE_NAME} .

run: build
	kubectl create -f bench.yaml

stop:
	kubectl delete -f bench.yaml
