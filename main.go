package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"plan_go/tree"
)


var V int

func F() { fmt.Printf("Hello, number %d\n", V) }

type db struct {
	H string
}

func (m db) Test() {

}

type DB *db

type D = int
type I int

type Field struct {
	A int
	B []int64
}

func Hello(a *Field) {
	a.B = append(a.B, 4)
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

//func main() {
	// 第一个[{}] 表示参数解释，第二个表示返回值解释
	//var str string
	//str = "FirstDeposit 是否首次入金  [\"入金金额\"];[\"depostNum\", \"入金金额\"]\n"
	//str = "FirstDeposit 是否首次入金\n"
	//str = "// FirstDeposit "
	//str = "# // % FirstDeposit 入金 [];[\"depostNum\", \"入金金额\"];"
	//str = "# // % FirstDeposit [{}]"
	//pattern := "[\\s/a-zA-Z]*(?P<desc>[\u4e00-\u9fa5\\w]*)\\s*(?P<params>[\"\u4e00-\u9fa5\\w ,\\[\\]]*);*\\s*(?P<renVal>[\"\u4e00-\u9fa5\\w ,\\[\\]]*);*"
	//re, err := regexp.Compile(pattern)
	//if err != nil {
	//	fmt.Printf("ERR:%v\n", err)
	//	return
	//}
	//rest :=re.FindStringSubmatch(str)
	//names := re.SubexpNames()
	//for _, name := range names {
	//	if name == "" {
	//		continue
	//	}
	//	idx := re.SubexpIndex(name)
	//	fmt.Printf("name:%s ret:%s\n", name, rest[idx])
	//}
	//pattern = "^[\\w\\d_]+$"
	//t, _ := regexp.Compile(pattern)
	//ret := t.Find([]byte("aab_t"))
	//fmt.Printf(string(ret))

	//re, _ := regexp.Compile("(?<=[\\D]*)\\d+")
	//t := re.Find([]byte("abd138784"))
	//str := string(t)
	//fmt.Printf("%s", str)
//}

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	go handle(ctx, 500*time.Millisecond)
//	select {
//	case v:= <-ctx.Done():
//		fmt.Println("main", ctx.Err(), v)
//	}
//}
//
//func handle(ctx context.Context, duration time.Duration) {
//	select {
//	case <-ctx.Done():
//		fmt.Println("handle", ctx.Err())
//	case <-time.After(duration):
//		fmt.Println("process request with", duration)
//	}
//}


//func main() {
//	cli, err := xzkp.CreateClient([]string{"127.0.0.1"}, time.Second * 5)
//	if err != nil {
//		fmt.Printf("err:%v", err)
//	}
//	path := "/hello"
//	ev, err := cli.WatchExists(path)
//	paths, err := cli.GetChild(path)
//	if err != nil {
//		return
//	}
//	fmt.Printf("paths:%v ev:%+v", paths, ev)
//	fmt.Printf("cli:%v", cli)
//}

func main()  {
	tree.CheckRedBlackTree()
}