package main

import (
	"gopkg.in/mgo.v2"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const url  = "mongodb://123.207.215.200:27017";

type Operater struct {
	mogSession *mgo.Session
	dbname string
	document string
}

type person struct {
	AGE    int  `bson:"age"`
	NAME   string `bson:"name"`
	HEIGHT int    `bson:"height"`
}

func main() {
	mgo.SetDebug(true)
	mgo.SetStats(true)
	op:=new(Operater)
	op.dbname="mytest"
	op.document="people"
	err:=op.connect()
	if err !=nil {
		fmt.Println("连接出错",err)
		return
	}

	p:=person{
		33,
		"周杰伦",
		175,
	}
	err=op.insert(p)
	if err != nil {
		fmt.Println("插入出错",err)
	}
	op.update()
	op.query();

	count,err:=op.count()
	if err !=nil {
		fmt.Println("统计出错",err)
		return
	}

	err=op.delete(&bson.M{"height": 0})
	if err!=nil {
		fmt.Println("删除错误",err)
	}else{
		fmt.Println("删除成功")
	}

	fmt.Println("共有数据",count)
	op.mogSession.Close()
}
//连接数据库
func (operater *Operater) connect() error {
	mogsession, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	operater.mogSession=mogsession
	return nil
}
//插入
func (operater *Operater) insert( p person) error {
	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
	err:=collcetion.Insert(p)
	return err
}
//统计文档中数据的个数
func (operater *Operater) queryAll() ([]person,error) {
	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
	p:=new(person)
	p.AGE=33
	query:=collcetion.Find(nil)
	ps:=[]person{}
	query.All(&ps)
	iter:=collcetion.Find(nil).Iter()
	//
	result:=new(person)
	for iter.Next(&result) {
		fmt.Println("一个一个输出：", result)
	}
	return ps,nil
}
//统计文档中数据的个数
func (operater *Operater) query() ([]person,error) {
	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
	p:=new(person)
	p.AGE=33
	query:=collcetion.Find(bson.M{"age":21})
	ps:=[]person{}
	query.All(&ps)
	fmt.Println(ps)
	return ps,nil
}

//更新一行
func (operater *Operater) update() (error) {
	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
	update:=person{
		33,
		"詹姆斯",
		201,
	}
	err:=collcetion.Update(bson.M{"name": "周杰伦"},update)
	if err !=nil {
		fmt.Println(err)
	}
	return err
}
//更新所有数据
func (operater *Operater) updateAll() (error) {
	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
	update:=person{
		33,
		"詹姆斯",
		201,
	}
	err:=collcetion.UpdateId(1,update)
	if err !=nil {
		fmt.Println(err)
	}
	changeinfo,err:=collcetion.UpdateAll(bson.M{"name": "周杰伦"},update)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println("共有多少行",changeinfo.Matched,"影响")
	return nil
}


//单行删除
func (operater *Operater) delete(seletor interface{}) (error) {
	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
	return collcetion.Remove(seletor)
}

//统计文档中数据的个数
func (operater *Operater) count() (int,error) {
	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
	i,err:=collcetion.Count()
	return i,err
}








