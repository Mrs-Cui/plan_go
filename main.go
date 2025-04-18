package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"plan_go/tree"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	"unsafe"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmot"

	// "git.tigerbrokers.net/libs/grpc-service-go/ams/account_service/v1"
	// "git.tigerbrokers.net/libs/grpc-service-go/ams/user/v2"
	// "git.tigerbrokers.net/libs/grpc-service-go/ams/tesseract/v1"
	// "git.tigerbrokers.net/libs/grpc-service-go/ams/customer/v1"
	"git.tigerbrokers.net/libs/grpc-service-go/ams/virtual_account/v1"
	"google.golang.org/protobuf/encoding/protojson"
	// "google.golang.org/protobuf/types/known/fieldmaskpb"
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
	C []string
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

func IsIbAccount(aid string) bool {
	var flag = false
	for _, x := range aid {
		if unicode.IsDigit(x) {
			continue
		}
		flag = true
		break
	}
	return flag
}

func IsCycle(cycle int, ct time.Time) bool {
	mid := time.Date(ct.Year(), ct.Month(), ct.Day(), 0, 0, 0, 0, time.Local)
	now := time.Now()
	if cycle == 1 {
		return now.Sub(mid) < 24 * 3600 * time.Second
	}
	return true
}

var AlphanumericSet = []rune{
	// '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	// 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

const (
	PRIME1 = 7 // 与字符集长度 24 互质
	PRIME2 = 5 // 与邀请码长度 6 互质
	SALT   = 20230712	// 随意取一个数值
)

func GetInvCodeByUID(uid uint64, l int) string {
	// 放大 + 加盐
	uid = uid*PRIME1 + SALT

	var code []rune
	slIdx := make([]byte, l)

	// 扩散
	for i := 0; i < l; i++ {
		slIdx[i] = byte(uid % uint64(len(AlphanumericSet)))                   // 获取 24 进制的每一位值
		slIdx[i] = (slIdx[i] + byte(i)*slIdx[0]) % byte(len(AlphanumericSet)) // 其他位与个位加和再取余（让个位的变化影响到所有位）
		uid = uid / uint64(len(AlphanumericSet))                              // 相当于右移一位（24进制）
	}

	// 混淆
	for i := 0; i < l; i++ {
		idx := (byte(i) * PRIME2) % byte(l)
		code = append(code, AlphanumericSet[slIdx[idx]])
	}
	return string(code)
}

func GetManyDay(t time.Time, stime time.Time, d int) []time.Time {
	st := time.Date(stime.Year(), stime.Month(), stime.Day(), 0, 0, 0, 0, t.Location())
	et := st.AddDate(0, 0, d)
	for et.Sub(t).Seconds() < 0 {
		st = et
		et = et.AddDate(0, 0, d)
	}
	return []time.Time{st, et}
}

func check1 (w *sync.WaitGroup) {
	defer w.Done()
	fmt.Printf("hello world")
}

func check(w *sync.WaitGroup) {
	defer w.Done()
	panic("error")
}

func check2() {
	w := sync.WaitGroup{}
	w.Add(2)
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Printf("Err:%v", err)
		}	
	}()
	go check1(&w)
	go check(&w)
	w.Wait()

}

func CondCheck() {
	cond := sync.NewCond(&sync.Mutex{})
	var ready int
	var m sync.Mutex
	for i:= 0; i < 10; i ++ {
		go func (i int)  {
			time.Sleep(time.Duration(1) * time.Second)
			m.Lock()
			defer m.Unlock()
			ready += 1
			fmt.Printf("运动员: %d 已准备好\n", i)
			cond.Broadcast()
		}(i)
	}
	cond.L.Lock()
	for ready != 10 {
		cond.Wait()
		fmt.Printf("教练员被唤醒\n")
	}
	cond.L.Unlock()
	fmt.Printf("所有运动员准备完毕 1, 2, 3 ......")
}

func Server() {
	runtime.GOMAXPROCS(1) // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1) // 开启对阻塞操作的跟踪

	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	select{}
}

func ObjectPool() {
	pool := sync.Pool {
		New: func() any {
			return Field{}
		},
	}
	for i :=0; i <= 10000; i++ {
		obi := pool.Get()
		ob := obi.(Field)
		fmt.Printf("取出的对象:%v\n", ob)
		for j := 0; j <= i; j ++ {
			ob.C = append(ob.C, strconv.FormatInt(int64(j+10000000), 10))
		}
		size := unsafe.Sizeof(ob.C)
		pool.Put(ob)
		fmt.Printf("重新放入了对象，Size:%d\n", size)
		time.Sleep(2 * time.Second)
	}
}

type Field2 struct {
	*Field
}

type ListNode struct {
	Val int
	Next *ListNode
}
func oddEvenList(head *ListNode) *ListNode {
	odd := head
	even := head.Next
	for even != nil {
		mid := odd.Next
		odd.Next = even.Next
		even.Next = mid
		odd = odd.Next
		even = even.Next
	}
	return head
}

func Check() {
	num := 100
	var nodes []*ListNode
	for i :=0; i < num; i++ {
		nodes = append(nodes, &ListNode{Val:i})
	}
	wait := sync.WaitGroup{}
	wait.Add(len(nodes))
	for _, node := range nodes {
		nodeT := node
		go func(){
			defer wait.Done()
			fmt.Println("NodeTemp:%+v", nodeT)
		}()
	}
	wait.Wait()
}

func T( t map[string]bool) {
	for _, val := range []string{"A", "B", "C"} {
		t[val] = true
	}
}

func printContextInternals(ctx interface{}, inner bool) {
    contextValues := reflect.ValueOf(ctx).Elem()
    contextKeys := reflect.TypeOf(ctx).Elem()

	for contextKeys.Kind() == reflect.Pointer || contextKeys.Kind() == reflect.Ptr {
		contextKeys = contextKeys.Elem()
	}
    if !inner {
        fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
    }

    if contextKeys.Kind() == reflect.Struct {
        for i := 0; i < contextValues.NumField(); i++ {
            reflectValue := contextValues.Field(i)
            reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

            reflectField := contextKeys.Field(i)

            if reflectField.Name == "Context" {
                printContextInternals(reflectValue.Interface(), true)
            } else {
                fmt.Printf("field name: %+v\n", reflectField.Name)
                fmt.Printf("value: %+v\n", reflectValue.Interface())
            }
        }
    } else {
        fmt.Printf("context is empty (int)\n")
    }
}

func Trace() {
	opentracing.SetGlobalTracer(apmot.New())  
	tracer := opentracing.GlobalTracer()
 	parentSpan := tracer.StartSpan("parent", opentracing.Tag{Key: "hello", Value:"world"},)
	parentSpan.SetBaggageItem("Parent", "Parent")
	parentSpan.Finish()  
 	ctx := opentracing.ContextWithSpan(context.Background(), parentSpan)  
	sp := apm.SpanFromContext(ctx)
	fmt.Println("TraceId", sp.TraceContext())
	// 模拟一些工作  
	childSpan := tracer.StartSpan("child", opentracing.ChildOf(parentSpan.Context()))  
	childSpan.SetTag("key", "value") 
	childSpan = childSpan.SetBaggageItem("TTTT", "VVVV")
	childSpan.Finish()
	
	car :=opentracing.TextMapCarrier(map[string]string{})
	
	err := childSpan.Tracer().Inject(childSpan.Context(), opentracing.TextMap, car)
	fmt.Println("Err1:", err, car)

	car.Set("trace_id", "aaaaaaa")
	ctx = opentracing.ContextWithSpan(ctx, childSpan)

	spc, err := tracer.Extract(opentracing.TextMap, car)
	fmt.Println("Errttt:", err, car, spc)

	childSpan1 := opentracing.StartSpan("engine-2", opentracing.ChildOf(spc))
	ctx = opentracing.ContextWithSpan(ctx, childSpan1)

}

type User struct {
	Name string `json:"name,omitempty"`
	Age int `json:"age,omitempty"`
}

func main()  {
	// node := &tree.TriNode{}
	// strs := []string{"mpmyyz","czymyy","ycazdiog","ppcspz","mmyz","ymyo","ppczpapsuz","ppmmd","ad","mpo","uyzpzudupy","appppaz","czpss","mpdpgmdzyo","sdpui","zgyzii","z","ddazzpmcpg","mmdomszai","gpocoa","udsydmdymo","zmsp","dp","omdmdszud","pu","mpp","y","yyuppygu","zpzm","sd","p","psgsz","ypimc","pgsaydp","dy","aomuuc","opd","admzayo","ys","udzuzyyudc","s","dssyopdpi","po","ydy","pmgmsmm","ugzpipzpyy","sdoycusssm","pdgoi","mzu","acdazmps","adya","ocdm","izddagm","gzdudip","zmpmopu","zpyzgdmmmy","ydc","ymyydpsgm","pmpmo","zpzddays","capip","gpoczm","mpgzpmmpa","cgmg","ppicau","gsys","id","iiiooz","zyom","oyzyamdap","zdyypsdaua","o","iiu","oy","dg","ua","yddoa","gc","gdcd","mpmcd","cmyyys","m"}
	// node.CreateTri(strs)
	// fmt.Printf("%v\n", node)
	// node.AcPatternBuild()
	// fmt.Printf("%v\n", node)
	// rets := make([][]int, len(strs))
	// ret := node.Search("sdpuiycmyyysagpocoapsgszgdcdadmppsdgugzpipzpyyoyzyamdapsugzpipzpyympoudsydmdymodpsdyyuppyguu", rets, strs)
	// fmt.Printf("%v\n", ret)
	// xelasticsearch.NewClient()
	auLocal, _ := time.LoadLocation("Australia/Victoria")
	local, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Date(2023, 2, 9, 1, 0 ,0, 0, auLocal)
	// t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-12-07 13:01:01", local)
	k := t.In(local)
	fmt.Printf("%v", k)
	fmt.Printf("%v\n", IsCycle(1, t))
	a := " a b  c"
	fmt.Printf("%s\n", strings.ReplaceAll(a, " ", ""))
	f := 116.404844
	bits := math.Float64bits(f)
    bytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(bytes, bits)
	fmt.Printf("%v\n", bytes)
	location, _ := time.LoadLocation("Asia/Tomsk")
	kkk := time.Now().In(location).String()[:10]
	fmt.Printf("kkk:%v\n", kkk)
	context.Background()
	fmt.Printf("%v\n", time.Time{})
	fmt.Printf("%v\n", time.Unix(1678377600, 0))
	ctime := time.Now()
	mid := ctime.AddDate(0, 0, 1)
	tomorrow := time.Date(mid.Year(), mid.Month(), mid.Day(), 0, 0, 0, 0, time.Local)
	second := tomorrow.Sub(ctime).Seconds()
	fmt.Printf("second:%d\n", int64(second) / 3600)
	tt := time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local)
	date, _ := time.Parse("2006-01-02 15:04:05", "2023-03-05 10:00:00")
	if date.Sub(tt) < 0 {
		fmt.Printf("date:%v", date)
	}
	var word string = "hello word"
	for _, w := range []rune(word) {
		fmt.Printf("%s\n", string(w))
	}
	filed := new(Field)
	var cond sync.Cond
	cond.Signal()
	var s sync.Mutex
	s.Lock()
	fmt.Printf("%d\n", filed.A)
	fmt.Printf("%d\n", time.Now().UnixMilli())
	fmt.Printf("%v\n", time.Unix(1677834241843/ 1000, 0))
	escapedString := "\"Hello,\\nWorld!\""
	unescapedString, err := strconv.Unquote(escapedString)
	fmt.Printf("UnEscaped:%s, Err:%v\n", unescapedString, err)
	sum := md5.Sum([]byte("10000000320199"))
	fmt.Printf("%v\n", sum)
	fmt.Printf("%v\n", byte(62))
	fmt.Printf("%s\n", GetInvCodeByUID(10000000320196, 8))
	local1, _ := time.LoadLocation("Asia/Shanghai")
	date2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-06-28 08:52:39", local1)
	// stime, _ := time.Parse("2006-01-02 15:04:05", "2023-07-13 09:20:59")
	date1, _ := time.Parse("2006-01-02 15:04:05", "2023-07-05 06:53:03")
	fmt.Printf("Date1:%v Date2:%v\n", date1, date2.IsZero())
	L := date1.Sub(date2).Hours() / 24
	fmt.Printf("Hour:%f\n", L)
	timeLocal, _ := time.LoadLocation("Asia/Shanghai")
	t2 := time.Now().In(timeLocal)
	stime := time.Unix(1689326318, 0)
	etime := time.Unix(1691337599, 0)
	rtimes := GetManyDay(t2, stime, 7)
	if rtimes[1].Sub(etime).Seconds() > 0 {
		rtimes[1] = etime
	}
	fmt.Printf("%v\n", rtimes)
	fmt.Printf("%d\n", 3 &^ 2)
	tt1 := make(chan struct{}, 0)
	close(tt1)
	// xxx, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-11-10 10:38:07", time.Local)
	pattern :=regexp.MustCompile(`(?i)first`) 
	fmt.Println("Pattern:%v", pattern.Match([]byte("ttirstbbc")))
	fmt.Println("Time:%s", "2012-10-10"[:7])
	t6, _ := time.Parse("2006-01-02 15:04:05", "2024-04-07 06:53:03")
	weekDay := int(t6.Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	fmt.Println(t6, weekDay)
	dayFix := 1 // monday is the first day of week
	today := time.Date(t6.Year(), t6.Month(), t6.Day(), 0, 0, 0, 0, t6.Location())
	weekDay -= dayFix
	tiStart := today.AddDate(0, 0, -1*weekDay)
	tiEnd := tiStart.AddDate(0, 0, 7)
	fmt.Println("Start:%v End:%v", tiStart, tiEnd)

	fmt.Printf("%d", int('3' - '1'))
	
	fmt.Println("Time:%v", time.Unix(0, 0))

	
	// myProtoStruct := user.GetProfileRequest{
	// 	Identifier: &user.UserIdentifier{
	// 		Identifier: &user.UserIdentifier_UserId{
	// 			UserId: 10001,
	// 		},
	// 	},
	// 	FieldMask: &fieldmaskpb.FieldMask{
	// 		Paths: []string{"channel"},
	// 	},
	// }
	// myProtoStruct := user.GetProfileResponse{
	// 	UserId: 10001,
	// 	Uuid: 100000000000001,
	// 	Region: "CHN",
	// 	SignupAt: 1722355200,
	// 	Partition: 1,
	// 	AppCurrentAccountLicense: 1,
	// 	Status: 3,
	// 	Channel: &user.GetProfileResponse_RegistrationChannel{
	// 		AppName: "FACEBOOK",
	// 		RegSource: "CHANNL",
	// 	},
	// }
	// myProtoStruct := customer.GetUserProfileResponse{
	// 	UserProfile: &customer.UserProfile{
	// 		UserId: 10001,
	// 		IdNo: "IDNO",
	// 		HashIdNo: "HASHIDNO",
	// 	},
	// }
	myProtoStruct := virtual_account.GetAccountRequest{
		AccountId: 10000,
	}
	// myProtoStruct.FieldMask, _ = fieldmaskpb.New(new(user.GetProfileResponse), "channel", "transference")
	data, err := protojson.Marshal(myProtoStruct.ProtoReflect().Interface())
	if err != nil {
		log.Fatalf("Failed to JSON marhsal protobuf. Error: %s", err.Error())
	}
	fmt.Println("JSON: %s", string(data))
	next := tree.BuildNext("ababcabd")
	fmt.Println("Next:%v", next)
	pattern, err = regexp.Compile("[0-9.]+")
	if err != nil {
		return
	}
	m := pattern.Find([]byte("-0.013%"))
	fmt.Println("Match:%s", string(m))

	// 生成一个新的UUID
	newUUID := uuid.New()
 
	// 打印生成的UUID
	fmt.Println("Generated UUID:", newUUID)
 
	// 生成UUID的字符串表示形式
	uuidString := newUUID.String()
	fmt.Println("UUID as string:", uuidString)
	pool := sync.Map{}
	pool.Put()
}