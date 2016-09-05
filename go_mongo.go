package main
import(
    "fmt"
    // "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Pperson struct{
    Name string
    Phone string
}

type Men struct{
    Persons []Pperson
}

const URL="127.0.0.1:27017"


func main(){
    session,err:=mgo.Dial(URL) //连接数据库
    if err!=nil{
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic,true)

    db:=session.DB("cyy_test") //数据库名称
    collection:=db.C("person") //如果该集合已经存在的话，直接返回

    //****集合中元素数目****
    countNum,err:=collection.Count()
    if err!=nil{
        panic(err)
    }
    fmt.Println("Things objects count: ", countNum)

    //****插入元素****
    temp := &Pperson{
        Name: "cyy",
        Phone: "wewq2423",
    }

    //一次可以插入多个对象，插入两个Pperson对象
    err = collection.Insert(&Pperson{"xhc","2343252353"},temp)
    if err!=nil{
        panic(err)
    }

    //****查询单条数据****
    result:=Pperson{}
    err=collection.Find(bson.M{"phone":"456"}).One(&result)
    fmt.Println("Phone:",result.Name,result.Phone)

    //*****查询多条数据*****
    var personAll Men //存放结果
    iter:=collection.Find(nil).Iter()
    for iter.Next(&result){
        fmt.Printf("Result: %v\n", result.Name)
        personAll.Persons = append(personAll.Persons, result)
    }

    //*******更新数据**********
  err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
  err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
  err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

  //******删除数据************
    _, err = collection.RemoveAll(bson.M{"name": "Ale"})
}
