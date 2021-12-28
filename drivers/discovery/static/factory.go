package static

import (
	"reflect"

	"github.com/eolinker/eosc"
)

var name = "discovery_static"

//Register 注册静态服务发现的驱动工厂
func Register(register eosc.IExtenderDriverRegister) {
	register.RegisterExtenderDriver(name, NewFactory())
}

type factory struct {
}

//NewFactory 创建静态服务发现的驱动工厂
func NewFactory() eosc.IExtenderDriverFactory {
	return &factory{}
}

//Create 创建静态服务发现驱动
func (f *factory) Create(profession string, name string, label string, desc string, params map[string]interface{}) (eosc.IExtenderDriver, error) {
	return &driver{
		profession: profession,
		name:       name,
		label:      label,
		desc:       desc,

		configType: reflect.TypeOf((*Config)(nil)),
	}, nil
}