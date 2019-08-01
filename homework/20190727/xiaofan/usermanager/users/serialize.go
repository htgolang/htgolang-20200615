package users

import (
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

const (
	jsonfile = "user.json"
	gobfile  = "user.gob"
	csvfile  = "user.csv"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
	Desc     string
}

// 用户序列化的接口
type UserSerial interface {
	Load() map[int]User
	Store(map[int]User)
}

// 序列化为JSON格式
type JsonFile map[int]User

// 序列化为GOB格式
type GobFile map[int]User

// 序列化为CSV格式
type CsvFile map[int]User

// 从user.json文件反序列化
func (m JsonFile) Load() map[int]User {
	if file, err := ioutil.ReadFile(jsonfile); err == nil {
		_ = json.Unmarshal(file, &m)
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("[-]发生错误: ", err)
		}
	}
	return m
}

// 序列化为user.json文件
func (m JsonFile) Store(users map[int]User) {
	if _, err := os.Stat(jsonfile); err == nil {
		_ = os.Rename(jsonfile, strconv.FormatInt(time.Now().Unix(), 10)+"."+jsonfile)
	}

	if names, err := filepath.Glob("*." + jsonfile); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		fmt.Println(names)
		if len(names) > 3 {
			for _, name := range names[3:] {
				_ = os.Remove(name)
			}
		}
	}

	marshal, _ := json.Marshal(users)
	if err := ioutil.WriteFile(jsonfile, marshal, os.ModePerm); err != nil {
		fmt.Println("Store Json Failed")
	}
}

// 从user.csv文件反序列化
func (m CsvFile) Load() map[int]User {
	if file, err := os.Open(csvfile); err == nil {
		defer file.Close()
		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println("[-]发生错误:", err)
				}
				break
			}
			id, _ := strconv.Atoi(line[0])
			birthday, _ := time.Parse("2006-01-02", line[2])

			m[id] = User{
				ID:       id,
				Name:     line[1],
				Birthday: birthday,
				Tel:      line[3],
				Addr:     line[4],
				Desc:     line[5],
			}
		}
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("[-]发生错误: ", err)
		}
	}
	return m
}

// 反序列化为user.csv文件
func (m CsvFile) Store(users map[int]User) {
	if _, err := os.Stat(csvfile); err == nil {
		_ = os.Rename(csvfile, strconv.FormatInt(time.Now().Unix(), 10)+"."+csvfile)
	}

	if names, err := filepath.Glob("*." + csvfile); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		fmt.Println(names)
		if len(names) > 3 {
			for _, name := range names[3:] {
				_ = os.Remove(name)
			}
		}
	}

	if file, err := os.Create(csvfile); err == nil {
		defer file.Close()
		writer := csv.NewWriter(file)
		for _, user := range users {
			_ = writer.Write([]string{
				strconv.Itoa(user.ID),
				user.Name,
				user.Birthday.Format("2006-01-02"),
				user.Tel,
				user.Addr,
				user.Desc,
			})
		}
		writer.Flush()
	}
}

// 从user.gob文件反序列化
func (m GobFile) Load() map[int]User {
	if file, err := os.Open(gobfile); err == nil {
		defer file.Close()
		decoder := gob.NewDecoder(file)
		_ = decoder.Decode(&m)
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("[-]发生错误: ", err)
		}
	}
	return m
}

// 序列化为user.gob文件
func (m GobFile) Store(users map[int]User) {
	if _, err := os.Stat(gobfile); err == nil {
		_ = os.Rename(gobfile, strconv.FormatInt(time.Now().Unix(), 10)+"."+gobfile)
	}

	if names, err := filepath.Glob("*." + gobfile); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		fmt.Println(names)
		if len(names) > 3 {
			for _, name := range names[3:] {
				_ = os.Remove(name)
			}
		}
	}

	if file, err := os.Create(gobfile); err == nil {
		defer file.Close()
		encoder := gob.NewEncoder(file)
		_ = encoder.Encode(users)
	}
}
