# Kubernetes

### Contents
- [Kubernetes](#kubernetes)
    - [Pod](#pod)
    - [Networking](#networking)
    - [Controllers](#controllers)
    - [Managing State](#managing-state)
    - [Storage](#storage)
    - [Configuration and Secrets](#configuration-and-secrets)
    - [Authentication and Authorization](#authentication-and-authorization)
    - [Logging and Monitoring](#logging-and-monitoring)
    - [Ingress](#ingress)
    - [Running workloads to completion](#running-workloads-to-completion)
    - [TLS & Certificates](#tls-and-cerificates)
    - [Deployment and packaging of manifests](#deployment-and-packaging-of-manifests)

## Pod 
- InitContainers. 
- Probes
- PodPresets
- PodLifecycle states
#### Assignments
- [ ] Deploy a Pod which serves a static file. 
The static file should be created when the Pod starts (init container).
The pod is able to serve traffic only if the file exists. Gets restarted if the file is deleted. 
Also, configure liveness and readiness probes

## Networking 
- Fundamentals of the kubernetes networking approach. 
- The Pause container. 
- CNI - Container Networking Interface, CNI Plugins, what is the specification

## Controllers
- Deployment
- Daemonset
- Replicaset / Replication controller
#### Assignments
- [ ] Run a Pod on each node of the kubernetes cluster. 
- [ ] Run an nginx pod, upgrade the pod, rollback the deployment, record history.

## Managing state	
- Statefulsets	
#### Assignments
- [ ] Convert a Postgres deployment to a statefulset and deploy that on the cluster.

## Storage
- PV
- PVC
- StorageClass
#### Assignments
- [ ] Deploy a Pod which uses a scratch disk. 
- [ ] Deploy a Pod with multiple containers, One of the container creates random files, and the other list the contents of the created file.

## Configuration and Secrets	
- ConfigMap
- Secrets	
#### Assignments
- [ ] Deploy a pod which uses a secret. 
- [ ] Deploy a pod which uses the nginx reverse proxy configuration from a configmap"

## Authentication and Authorization	
- Service accounts and external users.  
- Authentication strategies. 
- RBAC
#### Assignments
- [ ] Deploy a Pod which uses a service account. The pod tries to use a configmap. Setup RBAC for the service account such that the service account does not have access to the configmap resource.

## Logging and Monitoring	
- Core metrics. 
- Metrics server. 
- Logging. 
- Custom metrics pipeline. 
- Horizontal Pod Autoscaler"	
#### Assignments
- [ ] Deploy an application with resource requests. 
- [ ] Configure HPA to scale based on the CPU utilization when it goes beyond a certain threshold"

## Ingress
- Ingress resource and ingress controllers

## Running workloads to completion	
- Job
- CronJob

## TLS & Certificates	
- Kubernetes PKI infrastructure.

## Deployment and packaging of manifests	
- Helm charts
