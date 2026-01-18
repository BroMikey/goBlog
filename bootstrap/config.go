package bootstrap

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`  // MySQL配置
	Logger Logger `yaml:"logger"` // 日志配置
	System System `yaml:"system"` // 系统配置
}

type Mysql struct {
	Host        string `yaml:"host"`          // 数据库主机
	Port        int    `yaml:"port"`          // 数据库端口
	Config      string `yaml:"config"`        // 额外配置
	DBname      string `yaml:"dbname"`        // 数据库名称
	Username    string `yaml:"username"`      // 数据库用户名
	Password    string `yaml:"password"`      // 数据库密码
	MaxIdleConn int    `yaml:"max_idle_conn"` // 最大空闲连接数
	MaxOpenConn int    `yaml:"max_open_conn"` // 最大连接数
	LogLevel    string `yaml:"log_level"`     // 日志级别
}

type Logger struct {
	Level        string `yaml:"level"`          // 日志级别
	Prefix       string `yaml:"prefix"`         // 日志前缀
	Director     string `yaml:"director"`       // 日志文件目录
	ShowLine     bool   `yaml:"show_line"`      // 是否显示行号
	LogInConsole bool   `yaml:"log_in_console"` // 是否在控制台输出日志
}

type System struct {
	Host string `yaml:"host"` // 系统主机
	Port int    `yaml:"port"` // 系统端口
	Env  string `yaml:"env"`  // 运行环境
}

func LoadConfig(path string) (*Config, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %w", err)
	}

	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("unmarshal config file failed: %w", err)
	}

	return &c, nil
}

// database source name
func (m Mysql) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DBname,
		m.Config,
	)
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
