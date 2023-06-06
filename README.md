# dynamicRoutingApp

## Prerequisites
- Kubernetes kind/minikuke cluster
- Linkerd must be installed and configured with the cluster(https://linkerd.io/2.13/tasks/install/)
- [Linkerd-SMI](https://linkerd.io/2.13/tasks/linkerd-smi/#cli) and [Linkerd-viz](https://linkerd.io/2.13/tasks/troubleshooting/#l5d-viz-ns-exists) extension must be installed.

## Create and configure namespace
Create and configure namespace with linkerd proxy so that Deployments created in the namespace will be automatically configured with linkerd sidecar/proxy
  ```
  kubectl create ns test --dry-run=client -o yaml \
    | linkerd inject - \
    | kubectl apply -f -
  ```


## Build the docker images
For kind cluster build all the required images.
  ```
  $ make build-all
  $ make load-kind
  ```
For minikube cluster build the required images as follows:
  ```
  $ eval $(minikube docker-env)
  $ make build-all
  ```
## Deploy Application on the kubernetes cluster
Deploy backend-a and backend-b 
  ```
  $ make deploy-all
  ```

NOTE: We should not enable Dynamic routing or traffic-splitting simultaneously.
## Header based Dynamic Routing
### Enable Dynamic Routing
  - To dynamically route traffic to another version of application i.e backend-b we need configure httpRoute with the following
    ```
    $ make configure-httpRoute
    ```
### Test With Request Header
  - We can dynamically route the traffic to backend-b by using the request header `x-backend: test` 
  - If we want to configure dynamic routing with some different header we can do it by editing the header name and value in the linkerd/httpRoute.yaml file
  - Expose forwarder service
    ```
    $ kubectl port-forward svc/forwarder-service 4000:8080 -n test
    ```
  - Use `x-backend: test` header to control traffic
    ```
    $ curl -sX GET -H 'x-backend: test' localhost:4000/
    ```

## Traffic Splitting
- We can enable traffic split by the ratio 50:50 between the two versions/backend as follows:
  ```
  $ make configure-traffic-split
  ```
- Expose forwarder service
  ```
  $ kubectl port-forward svc/forwarder-service 4000:8080 -n test
  ```
- Test traffic-split with the `curl -sX GET localhost:4000/` such that request will be routed to both version/backend according traffic split ratio.
- We can adjust ratio of traffic by editing the value of weight in the linkerd/traffic-split.yaml file
