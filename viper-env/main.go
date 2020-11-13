package main

import "github.com/spf13/viper"

func main() {
	//replacer := strings.NewReplacer(".", "_")
	//viper.SetEnvKeyReplacer(replacer)
	////viper.AddConfigPath("./")
	////viper.SetConfigType("yaml")
	//viper.AutomaticEnv()
	//
	//// If a config file is found, read it in.
	////if err := viper.ReadInConfig(); err != nil {
	////	fmt.Println(err)
	////	os.Exit(1)
	////}
	////fmt.Printf("Using config file:%s\n", viper.ConfigFileUsed())

	config := viper.New()
	config.AddConfigPath("/Users/snlan/nfs-data/baas/baas-artifacts/admin/chain1")
	config.SetConfigName("configtx")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
