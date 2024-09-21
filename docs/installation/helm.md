This guide provides step-by-step instructions to install tvhgo using helm.
Instead of a custom helm chart, we are using the [app-template chart](https://bjw-s.github.io/helm-charts/docs/app-template/) by bjw-s.

## Prerequisites

Before you start, make sure you have the following ready:

- Kubernetes cluster
- Helm

## Install tvhgo

Create a file called `values.yaml` with the following content:

```yaml
controllers:
  tvhgo:
    replicas: 1 # adjust to your needs
    strategy: RollingUpdate # or Recreate when used with sqlite
    containers:
      app:
        image:
          repository: ghcr.io/davidborzek/tvhgo
          tag: latest # should be a specific version in production
        env:
          # see https://davidborzek.github.io/tvhgo/latest/configuration/ for further configuration options
          TVHGO_SERVER_HOST: <TVHEADEND_HOST> # replace with the actual hostname or ip of your tvheadend server
          TVHGO_SERVER_HOST: "0.0.0.0"
          TVHGO_SERVER_PORT: &port 8080
          # metrics
          TVHGO_METRICS_ENABLED: "true"
          TVHGO_METRICS_PORT: &metricsPort 8081
          TVHGO_METRICS_PATH: &metricsPath "/metrics"
        probes:
          liveness:
            enabled: true
            custom: true
            spec: &probeSpec
              httpGet:
                path: /health/liveness
                port: *port
              initialDelaySeconds: 3
              periodSeconds: 10
              timeoutSeconds: 1
              failureThreshold: 3
          readiness:
            enabled: true
            custom: true
            spec:
              <<: *probeSpec
              httpGet:
                path: /health/readiness
                port: *port
        securityContext:
          allowPrivilegeEscalation: false
          readOnlyRootFilesystem: true
          capabilities: { drop: ["ALL"] }
        resources:
          requests:
            cpu: 10m
          limits:
            memory: 128Mi

defaultPodOptions:
  securityContext:
    runAsNonRoot: true
    runAsUser: 568
    runAsGroup: 568
    fsGroup: 568
    fsGroupChangePolicy: OnRootMismatch
    seccompProfile: { type: RuntimeDefault }

service:
  app:
    controller: tvhgo
    ports:
      http:
        port: *port
      metrics:
        port: *metricsPort

# Uncomment the following block to enable prometheus scraping
# serviceMonitor:
#  app:
#    serviceName: tvhgo
#    endpoints:
#      - port: metrics
#        scheme: http
#        path: *metricsPath
#        interval: 1m
#        scrapeTimeout: 10s

# Uncomment the following block to enable ingress
#ingress:
#  app:
#    hosts:
#      - host: tvhgo.yourdomain.com
#        paths:
#          - path: /
#            service:
#              identifier: app
#              port: http


# Uncomment the following block to enable persistence.
# You need a storage provider that supports dynamic provisioning.
# A better option for persistence is to use a external postgres database.
#persistence:
#  config:
#    type: persistentVolumeClaim
#    accessMode: ReadWriteOnce
#    size: 1Gi
#    globalMounts:
#      - path: /tvhgo
```

Replace `<TVHEADEND_HOST>` with the actual hostname or ip of your tvheadend server.

You can find more configuration options [here](../configuration.md).

Now you can install tvhgo using the following commands:

```bash
$ helm repo add bjw-s https://bjw-s.github.io/helm-charts
$ helm install tvhgo bjw-s/app-template -f values.yaml
```

You can find more configuration options [here](#configuration).

## Create a user

To complete the setup you need to create a user.

```bash
kubectl exec -it \
    $(kubectl get pods -n default -l app.kubernetes.io/name=tvhgo -o jsonpath='{.items[0].metadata.name}') \
    tvhgo admin user add
```

> Note: Replace `default` with the namespace you installed tvhgo in.

Follow the interactive setup to create a new user.
