RELEASE_NAME=k6-bench
IMAGE=$(RELEASE_NAME)

HELM_DIR=$(shell pwd)/helm
NAMESPACE=default

build:
	docker build -t ${RELEASE_NAME} .

run:
	kubectl create -f bench.yaml

stop:
	kubectl delete -f bench.yaml
