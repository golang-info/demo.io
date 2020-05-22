[Writing Operators using the Operator Framework SDK](https://linuxera.org/writing-operators-using-operator-framework/)

```bash
operator-sdk new reverse-words-operator --repo=github.com/golang-info/everse-words-operator
INFO[0000] Creating new Go operator 'reverse-words-operator'.
INFO[0000] Created go.mod
INFO[0000] Created tools.go
INFO[0000] Created cmd/manager/main.go
INFO[0000] Created build/Dockerfile
INFO[0000] Created build/bin/entrypoint
INFO[0000] Created build/bin/user_setup
INFO[0000] Created deploy/service_account.yaml
INFO[0000] Created deploy/role.yaml
INFO[0000] Created deploy/role_binding.yaml
INFO[0000] Created deploy/operator.yaml
INFO[0000] Created pkg/apis/apis.go
INFO[0000] Created pkg/controller/controller.go
INFO[0000] Created version/version.go
INFO[0000] Created .gitignore
INFO[0000] Validating project
go: downloading k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
go: downloading sigs.k8s.io/controller-runtime v0.2.0
go: downloading github.com/operator-framework/operator-sdk v0.11.0
go: downloading k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
go: downloading github.com/prometheus/client_golang v1.0.0
go: downloading k8s.io/klog v0.3.3
go: downloading gomodules.xyz/jsonpatch/v2 v2.0.1
go: downloading github.com/Azure/go-autorest v11.7.0+incompatible
go: downloading golang.org/x/crypto v0.0.0-20190621222207-cc06ce4a13d4
go: downloading github.com/gophercloud/gophercloud v0.2.0
go: downloading k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
go: downloading cloud.google.com/go v0.37.2
go: downloading github.com/googleapis/gnostic v0.2.0
go: downloading github.com/go-logr/zapr v0.1.1
go: downloading github.com/prometheus/common v0.4.1
go: downloading github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
go: downloading github.com/coreos/prometheus-operator v0.31.1
go: downloading golang.org/x/sys v0.0.0-20190515120540-06a5c4944438
go: downloading contrib.go.opencensus.io/exporter/ocagent v0.4.12
go: downloading go.opencensus.io v0.20.2
go: downloading k8s.io/kube-openapi v0.0.0-20190401085232-94e1e7b7574c
go: downloading go.uber.org/atomic v1.3.2
go: downloading google.golang.org/grpc v1.19.1
go: downloading github.com/census-instrumentation/opencensus-proto v0.2.0
go: downloading google.golang.org/api v0.3.2
go: downloading github.com/prometheus/procfs v0.0.2
go: downloading github.com/grpc-ecosystem/grpc-gateway v1.8.5
go: downloading google.golang.org/genproto v0.0.0-20190404172233-64821d5d2107
go: downloading gopkg.in/yaml.v2 v2.2.2
go: downloading k8s.io/kube-state-metrics v1.7.2
go: downloading github.com/emicklei/go-restful v2.8.1+incompatible
go: downloading github.com/go-openapi/swag v0.19.0
go: downloading github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983
INFO[0200] Project validation successful.
INFO[0200] Project creation complete.


~/go_demo/demo.io/demo-k8s/06-operators-operator-framework-sdk/reverse-words-operator$ operator-sdk add api --api-version=linuxera.org/v1alpha1 --kind=ReverseWordsApp
INFO[0000] Generating api version linuxera.org/v1alpha1 for kind ReverseWordsApp.
INFO[0000] Created pkg/apis/linuxera/group.go
INFO[0020] Created pkg/apis/linuxera/v1alpha1/reversewordsapp_types.go
INFO[0020] Created pkg/apis/addtoscheme_linuxera_v1alpha1.go
INFO[0020] Created pkg/apis/linuxera/v1alpha1/register.go
INFO[0020] Created pkg/apis/linuxera/v1alpha1/doc.go
INFO[0020] Created deploy/crds/linuxera.org_v1alpha1_reversewordsapp_cr.yaml
INFO[0021] Created deploy/crds/linuxera.org_reversewordsapps_crd.yaml
INFO[0021] Running deepcopy code-generation for Custom Resource group versions: [linuxera:[v1alpha1], ]
INFO[0030] Code-generation complete.
INFO[0030] Running OpenAPI code-generation for Custom Resource group versions: [linuxera:[v1alpha1], ]
INFO[0040] Created deploy/crds/linuxera.org_reversewordsapps_crd.yaml
INFO[0040] Code-generation complete.
INFO[0040] API generation complete.

curl -Ls https://linuxera.org/assets/post_resources/2019-05-18-openshift-operator-framework/reversewordsapp_types.go.txt -o pkg/apis/linuxera/v1alpha1/reversewordsapp_types.go

~/go_demo/demo.io/demo-k8s/06-operators-operator-framework-sdk/reverse-words-operator$ operator-sdk generate k8s
INFO[0000] Running deepcopy code-generation for Custom Resource group versions: [linuxera:[v1alpha1], ]
INFO[0009] Code-generation complete.

~/go_demo/demo.io/demo-k8s/06-operators-operator-framework-sdk/reverse-words-operator$ operator-sdk generate openapi
INFO[0000] Running OpenAPI code-generation for Custom Resource group versions: [linuxera:[v1alpha1], ]
API rule violation: list_type_missing,demo.io/demo-k8s/06-operators-operator-framework-sdk/reverse-words-operator/pkg/apis/linuxera/v1alpha1,ReverseWordsAppStatus,AppPods
INFO[0010] Created deploy/crds/linuxera.org_reversewordsapps_crd.yaml
INFO[0010] Code-generation complete.

~/go_demo/demo.io/demo-k8s/06-operators-operator-framework-sdk/reverse-words-operator$ operator-sdk add controller --api-version=linuxera.org/v1alpha1 --kind=ReverseWordsApp
INFO[0000] Generating controller version linuxera.org/v1alpha1 for kind ReverseWordsApp.
INFO[0000] Created pkg/controller/reversewordsapp/reversewordsapp_controller.go
INFO[0000] Created pkg/controller/add_reversewordsapp.go
INFO[0000] Controller generation complete.
```