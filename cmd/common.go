package cmd

import (
	"fmt"
	"log"
	"zhengzaitv-go/wap"

	"github.com/spf13/cobra"
)

func getFlag(cmd *cobra.Command, key string) string {
	value, err := cmd.Flags().GetString(key)
	if err != nil || value == "" {
		log.Fatalf("获取 %s 发生错误,请检查输入是否正确, err: %+v\n", key, err)
		return ""
	}

	log.Printf("%v 值为: %v", key, value)
	return value
}

func getFlagInt(cmd *cobra.Command, key string) int {
	value, err := cmd.Flags().GetInt(key)
	if err != nil {
		log.Fatalf("获取 %s 发生错误,请检查输入是否正确, err: %+v\n", key, err)
		return 0
	}

	log.Printf("%v 值为: %v", key, value)
	return value
}

func PreRunZhengzaitv(cmd *cobra.Command, args []string) error {
	//参数校验
	if cookies == "" {
		return fmt.Errorf("配置为空")
	}

	zhengzai = wap.NewZhengzaitv(cookies)
	err := zhengzai.InjectUserInfo()
	if err != nil {
		log.Printf("inject user err: %v\n", err)
		return err
	}

	return nil
}
