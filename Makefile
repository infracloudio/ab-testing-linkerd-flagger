build-a:
	docker build -t backend-a -f Dockerfile.backend-a .

build-b:
	docker build -t backend-b -f Dockerfile.backend-b .

build-forwarder:
	docker build -t forwarder -f Dockerfile.forwarder .

build-load-generator:
	docker build -t load-generator -f Dockerfile.load-generator .

build-flagger-release:
	docker build -t backend-b:v1 -f Dockerfile.backend-b .

build-all: build-a build-b build-forwarder build-load-generator build-flagger-release


load-kind:
	kind load docker-image backend-a
	kind load docker-image backend-b
	kind load docker-image forwarder
	kind load docker-image load-generator
	kind load docker-image backend-b:v1

deploy-flagger-release:
	kubectl apply -f deploy/backend-A/
	kubectl apply -f deploy/forwarder/

deploy-all:
	kubectl apply -f deploy/backend-A -f deploy/backend-B -f deploy/forwarder

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
