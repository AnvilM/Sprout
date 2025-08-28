package base

import (
	"foundryinstaller/internal/embedassets"
	"foundryinstaller/internal/utils/logger"
)

func ExportConfigs(){
	if err := embedassets.ExtractAssets(); err != nil {
		logger.Fatal("%s", err.Error())
	}

	logger.Log(logger.SUCCESS, "Configuration successfully exported!")
}