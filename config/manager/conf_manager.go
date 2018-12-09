package manager

import (
	"fmt"
	"github.com/0990/golearn/config/conf"
	"sync"
)

type ConfManager struct {
	path string
	cfg  *conf.Config
	//mutex sync.RWMutex
}

func (p *ConfManager) Init(path string) error {
	p.cfg = &conf.Config{}
	p.path = path
	err := p.cfg.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func (p *ConfManager) Reload(command string) error {
	var loadError error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reloadCfg := &conf.Config{}
		err := reloadCfg.Reload("student_config.pbt")
		if err != nil {
			loadError = err
			fmt.Println("reload error,use last config")
			return
		}

		//p.mutex.Lock()
		//defer p.mutex.Unlock()

		switch command {
		case "student":
			p.cfg.StudentConf = reloadCfg.StudentConf
		case "teacher":
			p.cfg.TeacherConf = reloadCfg.TeacherConf
		case "all":
			p.cfg = reloadCfg
		}
	}()
	wg.Wait()

	if loadError != nil {
		return loadError
	}

	return nil
}

func (p *ConfManager) GetConfig() *conf.Config {
	//p.mutex.RLock()
	//p.mutex.RUnlock()
	return p.cfg
}
