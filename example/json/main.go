package main

import (
	"encoding/json"
	"fmt"
	"github.com/woaijssss/godbx/ttypes"
	"time"
)

type Data struct {
	Id       int64
	Name     string
	CreateAt ttypes.NormalDatetime
	Modify   ttypes.NilableDatetime
	S        ttypes.NilableString
	Str      string
}

func main() {
	d := &Data{
		Id:       2,
		Name:     "godbx",
		CreateAt: ttypes.NormalDatetime(time.Now()),
		Modify:   *ttypes.FromDatetime(time.Now()),
		S:        *ttypes.FromString(`abc`),
		Str:      `gr`,
	}
	j, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(j))
	var t Data
	json.Unmarshal(j, &t)

	s := t.S.String
	fmt.Println(s)
}
