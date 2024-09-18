package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/woaijssss/godbx"
	"github.com/woaijssss/godbx/example/dal"
	"github.com/woaijssss/godbx/ttypes"
	txrequest "github.com/woaijssss/godbx/tx"
	"log"
	"time"
)

var datasource godbx.Datasource

func init() {
	conf := &godbx.DbConf{
		DbUrl:  "root:123456@tcp(localhost:3306)/godbx",
		LogSQL: true,
		Size:   1,
	}
	var err error
	datasource, err = godbx.NewDatasource(conf)
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	defer datasource.Shutdown()

	//testMapConn()
	//createUserUseAutoTrans()
	//createUser()
	//create()
	//query()
	//queryUser()
	queryUserPageForUpdate()
	//queryRawSQLForCount()
	//queryByIds()
	//queryByIdsUsingDao()
	//queryByMatcher()
	//queryAll()
	//queryByMatcherOrder()
	//countByMatcher()
	//update()

	//deleteById()
}

func queryUser() {
	tcCreate := func() (*godbx.TransContext, error) {
		return godbx.NewTransContext(datasource, txrequest.RequestNone, "trace-1001")
	}
	g, err := godbx.AutoTransWithResult(tcCreate, func(tc *godbx.TransContext) (*dal.UserInfo, error) {
		return godbx.GetById(tc, 1, dal.UserInfoMeta)
	})
	if err != nil {
		fmt.Println(err)
	}
	j, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("queryUser", string(j))
	var rg dal.UserInfo
	json.Unmarshal(j, &rg)
	fmt.Println(rg)
}

func queryUserPageForUpdate() {
	tcCreate := func() (*godbx.TransContext, error) {
		return godbx.NewTransContext(datasource, txrequest.RequestNone, "trace-1001")
	}
	mat := godbx.NewMatcher()
	mat.In(dal.UserInfoFields.Id, []any{1, 3, 5})
	pager := &godbx.Pager{
		PageNumber: 1,
		PageSize:   2,
	}
	viewColumns := []string{
		dal.UserInfoFields.Id,
		dal.UserInfoFields.Name,
		dal.UserInfoFields.CreateAt,
	}
	userInfos, err := godbx.AutoTransWithResult(tcCreate, func(tc *godbx.TransContext) ([]*dal.UserInfo, error) {
		return godbx.QueryPageListMatcherWithViewColumnsForUpdate(tc, mat, dal.UserInfoMeta, viewColumns, pager, true)
	})
	if err != nil {
		fmt.Println(err)
	}
	j, err := json.Marshal(userInfos)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("queryUser", string(j))

}

func createUserUseAutoTrans() {
	t := &dal.UserInfo{
		Name:     "godbx",
		CreateAt: ttypes.NormalDatetime(time.Now()),
		ModifyAt: *ttypes.FromDatetime(time.Now()),
	}
	affect, err := godbx.AutoTransWithResult[int64](func() (*godbx.TransContext, error) {
		return godbx.NewTransContext(datasource, txrequest.RequestWrite, "traceId-01")
	}, func(tc *godbx.TransContext) (int64, error) {
		return godbx.Insert(tc, t, dal.UserInfoMeta)
	})
	fmt.Println(affect, t.Id, err)
}
