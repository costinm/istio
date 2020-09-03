#!/bin/bash

if [[ -n ${PROJECT} ]]; then
  echo gcloud container clusters get-credentials ${CLUSTER} --zone ${ZONE} --project ${PROJECT}
  gcloud container clusters get-credentials ${CLUSTER} --zone ${ZONE} --project ${PROJECT}
  # TODO: check secret manager for a .kubeconfig - use it for non-GKE projects AND ingress secrets
  # AND citadel root CA
fi

# Disable webhook config patching - manual configs used, proper DNS certs means no cert patching needed.
# If running in KNative without a DNS cert - we may need it back, but we should require DNS certs.
# Even if user doesn't have a DNS name - they can still use an self-signed root and add it to system trust,
# to simplify
export VALIDATION_WEBHOOK_CONFIG_NAME=
export INJECTION_WEBHOOK_CONFIG_NAME=

# No longer needed.
#export DISABLE_LEADER_ELECTION=true

# No mTLS for control plane
export USE_TOKEN_FOR_CSR=true
export USE_TOKEN_FOR_XDS=true

# Disable the DNS-over-TLS server - no ports
export DNS_ADDR=

# TODO: parse service name and extra project, revision, cluster

export REVISION=${REV:-managed}

# TODO: should be auto-set now, verify safe to remove
export GKE_CLUSTER_URL=https://container.googleapis.com/v1/projects/${PROJECT}/locations/${ZONE}/clusters/${CLUSTER}

export CLUSTER_ID=${PROJECT}/${ZONE}/${CLUSTER}

# Emulate K8S - with one namespace per tenant
export ASM_CONTROL_PLANE_POD_NAMESPACE=${K_CONFIGURATION}
# Revision is equivalent with a deployment - unfortunately we can't get instance id.
export POD_NAME=${K_REVISION}-$(date +%N)
export ASM_CONTROL_PLANE_POD_NAME=${POD_NAME}

# Test: see the IP, if unique we can add it to pod name
#ip addr
#hostname

if [[ "${CA}" == "1" ]]; then
  export CA_ADDR=meshca.googleapis.com:443
  export TRUST_DOMAIN=${PROJECT}.svc.id.goog
  export AUDIENCE=${TRUST_DOMAIN}
  export CA_PROVIDER=${CA_PROVIDER:-istiod}
else
  # If not set - the template default is a made-up istiod address instead of discovery.
  # TODO: fix template
  # TODO: if we fetch MeshConfig from cluster - leave trust domain untouched.
  export CA_ADDR=${K_SERVICE}${ISTIOD_DOMAIN}:443
  export TRUST_DOMAIN=cluster.local
  export AUDIENCE=${PROJECT}.svc.id.goog
  export CA_PROVIDER=istiod
fi

# TODO:
# - copy inject template and mesh config to cluster (first time) or from cluster
# - revision support
# - option to enable 'default' ingress class, remote install/control Gateway

kubectl get ns istio-system
if [[ "$?" != "0" ]]; then
  echo "Initializing istio-system and CRDs, fresh cluster"
  kubectl create ns istio-system
  #kubectl apply -k github.com/istio/istio/manifests/charts/base
  kubectl apply -f /var/lib/istio/config/gen-istio-cluster.yaml \
      --record=false --overwrite=false   --force-conflicts=true --server-side
fi
# TODO: check CRD revision, upgrade if needed.

if [[ -n ${MESH} ]]; then
  echo ${MESH} > /etc/istio/config/mesh
else
  cat /etc/istio/config/mesh_template.yaml | envsubst > /etc/istio/config/mesh
  cat /etc/istio/config/mesh
fi

# TODO: fix OSS template to use only MeshConfig !
cat /var/lib/istio/inject/values_template.yaml | envsubst > /var/lib/istio/inject/values

# TODO: istio must watch it - no file reloading
kubectl get -n istio-system cm istio-${REVISION}
if [[ "$?" != "0" ]]; then
  echo "Initializing revision"
  kubectl -n istio-system create cm istio-${REVISION} --from-file /etc/istio/config/mesh

  # Sidecars will report to stackdriver - requires proper setup.
  if [[ "${ASM}" == "1" ]]; then
    cat /var/lib/istio/config/telemetry-sd.yaml | envsubst | kubectl apply -f -
  else
    # Prometheus only.
    cat /var/lib/istio/config/telemetry.yaml | envsubst | kubectl apply -f -
  fi
fi


# Make sure the mutating webhook is installed, and prepare CRDs
# This also 'warms' up the kubeconfig - otherwise gcloud will slow down startup of istiod.
kubectl get mutatingwebhookconfiguration istiod-${REVISION}
if [[ "$?" == "1" ]]; then
  echo "Mutating webhook missing, initializing"
  cat /var/lib/istio/inject/mutating_template.yaml | envsubst > /var/lib/istio/inject/mutating.yaml
  cat /var/lib/istio/inject/mutating.yaml
  kubectl apply -f /var/lib/istio/inject/mutating.yaml
else
  echo "Mutating webhook found"
fi

echo Starting $*

# What audience to expect for Citadel and XDS - currently using the non-standard format
# TODO: use https://... - and separate token for stackdriver/managedCA
export TOKEN_AUDIENCES=${PROJECT}.svc.id.goog,istio-ca

# Istiod will report to stackdriver
export ENABLE_STACKDRIVER_MONITORING=${ENABLE_STACKDRIVER_MONITORING:-1}

env

exec /usr/local/bin/pilot-discovery discovery \
   --httpsAddr "" \
   --trust-domain ${TRUST_DOMAIN} \
   --secureGRPCAddr "" \
   --monitoringAddr "" \
   --grpcAddr "" \
   ${EXTRA_ARGS} ${LOG_ARGS} $*
