package access

import (
	"fmt"

	"github.com/eolinker/eosc"
)

var (
	permits = eosc.BuildUntyped[string, *permit]()
)

type Template struct {
	Name     string     `yaml:"name" json:"name,omitempty"`
	CName    string     `yaml:"cname" json:"cname,omitempty"`
	Value    string     `yaml:"value" json:"value,omitempty"`
	Children []Template `yaml:"children" json:"children,omitempty"`
}

type permit struct {
	group string
	// permits 当前权限下的API列表
	permits eosc.Untyped[string, []string]
	// access api对应需要的权限
	access eosc.Untyped[string, string]
	// template 模版
	template []Template
}

func newPermit(group string, access []Access) *permit {
	p := &permit{
		group:    group,
		permits:  eosc.BuildUntyped[string, []string](),
		access:   eosc.BuildUntyped[string, string](),
		template: nil,
	}
	p.Add(access)
	return p
}

func (p *permit) Valid(access string) error {
	_, has := p.access.Get(access)
	if !has {
		return fmt.Errorf("permit %s not found", access)
	}
	return nil
}

func (p *permit) Add(as []Access) error {
	result, templates := formatAccess(as)
	for k, vs := range result {
		k = fmt.Sprintf("%s.%s", p.group, k)
		p.permits.Set(k, vs)
		for _, v := range vs {
			p.access.Set(v, k)
		}
	}
	p.template = templates
	return nil
}

func (p *permit) GetTemplate() []Template {
	return p.template
}

func (p *permit) GetPermits(access string) ([]string, error) {
	perms, has := p.permits.Get(access)
	if !has {
		return nil, fmt.Errorf("permit %s not found", access)
	}
	return perms, nil
}

func (p *permit) AccessKeys() []string {
	return p.permits.Keys()
}

func GetPermit(group string) (*permit, bool) {
	return permits.Get(group)
}
