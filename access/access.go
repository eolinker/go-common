package access

import (
	"fmt"
	"strings"
)

type Access struct {
	Name  string `yaml:"name" json:"name,omitempty"`
	CName string `yaml:"cname" json:"cname,omitempty"`
	Desc  string `yaml:"desc" json:"desc,omitempty"`
}

var (
	access = make(map[string][]Access)
)

func All() map[string][]Access {
	return access
}
func Get(name string) ([]Access, bool) {
	list, has := access[name]
	return list, has
}

func Add(group string, asl []Access) {
	gl := make([]Access, 0, len(asl))
	group = formatGroup(group)
	gp := fmt.Sprint(group, ".")
	for _, a := range asl {
		a.Name = strings.ToLower(a.Name)
		if !strings.HasPrefix(a.Name, gp) {
			a.Name = fmt.Sprint(gp, a.Name)
		}
		gl = append(gl, a)
	}

	access[group] = append(access[group], gl...)
}
func formatGroup(group string) string {
	group = strings.ToLower(group)
	group = strings.TrimSpace(group)
	group = strings.Trim(group, ".")
	group = strings.ReplaceAll(group, "-", "_")
	group = strings.ReplaceAll(group, ".", "_")

	return group
}
