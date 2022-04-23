/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2022-02-16 22:43:15
 * @LastEditors: MoonKnight
 * @LastEditTime: 2022-02-21 22:35:10
 */

package exporter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

func Exec(tempdesc *prometheus.Desc, dsn string, q string, ch chan<- prometheus.Metric) error {
	dsn = strings.TrimPrefix(dsn, "mysql://")
	list, err := PingDB(dsn, q)

	if err != nil {
		return err
	}

	querydata, _ := json.Marshal(list)
	dataString := string(querydata)
	ch <- prometheus.MustNewConstMetric(tempdesc, prometheus.GaugeValue, 1, dataString)
	return nil
}

// func PingDB(dsn string, query string) ([]map[string]string, error) { //数据库连接
// 	db, _ := sql.Open("mysql", dsn)

// 	err := db.Ping()
// 	if err != nil {
// 		fmt.Println("数据库链接失败")
// 	}
// 	defer db.Close()
// 	//多行查询
// 	rows, _ := db.Query(query)
// 	columns, _ := rows.Columns()
// 	columnLength := len(columns)
// 	cache := make([]interface{}, columnLength) //临时存储每行数据
// 	for index, _ := range cache {              //为每一列初始化一个指针
// 		var a interface{}
// 		cache[index] = &a
// 	}
// 	var list []map[string]string //返回的切片
// 	for rows.Next() {
// 		_ = rows.Scan(cache...)

// 		item := make(map[string]string)
// 		for i, data := range cache {
// 			//(data.(type))
// 			show := Strval(data)
// 			item[columns[i]] = Strval(data)
// 			log.Infof("strange %s", show)
// 			//item[columns[i]] = *data.(*interface{}) //取实际类型
// 		}
// 		list = append(list, item)
// 	}
// 	_ = rows.Close()
// 	return list, nil
// }

func PingDB(dsn string, query string) ([]map[string]string, error) { //数据库连接
	db, _ := sql.Open("mysql", dsn)

	err := db.Ping()
	if err != nil {
		fmt.Println("数据库链接失败")
	}
	defer db.Close()
	//多行查询
	rows, _ := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	results := make([]map[string]string, 0, 10)
	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]string, 10)
		for k, v := range values {
			key := cols[k]
			row[key] = string(v)
			log.Infof("strange %s", string(v))
		}
		results = append(results, row)
	}

	return results, nil
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
