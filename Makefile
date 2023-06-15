build-a:
	docker build -t backend-a -f Dockerfile.backend-a .

build-a-v1:
	docker build -t backend-a:v1 -f Dockerfile.backend-a-v1 .

build-forwarder:
	docker build -t forwarder -f Dockerfile.forwarder .

build-load-generator:
	docker build -t load-generator -f Dockerfile.load-generator .

build-all: build-a build-a-v1 build-forwarder build-load-generator 

load-kind:
	kind load docker-image backend-a
	kind load docker-image forwarder
	kind load docker-image load-generator
	kind load docker-image backend-a:v1

delete-httpRoute-traffisplit:
	kubectl delete -f linkerd/httpRoute.yaml
	kubectl delete -f linkerd/traffic-split.yaml

deploy-flagger-release:
	kubectl apply -f deploy/backend-A/
	kubectl apply -f deploy/forwarder/

deploy-all:
	kubectl apply -f deploy/backend-A -f deploy/backend-a-v1 -f deploy/forwarder

deploy-load-generator:
	kubectl apply -f deploy/load-generator

configure-httpRoute:
	kubectl apply -f linkerd/httpRoute.yaml
configure-traffic-split:
	kubectl apply -f linkerd/traffic-split.yaml

setup-cluster-linkerd:
	./linkerd/setup-cluster.sh

setup-cluster-flagger:
	kubectl apply -k github.com/fluxcd/flagger/kustomize/linkerd

patch-flagger-release:
	kubectl patch deployment backend-a -n test --type='json' -p="$$(cat flagger/flagger-release-patch.json)"
