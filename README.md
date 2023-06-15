# dynamicRoutingApp

## Prerequisites
- Kubernetes kind/minikuke cluster
- Linkerd must be installed and configured with the cluster(https://linkerd.io/2.13/tasks/install/)
- [Linkerd-SMI](https://linkerd.io/2.13/tasks/linkerd-smi/#cli) and [Linkerd-viz](https://linkerd.io/2.13/tasks/troubleshooting/#l5d-viz-ns-exists) extension must be installed.
- We can install Linkerd and required addons on the Kubernetes cluster in one step with
  ```
  $ make setup-cluster-linkerd
  ```

## Create and configure the namespace
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
Deploy backend-a, backend-b and Forwarder
  ```
  $ make deploy-all
  ```

NOTE: We should not enable Dynamic routing and traffic-splitting simultaneously.
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

## A/B Testing with Flagger
### Delete the HTTPRoute and TrafficSplit object
Delete all HTTPRoute and TrafficSplit objects as flagger creates them internally while testing
  ```
  $ make delete-httpRoute-traffisplit
  ```
### Implement A/B Testing
  1. Install flagger
     ```
     $ make setup-cluster-flagger
     ```
  2. Deploy backend-a and forwarder as follows:
     ```
     $ make deploy-flagger-release
     ```
  3. Apply flagger header based canary object
     ```
     $ kubectl apply -f flagger/weight-based.yaml
     ```
  4. Configure and deploy Load generator for matrics analysis
        ```
        $ make build-load-generator
        $ make load-kind
        $ make deploy-load-generator
        ```
  5. Suppose you have made changes to the code or you have developed a different version of the API.
      1. Build the changes as follows:
         ```
         $ make build-flagger-release
         $ make load-kind
         ```
      3. Deploy new release
         ```
         $ make patch-flagger-release
         ```
  6. Observe the progress of a release
     ```
     $ watch kubectl -n test get canary
     ```
    
