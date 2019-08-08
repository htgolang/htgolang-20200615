package users

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

type Persistence interface {
	Load() (map[int]User, error)
	Store(users map[int]User) (err error)
	removeFile() error
}

type JSONFile struct {
	name string
}

func NewJSONFile(name string) JSONFile {
	return JSONFile{name + ".json"}
}

func (j JSONFile) Load() (map[int]User, error) {
	bytes, err := ioutil.ReadFile(j.name)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]User), nil
		}
		return nil, err
	}
	var users map[int]User
	err = json.Unmarshal(bytes, &users)
	return users, err
}

func (j JSONFile) Store(users map[int]User) (err error) {
	bytes, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("[-]用户数据转化失败:", err)
	}
	_, err = os.Stat(j.name)
	if err != nil {
		if os.IsNotExist(err) {
			ioutil.WriteFile(j.name, bytes, os.ModePerm)
			return nil
		}
	} else {
		// fmt.Println(filepath.Join(filepath.Dir(j.name) + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Base(j.name)))
		if err = os.Rename(j.name, filepath.Join(filepath.Dir(j.name), strconv.FormatInt(time.Now().Unix(), 10)+"."+filepath.Base(j.name))); err != nil {
			return err
		}
		if err = j.removeFile(); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(j.name, bytes, os.ModePerm)
}

func (j JSONFile) removeFile() error {
	if names, err := filepath.Glob(filepath.Join(filepath.Dir(j.name), "*."+filepath.Base(j.name))); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		// fmt.Println(names)
		if len(names) > 3 {
			for _, name := range names[3:] {
				os.Remove(name)
			}
		}
	} else {
		return err
	}
	return nil
}

type GOBFile struct {
	name string
}

func NewGOBFile(name string) GOBFile {
	return GOBFile{name + ".gob"}
}

func (g GOBFile) Load() (map[int]User, error) {
	file, err := os.Open(g.name)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]User), nil
		}
		return nil, err
	}
	defer file.Close()
	var users map[int]User
	decode := gob.NewDecoder(file)
	decode.Decode(&users)
	return users, err
}

func (g GOBFile) Store(users map[int]User) (err error) {
	if _, err := os.Stat(g.name); err == nil {
		// fmt.Println(filepath.Join(filepath.Dir(j.name) + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Base(j.name)))
		if err = os.Rename(g.name, filepath.Join(filepath.Dir(g.name), strconv.FormatInt(time.Now().Unix(), 10)+"."+filepath.Base(g.name))); err != nil {
			return err
		}
		if err = g.removeFile(); err != nil {
			return err
		}
	}
	file, err := os.Create(g.name)
	if err != nil {
		return fmt.Errorf("[-]发生错误:", err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(users)
}

func (g GOBFile) removeFile() error {
	if names, err := filepath.Glob(filepath.Join(filepath.Dir(g.name), "*."+filepath.Base(g.name))); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		if len(names) > 3 {
			for _, name := range names[3:] {
				os.Remove(name)
			}
		}
	} else {
		return err
	}
	return nil
}

