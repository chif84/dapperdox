package spec

import (
	config "dapperdox/config"
	"dapperdox/logger"
	"github.com/leonelquinteros/gotext"
	"io/ioutil"
	"os"
)

type localeType map[string]gotext.Po

// localeTheme содержит словарь текущей темы
var localeTheme localeType

// localeDefault содержит словарь дефолтной темы
var localeDefault localeType

// LoadLocales загружает ресурсы локали
func LoadLocales() {
	cfg, _ := config.Get()

	localeTheme = make(localeType)
	localeDefault = make(localeType)

	loadLocaleForDomain(cfg.Locale, cfg.DefaultAssetsDir, cfg.AssetsDir, cfg.ThemeDir, cfg.Theme)

	if config.DefaultLocale != cfg.Locale {
		loadLocaleForDomain(config.DefaultLocale, cfg.DefaultAssetsDir, cfg.AssetsDir, cfg.ThemeDir, cfg.Theme)
	}
}

func loadLocaleForDomain(domain, defaultAssetsDir, assetsDir, themeDir, theme string) {
	const localeFolderName = "locales"
	const localeFileName = "locales.po"

	fileLocaleTheme := ""
	fileLocaleDefault := ""

	if len(assetsDir) != 0 {
		fileLocaleTheme = assetsDir + "/" + localeFolderName + "/" + domain + "/" + localeFileName
		logger.Tracef(nil, "Looking in assets dir for %s\n", fileLocaleTheme)
		if _, err := os.Stat(fileLocaleTheme); os.IsNotExist(err) {
			fileLocaleTheme = ""
		}
	}
	if len(fileLocaleTheme) == 0 && len(themeDir) != 0 {
		fileLocaleTheme = themeDir + "/" + theme + "/" + localeFolderName + "/" + domain + "/" + localeFileName
		logger.Tracef(nil, "Looking in theme dir for %s\n", fileLocaleTheme)
		if _, err := os.Stat(fileLocaleTheme); os.IsNotExist(err) {
			fileLocaleTheme = ""
		}
	}
	if len(fileLocaleTheme) == 0 {
		fileLocaleTheme = defaultAssetsDir + "/themes/" + theme + "/" + localeFolderName + "/" + domain + "/" + localeFileName
		logger.Tracef(nil, "Looking in default theme dir for %s\n", fileLocaleTheme)
		if _, err := os.Stat(fileLocaleTheme); os.IsNotExist(err) {
			fileLocaleTheme = ""
		}
	}
	localeTheme[domain] = *loadLocale(fileLocaleTheme)

	if len(fileLocaleDefault) == 0 {
		fileLocaleDefault = defaultAssetsDir + "/themes/default/" + localeFolderName + "/" + domain + "/" + localeFileName
		logger.Tracef(nil, "Looking in default theme %s\n", fileLocaleDefault)
		if _, err := os.Stat(fileLocaleDefault); os.IsNotExist(err) {
			fileLocaleDefault = ""
		}
	}
	localeDefault[domain] = *loadLocale(fileLocaleDefault)
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

	cfg, _ := config.Get()

	if po, ok := localeTheme[cfg.Locale]; ok {
		desc := po.Get(str)
		if desc != str {
			return desc
		}
	}

	if po, ok := localeDefault[cfg.Locale]; ok {
		desc := po.Get(str)
		if desc != str {
			return desc
		}
	}

	if po, ok := localeTheme[config.DefaultLocale]; ok {
		desc := po.Get(str)
		if desc != str {
			return desc
		}
	}

	if po, ok := localeDefault[config.DefaultLocale]; ok {
		desc := po.Get(str)
		if desc != str {
			return desc
		}
	}

	return str
}
