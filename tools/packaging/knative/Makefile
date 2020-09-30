# Example and tests for running Istiod on CloudRun or KNative
# This is an experimental mode, to evaluate the k8s dependencies

# Project used for testing the integration.
ISTIO_REGION ?= us-central1
ISTIO_PROJECT ?= wlhe-cr
ISTIO_SA ?= 601426346923-compute@developer.gserviceaccount.com
ZONE ?= us-central1-c
ISTIO_SUFFIX?=-icq63pqnqq-uc.a.run.app
URL=istiod-${PROJECT_ID}-${CLUSTER}-${REV}${ISTIO_SUFFIX}

# Test project/cluster where workloads run.
PROJECT_ID ?= costin-asm1
CLUSTER ?= big1

CPU ?= 1
MEM ?= 2G

# Istiod revision to deploy
REV ?= oss

# Image tag
TAG ?= asm-cr

# Hub for the proxy
PROXY_HUB ?= grc.io/wlhe-cr

$(ISTIO_OUT)/knative/telemetry-sd.yaml: manifests/charts/istio-control/istio-discovery/templates/telemetryv2_1.8.yaml tools/packaging/knative/values.telemetry-sd.yaml
	mkdir -p $(ISTIO_OUT)/knative/templates
	cp manifests/charts/istio-control/istio-discovery/templates/telemetryv2_1.8.yaml ${ISTIO_OUT}/knative/templates
	cp manifests/charts/istio-control/istio-discovery/Chart.yaml ${ISTIO_OUT}/knative/Chart.yaml
	helm template $(ISTIO_OUT)/knative -f tools/packaging/knative/values.telemetry-sd.yaml > $@

$(ISTIO_OUT)/knative/telemetry.yaml: manifests/charts/istio-control/istio-discovery/templates/telemetryv2_1.8.yaml
	mkdir -p $(ISTIO_OUT)/knative/templates
	cp manifests/charts/istio-control/istio-discovery/templates/telemetryv2_1.8.yaml ${ISTIO_OUT}/knative/templates
	cp manifests/charts/istio-control/istio-discovery/Chart.yaml ${ISTIO_OUT}/knative/Chart.yaml
	helm template $(ISTIO_OUT)/knative -f tools/packaging/knative/values.telemetry.yaml > $@

# Setup the CloudRun Istiod revision
knative/setup:
	#gcloud config unset api_endpoint_overrides/run
	# Slow start - but pinned
	# Alternative: --cpu 4 --memory 4G
	gcloud alpha run deploy istiod-${PROJECT_ID}-${CLUSTER}-${REV} --allow-unauthenticated \
	 --project ${ISTIO_PROJECT} --region ${ISTIO_REGION} --platform managed \
     --image gcr.io/${ISTIO_PROJECT}/cloudrun:${TAG} \
     --set-env-vars=${ENVEXTRA}REV=${REV},TAG=${TAG},CLUSTER=${CLUSTER},ZONE=${ZONE},PROJECT=${PROJECT_ID},ISTIOD_DOMAIN=${ISTIO_SUFFIX} \
	 --port 8080 --concurrency 1000 --timeout 900 --cpu ${CPU} --memory ${MEM}

# Same, but with stackdriver and CA integrated
knative/setup_asm:
	REV=asm ENVEXTRA="ASM=1," $(MAKE) knative/setup
