package conf

import (
	"github.com/0990/golearn/config/gameproto"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

type Config struct {
	StudentConf
	TeacherConf
	cfgNode *cfgConfigNode
}

type cfgConfigNode struct {
	studentCfg *gameproto.StudentConfig
}

func (p *Config) Load(path string) error {
	//解析配置文件到内存对象中
	p.cfgNode = &cfgConfigNode{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	p.cfgNode.studentCfg = &gameproto.StudentConfig{}
	err = proto.UnmarshalText(string(data), p.cfgNode.studentCfg)
	if err != nil {
		return err
	}

	//从内存对象中分离出所用的配置信息
	err = p.StudentConf.loadConf(p.cfgNode.studentCfg)
	if err != nil {
		return err
	}
	err = p.TeacherConf.loadConf(p.cfgNode.studentCfg)
	if err != nil {
		return err
	}
	return nil
}

func (p *Config) Reload(path string) error {
	//解析配置文件到内存对象中
	p.cfgNode = &cfgConfigNode{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	p.cfgNode.studentCfg = &gameproto.StudentConfig{}
	err = proto.UnmarshalText(string(data), p.cfgNode.studentCfg)
	if err != nil {
		return err
	}

	//从内存对象中分离出所用的配置信息
	err = p.StudentConf.loadConf(p.cfgNode.studentCfg)
	if err != nil {
		return err
	}
	err = p.TeacherConf.loadConf(p.cfgNode.studentCfg)
	if err != nil {
		return err
	}
	return nil
}
