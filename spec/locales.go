package spec

import (
	"dapperdox/config"
	"dapperdox/logger"
	"github.com/leonelquinteros/gotext"
	"io/ioutil"
	"os"
)

// localeTheme содержит словарь текущей темы
var localeTheme gotext.Po

// localeDefault содержит словарь дефолтной темы
var localeDefault gotext.Po

// LoadLocales загружает ресурсы локали
func LoadLocales() {
	const localeFileName = "locales.po"

	fileLocaleTheme := ""
	fileLocaleDefault := ""

	cfg, _ := config.Get()

	if len(cfg.AssetsDir) != 0 {
		fileLocaleTheme = cfg.AssetsDir + "/" + localeFileName
		logger.Tracef(nil, "Looking in assets dir for %s\n", fileLocaleTheme)
		if _, err := os.Stat(fileLocaleTheme); os.IsNotExist(err) {
			fileLocaleTheme = ""
		}
	}
	if len(fileLocaleTheme) == 0 && len(cfg.ThemeDir) != 0 {
		fileLocaleTheme = cfg.ThemeDir + "/" + cfg.Theme + "/" + localeFileName
		logger.Tracef(nil, "Looking in theme dir for %s\n", fileLocaleTheme)
		if _, err := os.Stat(fileLocaleTheme); os.IsNotExist(err) {
			fileLocaleTheme = ""
		}
	}
	if len(fileLocaleTheme) == 0 {
		fileLocaleTheme = cfg.DefaultAssetsDir + "/themes/" + cfg.Theme + "/" + localeFileName
		logger.Tracef(nil, "Looking in default theme dir for %s\n", fileLocaleTheme)
		if _, err := os.Stat(fileLocaleTheme); os.IsNotExist(err) {
			fileLocaleTheme = ""
		}
	}
	localeTheme = *loadLocale(fileLocaleTheme)

	if len(fileLocaleDefault) == 0 {
		fileLocaleDefault = cfg.DefaultAssetsDir + "/themes/default/" + localeFileName
		logger.Tracef(nil, "Looking in default theme %s\n", fileLocaleDefault)
		if _, err := os.Stat(fileLocaleDefault); os.IsNotExist(err) {
			fileLocaleDefault = ""
		}
	}
	localeDefault = *loadLocale(fileLocaleDefault)
}

// loadLocale загружает ресурс локали из файла
func loadLocale(localeFile string) *gotext.Po {
	translator := gotext.NewPo()

	if len(localeFile) == 0 {
		logger.Tracef(nil, "No localefile file found.")
		return translator
	}
	logger.Tracef(nil, "Processing locale file: %s\n", localeFile)
	file, err := os.Open(localeFile)

	if err != nil {
		logger.Errorf(nil, "Error: %s", err)
		return translator
	}
	defer file.Close()

	buf, err := ioutil.ReadFile(localeFile)
	if err != nil {
		logger.Errorf(nil, "Error: %s", err)
		return translator
	}

	translator.Parse(buf)

	return translator
}

// GetLocales возвращает локализованую строку
func GetLocales(str string) string {
	desc := localeTheme.Get(str)

	if desc != str {
		return desc
	}

	desc = localeDefault.Get(str)
	return desc
}
