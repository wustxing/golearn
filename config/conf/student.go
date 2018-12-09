package conf

import "github.com/0990/golearn/config/gameproto"

type StudentConf struct {
	studentMap map[int32]*gameproto.StudentItemDefine
}

func (p *StudentConf) loadConf(conf *gameproto.StudentConfig) error {
	p.studentMap = make(map[int32]*gameproto.StudentItemDefine)

	for _, v := range conf.GetStudentItem() {
		p.studentMap[v.ID] = v
	}
	return nil
}

func (p *StudentConf) GetStudent(id int32) (*gameproto.StudentItemDefine, bool) {
	v, ok := p.studentMap[id]
	return v, ok
}
