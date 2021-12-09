/*
 * @Description: 
 * @version: 
 * @Author: MoonKnight
 * @Date: 2021-12-05 20:23:47
 * @LastEditors: MoonKnight
 * @LastEditTime: 2021-12-05 21:15:16
 */

package Yaobase_exporter

type Counter struct{
    count int64
}

func (c *Counter) Add(count int64) int64 {
    c.count += count
    return c.count
}
package main

import (
    "fmt"
    "net/http"
)


var counter  = new(Counter)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "lexporter_request_count{user=\"admin\"} %d", counter.Add(10))
}

func main () {
    http.HandleFunc("/metrics", HelloHandler)
    http.ListenAndServe(":8000", nil)
}