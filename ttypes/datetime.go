// A quickly mysql access component.

package ttypes

import (
	"bytes"
	"database/sql/driver"
	"github.com/woaijssss/godbx"
	"strings"
	"time"
)

var (
	// DatetimeFormat 指定时间格式
	DatetimeFormat = "2006-01-02 15:04:05"
)

// NormalDatetime 支持按日期格式输出的日期类型, 格式由 DatetimeFormat 全局变量指定, 实现fmt.Stringer, driver.Valuer, json.Unmarshaler, json.Marshaler 接口
type NormalDatetime time.Time

func ParseNormalDatetime(sDate string) (*NormalDatetime, error) {
	t, err := time.ParseInLocation(DatetimeFormat, sDate, time.Local)
	if err != nil {
		return nil, err
	}
	ndt := NormalDatetime(t)
	return &ndt, nil
}

func ParseNormalDatetimeInUTC(sDate string) (*NormalDatetime, error) {
	t, err := time.Parse(DatetimeFormat, sDate)
	if err != nil {
		return nil, err
	}
	ndt := NormalDatetime(t)
	return &ndt, nil
}

func ParseNormalDatetimeInLocation(sDate string, loc *time.Location) (*NormalDatetime, error) {
	t, err := time.ParseInLocation(DatetimeFormat, sDate, loc)
	if err != nil {
		return nil, err
	}
	ndt := NormalDatetime(t)
	return &ndt, nil
}

// Value 实现 driver.Valuer
func (ndt NormalDatetime) Value() (driver.Value, error) {
	return *ndt.ToTimePointer(), nil
}

// String 实现 fmt.Stringer 接口
func (ndt NormalDatetime) String() string {
	return ndt.ToTimePointer().Format(DatetimeFormat)
}

// UnmarshalJSON 实现 json.Unmarshaler
func (ndt *NormalDatetime) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || bytes.Compare(b, nullJsonValue) == 0 {
		return nil
	}

	value := strings.Trim(string(b), `"`)                             //get rid of "
	t, err := time.ParseInLocation(DatetimeFormat, value, time.Local) //parse time
	if err != nil {
		godbx.GLogger.SimpleLogError(err)
		return err
	}
	*ndt = NormalDatetime(t) //set result using the pointer
	return nil
}

// MarshalJSON 实现 json.Marshaler 接口
func (ndt NormalDatetime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ndt.ToTimePointer().Format(DatetimeFormat) + `"`), nil
}

func (ndt *NormalDatetime) ToTimePointer() *time.Time {
	return (*time.Time)(ndt)
}
