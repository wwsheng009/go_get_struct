package generator

import (
	"fmt"
	"github.com/cartmanis/go_get_struct/node"
	"os"
	"strings"
)

func CreateStruct(n *node.NodeXml) (*os.File, error) {
	goNode(n)
	return nil, nil
}

func goNode(n *node.NodeXml) (*os.File, error) {
	r := string('\u0060')
	st := fmt.Sprintf(`type %v struct {
	XMLName xml.Name %vxml:"%v"%v
}`, strings.Title(n.Namespace), r, n.Namespace, r)
	fmt.Println(st)
	return nil, nil
}

func getStruct(name string, filds []string) string {
	return ""
}
