# Golang-playground webapp

This repository contains a simple go based webapp, and the required manifests to deploy it onto Kubernetes.

## Webapp
The web application serves any requested url, except `/foo` and `/bar` that are "secured" with HTTP Basic authentication.

The http authentication logic is applied using [github.com/goji/httpauth ](https://github.com/goji/httpauth) and it uses `gopkg.in/yaml.v2` to parse an external config file for a list of password protected pages and the credentials name ([config.yaml](./webapp/config.yaml)).

## Monitoring (Next steps)
- Implement prometheus go client: https://prometheus.io/docs/guides/go-application/
