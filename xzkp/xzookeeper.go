package xzkp

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)


type ZkWatcher struct {}

func (m *ZkWatcher) Watch1(event zk.Event) {
	fmt.Println("###########################")
	fmt.Println("path: ", event.Path)
	fmt.Println("type: ", event.Type.String())
	fmt.Println("state: ", event.State.String())
	fmt.Println("---------------------------")
}

type ZkClientInter interface {
	Add(path string, data string) (string, error)
	Get(path string) (string, error)
	Modify(path string, data string) error
	Del(path string) error
	WatchExists(path string) (<-chan zk.Event, error)
	GetChild(path string) ([]string, error)
}

type zookeeperClient struct {
	client *zk.Conn
	ec    <- chan zk.Event
	Watch  *ZkWatcher
}

func (m *zookeeperClient) Add(path string, data string) (string, error) {
	acls := zk.WorldACL(zk.PermAll)
	var flags int32 = 0
	s, err := m.client.Create(path, []byte(data), flags, acls)
	if err != nil {
		return "", err
	}
	return s, nil
}

func (m *zookeeperClient) Get(path string) (string, error) {
	data, _, err := m.client.Get(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *zookeeperClient) Modify(path string, data string) error {
	_, stat, _ := m.client.Get(path)
	_, err := m.client.Set(path, []byte(data), stat.Version)
	if err != nil {
		return err
	}
	return nil
}

func(m *zookeeperClient) Del(path string) error {
	_, stat, _ := m.client.Get(path)
	err := m.client.Delete(path, stat.Version)
	if err != nil {
		return err
	}
	return nil
}

func (m *zookeeperClient) GetChild(path string) ([]string, error) {
	paths, _, err := m.client.Children(path)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func (m *zookeeperClient) WatchExists(path string) (<-chan zk.Event, error) {
	_,_,ev, err := m.client.ExistsW(path)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

func (m *zookeeperClient) GetExists(path string) error {
	_, _, _, err := m.client.GetW(path)
	if err != nil {
		return err
	}
	return nil
}


func CreateClient(servers []string, timeout time.Duration) (ZkClientInter, error) {
	cli := &zookeeperClient{
		Watch: &ZkWatcher{},
	}
	callbackOption := zk.WithEventCallback(cli.Watch.Watch1)
	var err error
	cli.client, cli.ec, err = zk.Connect(servers, timeout, callbackOption)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

