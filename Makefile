RELEASE_NAME=service
IMAGE=$(RELEASE_NAME)

HELM_DIR=$(shell pwd)/helm
NAMESPACE=default

build:
	docker build -t ${RELEASE_NAME} .

helm.check:
	helm template -f ${HELM_DIR}/values.local.yaml ${HELM_DIR}

deploy:
	helm upgrade --install ${RELEASE_NAME} ${HELM_DIR} --namespace ${NAMESPACE} -f ${HELM_DIR}/values.local.yaml --debug

stop:
	helm delete ${RELEASE_NAME} --purge

status:
	kubectl get deployments --namespace ${NAMESPACE} -l service=${RELEASE_NAME}

pods:
	kubectl get po --namespace ${NAMESPACE} -l service=${RELEASE_NAME}
