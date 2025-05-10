package command

import (
	"fmt"
	"github.com/meth-suchatchai/kz-blog-api/lib/kzcrypto"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func GenerateSecretKey(context *cli.Context) error {
	defer context.Done()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if viper.IsSet("jwt.secret") {
		fmt.Printf("key already exists. Overwrite it? [y/N]")
		var input string
		fmt.Scanln(&input)
		if input != "y" && input != "Y" {
			fmt.Println("Keeping existing key.")
			return nil
		}
	}

	newSecret, err := kzcrypto.GenerateRandomString(32)
	if err != nil {
		return err
	}

	viper.Set("jwt.secret", newSecret)
	err = viper.WriteConfig()
	if err != nil {
		err = viper.SafeWriteConfig()
		if err != nil {
			return fmt.Errorf("failed to write config: %w", err)
		}
	}

	fmt.Println("Web Secret key updated.")
	return nil
}
