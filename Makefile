build-a:
	docker build -t backend-a -f Dockerfile.backend-a .

build-b:
	docker build -t backend-b -f Dockerfile.backend-b .

build-forwarder:
	docker build -t forwarder -f Dockerfile.forwarder .

build-load-generator:
	docker build -t load-generator -f Dockerfile.load-generator .

build-flagger-release:
	@TAG=$$(uuidgen | tr '[:upper:]' '[:lower:]'); \
    docker build -t backend-a:$$TAG -f Dockerfile.backend-b .

build-all: build-a build-b build-forwarder build-load-generator


load-kind:
	kind load docker-image backend-a
	kind load docker-image backend-b
	kind load docker-image forwarder
	kind load docker-image load-generator

deploy-flagger-release:
	kubectl apply -f flagger/flagger-release.yaml

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

make path-flagger-release:
	kubectl patch deployment backend-a -n test --type='json' -p="$$(cat flagger/flagger-release-patch.json)"
