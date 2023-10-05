# Proximo

A simple, lightweight HTTP proxy to allow anyone to access cloud run endpoints behind authentication. This will proxy 
requests to the cloud run endpoint, and add the `Authorization` header to the request based on your credentials.

## Usage

### Login

You need to ensure that you are logged in to your Google account before using Proximo.

```shell
gcloud auth login
```

### Start the proxy

```shell
proximo -url https://my-cloud-run-endpoint.run.app
```

This will start the proxy on port 8080, proxying requests to the given Cloud Run URL and adding the Authorization 
header with your Google credentials. You should be able to access the endpoint at http://localhost:8080.
