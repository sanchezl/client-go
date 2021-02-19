module github.com/openshift/client-go

go 1.13

require (
	github.com/openshift/api v0.0.0-20210216211028-bb81baaf35cd
	github.com/spf13/pflag v1.0.5
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/api v0.20.0
	k8s.io/apimachinery v0.20.0
	k8s.io/client-go v0.20.0
	k8s.io/code-generator v0.20.0
)

replace github.com/openshift/api => github.com/sanchezl/api v0.0.0-20210222165532-3c943607d192
