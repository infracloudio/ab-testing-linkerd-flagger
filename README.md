# dynamicRoutingApp

## Prerequisites
- Kubernetes kind/minikuke cluster
- Cluster must be set up with Ingress controller 
   - [Setup Ingress on minikube](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)
   - [Setup Ingress on kind](https://dustinspecker.com/posts/test-ingress-in-kind/)
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


## Weight Based A/B Testing with Flagger
### Delete the HTTPRoute and TrafficSplit object
Delete all HTTPRoute and TrafficSplit objects as flagger creates them internally while testing
  ```
  $ make delete-httpRoute-traffisplit
  ```
### Follow following steps for weight Based A/B Testing
  1. Install flagger
     ```
     $ make setup-cluster-flagger
     ```
  2. Deploy book-svc and forwarder as follows:
     ```
     $ make deploy-flagger-release
     ```
  3. Apply flagger header based canary object
     ```
     $ kubectl apply -f flagger/weight-based.yaml
     ```
  4. Deploy Load generator for matrics analysis
     1. Deploy load generator   
        ```
        $ make deploy-load-generator
        ```
  5. Suppose you have made changes to the code or you have developed a different version of the API.
      1. Build the changes as follows:
         ```
         $ make build-flagger-release
         ```
      3. Deploy new release
         ```
         $ make patch-flagger-release
         ```
  6. Observe the progress of a release
     ```
     $ watch kubectl -n test get canary
     ```
   

    
## Header based A/B Testing

  1. Install flagger
     ```
     $ make setup-cluster-flagger
     ```
  2. Deploy book-svc and forwarder as follows:
     ```
     $ make deploy-flagger-release
     ```
  3. Apply flagger header based canary object
     ```
     $ kubectl apply -f flagger/header-based.yaml
     ```
  4. Configure and deploy Load generator for matrics analysis
     1. Add `NEW_VERSION_HEADER_KEY` with the same value as `HEADER` env variable in the `deploy/book-svc.yaml` i.e `x-backend`. <br />
        **NOTE :** we have configure this header value as `x-backend` in flagger/header-based.yaml `headers` field
        ```
        - name: NEW_VERSION_HEADER_KEY
          value: "x-backend"
        ```
      
     2. Add value `new` for header key `x-backend` to route the traffic to new release. <br />
        **NOTE :** We have specify this `new` value in flagger/weight-based.yaml file to configure traffic to new release
        ```
        - name: NEW_VERSION_HEADER_VAL
          value: "new"
        ```
     3. Deploy load generator   
        ```
        $ make deploy-load-generator
        ```
  5. Suppose you have made changes to the code or you have developed a different version of the API.
      Deploy new release
      ```
      $ make patch-flagger-release
      ```
  6. Test Application with Request Header
     - We can dynamically route the traffic to book-svc-v1 by using the request header `x-backend: new` 
     - If we want to configure dynamic routing with some different header we can do it by editing the header name and value in the flagger/header-based.yaml file
     - Use `x-backend: new` header to control traffic
       ```
       $ curl -sX GET -H 'x-backend: new' http://app.example.com/
       ```
  7. Observe the progress of a release
     ```
     $ watch kubectl -n test get canary
     ```