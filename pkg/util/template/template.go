package template

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/config"
	"github.com/openshift/openshift-azure/pkg/util/tls"
)

func Template(tmpl string, f template.FuncMap, cs *api.OpenShiftManagedCluster, extra interface{}) ([]byte, error) {
	t, err := template.New("").Funcs(template.FuncMap{
		"CertAsBytes":       tls.CertAsBytes,
		"CertChainAsBytes":  tls.CertChainAsBytes,
		"PrivateKeyAsBytes": tls.PrivateKeyAsBytes,
		"PublicKeyAsBytes":  tls.PublicKeyAsBytes,
		"YamlMarshal":       yaml.Marshal,
		"Base64Encode":      base64.StdEncoding.EncodeToString,
		"String":            func(b []byte) string { return string(b) },
		"quote":             strconv.Quote,
		"ImageOnly":         func(s string) string { return strings.Split(s, ":")[0] },
		"escape": func(b string) string {
			replacer := strings.NewReplacer("$", "\\$")
			return replacer.Replace(b)
		},
	}).Funcs(f).Parse(tmpl)
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}

	err = t.Execute(b, struct {
		ContainerService *api.OpenShiftManagedCluster
		Config           *api.Config
		Derived          interface{}
		Extra            interface{}
	}{ContainerService: cs, Config: &cs.Config, Derived: config.Derived, Extra: extra})
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
