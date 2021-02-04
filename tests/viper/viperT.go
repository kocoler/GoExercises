package main

import (
	"fmt"
	"github.com/kocoler/GoExercises/tests/viper/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	var cfg = pflag.StringP("config", "c", "", "apiserver config file path.")

	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}


	// Set gin mode.
	fmt.Println(viper.GetString("runmode"))

}
