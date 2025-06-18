package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	config               *Config
	ErrUserNotFoundError = fmt.Errorf("user not found in configuration")
)

type Config struct {
	// Add your configuration fields here
	AppUserList []AppUser `json:"app_users"`
}

type AppUser struct {
	AppName    string   `json:"app_name"`
	User       string   `json:"user"`
	PassHash   string   `json:"pass_hash"`
	CreaatedAt string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	Tags       []string `json:"tags"`
}

func init() {
	config, _ = Load()
}

func GetConfig() *Config {
	if config == nil {
		fmt.Fprintf(os.Stderr, "Configuration is not loaded\n")
		os.Exit(1)
	}
	return config
}

func (c *Config) AddUpdateAppUser(appName, user, passHash string, tags []string) {
	if c.AppUserList == nil {
		c.AppUserList = []AppUser{}
	}

	for i, appUser := range c.AppUserList {
		if appUser.AppName == appName && appUser.User == user {
			c.AppUserList[i].PassHash = passHash
			c.AppUserList[i].UpdatedAt = time.Now().Format(time.RFC3339)
			c.AppUserList[i].Tags = tags
			return
		}
	}

	appUser := AppUser{
		AppName:    appName,
		User:       user,
		PassHash:   passHash,
		CreaatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
		Tags:       tags,
	}

	c.AppUserList = append(c.AppUserList, appUser)
}

func (c *Config) FindAppUser(appName, user string) (*AppUser, error) {
	if c.AppUserList == nil {
		return nil, fmt.Errorf("no app users found")
	}

	for _, appUser := range c.AppUserList {
		if appUser.AppName == appName && appUser.User == user {
			return &appUser, nil
		}
	}

	return nil, ErrUserNotFoundError
}

func Load() (*Config, error) {
	config := &Config{}

	// Load configuration from a file or environment variables
	// For now, we will just return an empty config
	// In a real application, you would read from a file or environment variables
	userHome, _ := os.UserHomeDir()
	if userHome == "" {
		return nil, fmt.Errorf("HOME environment variable is not set")
	}
	configFilePath := userHome + "/.box/config.json"
	if _, err := os.Stat(userHome + "/.box"); os.IsNotExist(err) {
		os.MkdirAll(userHome+"/.box", 0755)
	}
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		file, err := os.Create(configFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create config file: %v", err)
		}
		data, err := json.Marshal(&Config{})
		file.Write(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal empty config: %v", err)
		}
		file.Close()
		time.Sleep(1 * time.Second) // Ensure the file is created before we try to read it
	}

	configFile, _ := os.Open(configFilePath)
	decoder := json.NewDecoder(configFile)

	err := decoder.Decode(config)

	if err != nil && err != io.EOF {
		if config.AppUserList == nil {
			config.AppUserList = []AppUser{} // Initialize with an empty list if decoding fails
		}
	}

	return config, err
}

func (c *Config) Save() error {
	userHome, _ := os.UserHomeDir()
	if userHome == "" {
		return fmt.Errorf("HOME environment variable is not set")
	}
	configFilePath := userHome + "/.box/config.json"

	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(*c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}
	file.Write(jsonData)
	fmt.Println("Configuration saved successfully to", configFilePath)

	return nil
}
