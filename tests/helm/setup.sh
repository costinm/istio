#/bin/bash

# Update or install istio-system with the config needed for the test cluster.
function testIstioUpgradeSystem() {
   (cd $TOP/src/istio.io/istio; helm upgrade istio-system \
    install/kubernetes/helm/istio \
    -f tests/helm/istio-system/values-small.yaml \
    --set global.tag=$TAG --set global.hub=$HUB $* )
}

function testIstioInstallSystem() {
    (cd $TOP/src/istio.io/istio; \
    helm install -n istio-system --namespace istio-system \
     install/kubernetes/helm/istio \
     -f tests/helm/istio-system/values-small.yaml )
}


# Update or install test ns with the config needed for the test cluster.
function testIstioUpgradeTest() {
   local F=${1:-"latest"}
   (cd $TOP/src/istio.io/istio; helm upgrade test \
   tests/helm  \
    --set fortioImage=istio/fortio:$F $*)
}

function testIstioInstallTest() {
    kubectl create ns test
    kubectl label namespace test istio-injection=enabled

    (cd $TOP/src/istio.io/istio; \
    helm install -n test --namespace test \
     tests/helm )
}

# Update or install load test ns with the config needed for the test cluster.
function testIstioUpgradeLoad() {
   (cd $TOP/src/istio.io/istio; helm upgrade load \
   tests/testdata/pilotload $*)
}

function testIstioInstallLoad() {
    kubectl create ns load
    kubectl label namespace load istio-injection=enabled
    (cd $TOP/src/istio.io/istio; \
    helm install -n load --namespace load \
     tests/testdata/pilotload )
}

# Install istio
function testInstall() {
    make istio-demo.yaml
    kubectl create ns istio-system


    kubectl -n test apply -f samples/httpbin/httpbin.yaml

    kubectl create ns bookinfo
    kubectl label namespace bookinfo istio-injection=enabled

    kubectl -n bookinfo apply -f samples/bookinfo/kube/bookinfo.yaml
}


# Setup DNS entries - currently using gcloud
# Requires DNS_PROJECT, DNS_DOMAIN and DNS_ZONE to be set
# For example, DNS_DOMAIN can be istio.example.com and DNS_ZONE istiozone.
# You need to either buy a domain from google or set the DNS to point to gcp.
# Similar scripts can setup DNS using a different provider
function testCreateDNS() {

    gcloud dns --project=$DNS_PROJECT record-sets transaction start --zone=$DNS_ZONE

    gcloud dns --project=$DNS_PROJECT record-sets transaction add ingress10.${DNS_DOMAIN}. --name=*.v10.${DNS_DOMAIN}. --ttl=300 --type=CNAME --zone=$DNS_ZONE

    gcloud dns --project=$DNS_PROJECT record-sets transaction execute --zone=$DNS_ZONE
}

# Set the gateway IP address in DNS
function testSetIngress() {
    local IP=$1

    gcloud dns --project=$DNS_PROJECT record-sets transaction start --zone=$DNS_ZONE

    gcloud dns --project=$DNS_PROJECT record-sets transaction add $IP --name=ingress10.${DNS_DOMAIN}. --ttl=300 --type=A --zone=$DNS_ZONE

    gcloud dns --project=$DNS_PROJECT record-sets transaction execute --zone=$DNS_ZONE
}
