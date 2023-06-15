#Http-route Using Linkerd

## Create and configure the namespace
Create and configure namespace with linkerd proxy so that Deployments created in the namespace will be automatically configured with linkerd sidecar/proxy
  ```
  kubectl create ns test --dry-run=client -o yaml \
    | linkerd inject - \
    | kubectl apply -f -
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