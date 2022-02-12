package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath("configs/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	fmt.Println("vp:", vp)

	err := vp.ReadInConfig()

	fmt.Println("setting:", err)
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
