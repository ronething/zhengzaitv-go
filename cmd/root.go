package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"

	//_ "net/http/pprof"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type NotifyConfig struct {
	Slack string
	WeCom string
}

var cfgFile string
var cookies string

//var globalClient = resty.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zhengzaitv-go",
	Short: "正在现场命令行样例",
	Long:  `正在现场命令行样例`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
	cobra.OnInitialize(initConfig) // todo: 其实这里也可以不要 自己进行 config 读取也可以的
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.zhengzaitv.yaml)",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//log.Println(home)
		// todo: 可以改为当前目录
		viper.AddConfigPath(home)
		viper.SetConfigName(".zhengzaitv") // 不需要后缀
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("使用的配置文件:", viper.ConfigFileUsed())
		cookies = viper.GetString("cookies")
		log.Printf("cookies is %+v\n", cookies)
	}

}
