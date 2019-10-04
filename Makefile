REPOSITORY=krol
TAG=1.0
NAME=weather-operator

create:
	operator-sdk new ${NAME} --repo github.com/krol3/${NAME}

# Custom Resource Definition
crd:
	operator-sdk add api --api-version=k8s.devopsdays.com/v1alpha1 --kind=WeatherReport

# Custom Controller
controller:
	operator-sdk add controller --api-version=k8s.devopsdays.com/v1alpha1 --kind=WeatherReport

# Generate code for the resource type
generate:
	operator-sdk generate k8s

publish:
	operator-sdk build ${REPOSITORY}/${NAME}:${TAG}
	docker push ${REPOSITORY}/${NAME}:${TAG}

deploy_:
# Replace image
	sed -i "s|REPLACE_IMAGE|${REPOSITORY}/${NAME}:|g" deploy/operator.yaml
	kubectl create -f deploy/crds/k8s_v1alpha1_weatherreport_crd.yaml
	kubectl create -f deploy/

deploy_cr:
	kubectl create -f deploy/crds/k8s_v1alpha1_weatherreport_cr.yaml

delete:
	kubectl delete -f deploy/crds
	kubectl delete -f deploy/