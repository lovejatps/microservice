package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"microservice/conf"
	"sync"
)

var(
	masterEngine	 *xorm.Engine
	slaveEngine 	 *xorm.Engine
	lock			 sync.Mutex
)

func InstanceMasterEngin()  *xorm.Engine {
	if masterEngine !=nil{
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine !=nil{
		return masterEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User,c.Pwd,c.Host,c.Port,c.Dbname)
	engine ,err := xorm.NewEngine(conf.DriverName,driveSource)
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000) //LRU算法的一个缓存，缓存方式是存放到内存中，缓存struct的记录数为1000条
	engine.SetDefaultCacher(cacher)
	if err !=nil{
		log.Fatal("dbheoper.InstanceMasterEngin error",err)
		return nil
	}else{
		masterEngine = engine
		return masterEngine
	}

}

func InstanceSlaveEngin() *xorm.Engine{
	if slaveEngine != nil{
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()
	if slaveEngine != nil{
		return slaveEngine
	}
	c := conf.SlaveDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User,c.Pwd,c.Host,c.Port,c.Dbname)
	engine ,err := xorm.NewEngine(conf.DriverName,driveSource)
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000) //LRU算法的一个缓存，缓存方式是存放到内存中，缓存struct的记录数为1000条
	engine.SetDefaultCacher(cacher)
	if err!= nil {
		log.Fatal("dbheoper.InstanceMasterEngin error",err)
		return nil
	}else{
		 slaveEngine = engine
		return slaveEngine
	}
}
