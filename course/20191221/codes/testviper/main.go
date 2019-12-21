package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	reader := viper.New()
	reader.SetConfigName("config")
	reader.SetConfigType("yaml")
	reader.AddConfigPath(".")

	err := reader.ReadInConfig()
	fmt.Println(err)

	name := reader.GetString("name")
	fmt.Printf("%T, %v\n", name, name)

	age := reader.GetInt("age")
	fmt.Printf("%T, %v\n", age, age)


	sex := reader.GetBool("sex")
	fmt.Printf("%T, %v\n", sex, sex)

	hobby := reader.GetStringSlice("hobby")
	fmt.Printf("%T, %v\n", hobby, hobby)

	score := reader.GetStringMap("score")
	fmt.Printf("%T, %v\n", score, score)
}