package main

import (
	"errors"
	"fmt"
)

type QueryError struct {
	Query string
	Err   error
}

func (qe QueryError) Error() string {
	return qe.Query + qe.Err.Error()
}

func (e *QueryError) Unwrap() error { return e.Err }

var databaseErr = errors.New("connection err!")

func main() {
	if err := myQueryOperation(); err != nil {
		fmt.Printf("%v", err)
		err = errors.Unwrap(err)
		fmt.Printf("%v", err)
		err = errors.Unwrap(err)
		fmt.Printf("%v", err)
	}

}
func makeErr() error {
	// 制作一个错误，查询数据库错误
	return QueryError{"query table failed ", databaseErr}
}
func myQueryOperation() error {
	err := makeErr()
	if err != nil {
		fmt.Println(errors.Is(err, databaseErr), errors.As(err, &databaseErr), err.(QueryError).Err == databaseErr)
		// if e, ok := err.(QueryError); ok && e.Err == databaseErr {
		if errors.As(err, &databaseErr) {
			return fmt.Errorf("query failed because of database err : %w", err)
		}
	}
	return nil
}

// 总结：此特性重在优化错误包裹错误的情况，一般通过fmt.Errorf来包裹额外信息。然而这会时的新的err丢失内部的成员信息。
// 1.记住error是接口，一个错误可以是任何类型，仅需要实现Error方法,方便打印字符串
// 2.所以一个错误实际可以再其原始类型中存储很多额外信息的，比如构造结构体时添加更多的成员。
// 3.有时会通过包裹另一个错误的方式，来简单构造一个新错误。新错误只保留了字符串信息，但丢失了结构体成员信息。
// 4.所以使用errors.Unwrap可以把err展开一层。注意，必须是通过fmt.Errorf来包裹的且使用%w，才可以用此方法展开，比较时采用值比较Is。
// 5.如果是通过自定义的结构体进行包裹，则需自己实现Unwrap方法，比较时需要使用类型比较As
