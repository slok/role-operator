
CODE_GENERATOR_IMAGE=quay.io/slok/kube-code-generator:v1.10.0
CODE_GENERATOR_PACKAGE=github.com/slok/role-operator
GROUP_VERSION="roleoperator:v1alpha1"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=${DIR}/../../

echo "Generating Kubernetes code..."
docker run --rm -it \
	-v ${ROOT_DIR}:/go/src/${CODE_GENERATOR_PACKAGE} \
	-e PROJECT_PACKAGE=${CODE_GENERATOR_PACKAGE} \
	-e CLIENT_GENERATOR_OUT=${CODE_GENERATOR_PACKAGE}/pkg/client/k8s \
	-e APIS_ROOT=${CODE_GENERATOR_PACKAGE}/pkg/apis \
	-e GROUPS_VERSION="${GROUP_VERSION}" \
	-e GENERATION_TARGETS="deepcopy,client" \
	${CODE_GENERATOR_IMAGE}