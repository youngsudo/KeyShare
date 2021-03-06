package setting

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ćšć±ćé
var Conf = new(AppConfig)

type AppConfig struct {
	*App         `mapstructure:"app"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*ETHConfig   `mapstructure:"eth"`
	*EmailConfig `mapstructure:"email"`
}

type App struct {
	Name            string        `mapstructure:"name"`
	Mode            string        `mapstructure:"mode"`
	Port            int           `mapstructure:"port"`
	StartTime       string        `mapstructure:"start_time"`
	MachineID       int64         `mapstructure:"machine_id"`
	TokenExpiration time.Duration `mapstructure:"token_expiration"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

type ETHConfig struct {
	Ethclient    string `mapstructure:"ethclient"`
	Contract     string `mapstructure:"contract"`
	PrivateKey   string `mapstructure:"privatekey"`
	KeystoreName string `mapstructure:"keystore_name"`
	GiveETH      bool   `mapstructure:"give_eth"`
	GiveNum      int    `mapstructure:"give_num"`
}

type EmailConfig struct {
	FromEmail string `mapstructure:"from_email"`
	Password  string `mapstructure:"pass_word"`
	EmCilent  string `mapstructure:"em_client"`
	EmPort    int    `mapstructure:"em_port"`
}

func Init() error {
	// viper.SetConfigFile("./config.yaml") // æćźéçœźæä»¶è·ŻćŸ
	viper.SetConfigName("config") // éçœźæä»¶ćç§°(æ æ©ć±ć)
	viper.SetConfigType("yaml")   // ćŠæéçœźæä»¶çćç§°äž­æČĄææ©ć±ćïŒćéèŠéçœźæ­€éĄč
	// æ„æŸéçœźæä»¶æćšçè·ŻćŸ
	// çžćŻčè·ŻćŸ: æŻçžćŻčäșćŻæ§èĄæä»¶çè·ŻćŸ
	// ç»ćŻčè·ŻćŸ: æŻçžćŻčäșçł»ç»çæ čçźćœ
	viper.AddConfigPath("./")
	// ćșæŹäžæŻéćéçœźäž­ćżäœżçšç,ćèŻviperćœćæ°æźäœżçšä»äčæ ŒćŒć»è§Łæ
	// viper.SetConfigType("JSON")
	// viper.AutomaticEnv() // èȘćšèŻ»ćçŻćąćé

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// éçœźæä»¶æȘæŸć°éèŻŻïŒćŠæéèŠćŻä»„ćżœç„
			fmt.Println("éçœźæä»¶æȘæŸć°")
		} else {
			// éçœźæä»¶èą«æŸć°ïŒäœäș§çäșćŠć€çéèŻŻ
			fmt.Println("éçœźæä»¶èą«æŸć°ïŒäœäș§çäșćŠć€çéèŻŻ")
		}
		return err
	}

	// ć°éçœźæä»¶äž­çéçœźäżĄæŻććșććć° Conf ćéäž­
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("éçœźæä»¶è§Łæć€±èŽ„")
		return err
	}
	// çćŹéçœźæä»¶ćć
	viper.WatchConfig()
	// ćœéçœźæä»¶ćçććæ¶ïŒäŒè§ŠćOnConfigChangeæčæł
	viper.OnConfigChange(func(e fsnotify.Event) {
		// éçœźæä»¶ćçćæŽäčćäŒè°çšçćè°ćœæ°
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper Unmarshal error:", err)
		}
	})
	return nil
}
