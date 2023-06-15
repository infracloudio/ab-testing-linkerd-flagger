# Header based And Weight Based A/B Testing with Flagger
## Delete the HTTPRoute and TrafficSplit object
Delete all HTTPRoute and TrafficSplit objects as flagger creates them internally while testing
  ```
  $ make delete-httpRoute-traffisplit
  ```
## Header based A/B Testing
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
     $ kubectl apply -f flagger/header-based.yaml
     ```
  4. Configure and deploy Load generator for matrics analysis
     1. Add `NEW_VERSION_HEADER_KEY` with the same value as `HEADER` env variable in the `deploy/backend-a.yaml` i.e `x-backend`. <br />
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
      2. Copy the image tag from the above image build and replace it with image tag in the value field of flagger/flagger-release-patch.json
         ```
         {
            "op": "replace",
            "path": "/spec/template/spec/containers/0/image",
            "value": "<image-name>:<tag>"
         }
         ```
         we are replacing image with new build and version as 'canary' for the deployment backend-a
      3. Deploy new release
         ```
         $ make patch-flagger-release
         ```
  6. Observe the progress of a release
     ```
     $ watch kubectl -n test get canary
     ```