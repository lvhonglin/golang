package main

import "github.com/spf13/viper"

func main() {
	viper.SetDefault("test", "xxx")
	println(viper.GetString("test"))
	v := viper.New()
	v.SetDefault("test", "yyy")
	println(v.GetString("test"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("com/lucy/test/study/viper/config/")
	err := viper.ReadInConfig()
	if err != nil {
		println(err.Error())
		return
	}
	println(viper.GetString("pass"))
	println(viper.GetString("a.b"))
}
