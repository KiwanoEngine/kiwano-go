package core

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	_db, err := leveldb.OpenFile("tmp", nil)
	db = _db
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type Mdata map[string]interface{}

func stringify(m Mdata) []byte {
	var enc_buf bytes.Buffer
	enc := gob.NewEncoder(&enc_buf)
	enc.Encode(m)

	return enc_buf.Bytes()
}

func parse(data []byte) Mdata {
	var dec_buf bytes.Buffer
	dec_buf.Write(data)

	dec := gob.NewDecoder(&dec_buf)

	var m Mdata
	dec.Decode(&m)
	return m
}

func Save(key string, m Mdata) {
	err := db.Put([]byte(key), stringify(m), nil)
	checkErr(err)
}

func Get(key string) Mdata {
	data, err := db.Get([]byte(key), nil)
	checkErr(err)
	return parse(data)
}

func GetAll() []Mdata {
	iter := db.NewIterator(nil, nil)
	marr := []Mdata{}
	for iter.Next() {
		key := string(iter.Key())
		value := parse(iter.Value())
		marr = append(marr, Mdata{
			key: value,
		})
	}
	iter.Release()
	return marr
}

func Del(key string) {
	err := db.Delete([]byte(key), nil)
	checkErr(err)
}

// func main() {
// 	// var buf bytes.Buffer
// 	// //fmt.Println(db)

// 	// enc := gob.NewEncoder(&buf)
// 	// //dec := gob.NewDecoder(&buf)

// 	// m1 := M1{
// 	// 	"a": 1,
// 	// }
// 	// err := enc.Encode(m1)
// 	// checkErr(err)
// 	// fmt.Println(buf.Bytes())

// 	// buf2 := bytes.NewBuffer(buf.Bytes()[:])
// 	// dec := gob.NewDecoder(buf2)

// 	// var m2 M1
// 	// dec.Decode(&m2)
// 	// fmt.Println(m2)
// 	//fmt.Println(123)
// 	SaveAsync("123", Mdata{
// 		"d1": 1234,
// 	})
// 	SaveAsync("1233", Mdata{
// 		"val": 41,
// 	})
// 	v := Get("123")
// 	fmt.Println(v)
// 	//Del("123")
// 	//v := GetAll()
// 	//fmt.Println(v)
// }
