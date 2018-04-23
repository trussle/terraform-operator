package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

/*package v1

import (
	"strings"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Policy describes an AWS IAM Policy
type Policy struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec PolicySpec `json:"spec"`
}

// PolicySpec is the spec for Policy
type PolicySpec struct {
	// The policy document. This is a json formatted string
	Document string `json:"policyDocument"`
	Name     string `json:"name"`
}
*/

func main() {
	prov := aws.Provider().(*schema.Provider)
	for k, v := range prov.ResourcesMap {
		foo := generateStructString(titleStr(k), v.Schema)
		fmt.Println(foo)
	}
}

func generateStructString(n string, m map[string]*schema.Schema) string {
	spec := fmt.Sprintf(`
type %sSpec struct {
`, n)
	for k, v := range m {
		spec = spec + generateFieldString(k, v) + "\n"
	}
	spec = spec + "}"
	return spec
}

func generateFieldString(k string, v *schema.Schema) string {
	name := titleStr(k)
	t := typeAsString(v.Type)
	if v.Computed {
		return ""
	}
	return fmt.Sprintf("\t%s\t%s\t`json:\"%s\"`", name, t, k)
}

func titleStr(s string) string {
	p := strings.Replace(s, "_", " ", -1)
	t := strings.Title(p)
	return strings.Replace(t, " ", "", -1)
}

func typeAsString(t schema.ValueType) string {
	switch t {
	case schema.TypeString:
		return "string"
	case schema.TypeBool:
		return "bool"
	case schema.TypeFloat:
		return "float"
	case schema.TypeInt:
		return "int"
	case schema.TypeList:
		return "[]interface{}"
	case schema.TypeMap:
		return "map[string]interface{}"
	}

	return "interface{}"
}
