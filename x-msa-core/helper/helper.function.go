package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"time"

	"google.golang.org/grpc"
)

func init() {
	Ctx, CloseCtx = context.WithCancel(context.Background())
	C = make(chan os.Signal, 3)
}

func CreateConn(addr string) (conn *grpc.ClientConn, ctx context.Context, close context.CancelFunc, err error) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	c := make(chan bool)
	defer func() { c <- true }()

	Wg.Add(1)
	go func() {
		defer Wg.Done()
		for {
			select {
			case <-Ctx.Done():
				return
			case <-ticker.C:
				close()
			case <-c:
				return
			}
		}
	}()

	tick := 0
	for {
		select {
		case <-Ctx.Done():
			close()
			err = fmt.Errorf("context closed")
			return
		case <-ticker.C:
			if conn != nil {
				return
			} else if err != nil {
				if tick < 5 {
					tick++
					fmt.Println(err)
					continue
				} else {
					close()
					err = fmt.Errorf("cannot create conection %v: %v", addr, err)
					return
				}
			}
			ctx, close = context.WithCancel(Ctx)
			conn, err = grpc.DialContext(ctx, addr,
				grpc.WithInsecure(),
				grpc.WithBlock(),
			)
		}
	}
}

func ParseConfig(path string, out interface{}) error {
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf(KeyErrorNotFound+": file: %v", path)
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf(KeyErrorOpen+": file: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf(KeyErrorRead+": body: %v", err)
	}

	err = json.Unmarshal(buf, out)
	if err != nil {
		return fmt.Errorf(KeyErrorParse+": json: %v", err)
	}

	return nil
}

func Dispatch(f interface{}, args ...interface{}) error {
	ff := reflect.ValueOf(f)
	if ff.Kind() == reflect.Func {
		in := make([]reflect.Value, ff.Type().NumIn())
		for i, arg := range args {
			v := reflect.ValueOf(arg)
			if v.Type().ConvertibleTo(ff.Type().In(i)) {
				in[i] = v.Convert(ff.Type().In(i))
			} else {
				return fmt.Errorf("parameter: %v, expected %v got %v", i+1, ff.Type().In(i), v.Type())
			}
		}

		Wg.Add(1)
		go func() {
			ff.Call(in)
			Wg.Done()
		}()
	}
	return nil
}
