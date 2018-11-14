package conf

import "github.com/0990/golearn/config/gameproto"

type TeacherConf struct {
	teacherMap map[int32]*gameproto.TeacherItemDefine
}

func (p *TeacherConf) loadConf(conf *gameproto.StudentConfig) error {
	p.teacherMap = make(map[int32]*gameproto.TeacherItemDefine)

	for _, v := range conf.GetTeacherItem() {
		p.teacherMap[v.ID] = v
	}
	return nil
}

func (p *TeacherConf) GetTeacher(id int32) (*gameproto.TeacherItemDefine, bool) {
	v, ok := p.teacherMap[id]
	return v, ok
}
