package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func GetFloatQueryParam(ctx *gin.Context, param string) (float64, error) {
	v, exists := ctx.GetQuery(param)
	if !exists {
		return 10, nil
	}
	l, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Errorf("error parsing query param %v", err)
		return 0, errors.New("query param should be a float")
	}

	return l, err
}

func GetTime(c *gin.Context, key string, defaultVal *time.Time) (*time.Time, error) {
	unixTimeStr, exists := c.GetQuery(key)
	if unixTimeStr == "" || !exists {
		return defaultVal, nil
	}
	unixTime, err := strconv.ParseInt(unixTimeStr, 10, 64)
	log.Infof("Unix time %v", unixTime)
	t := time.Unix(unixTime, 0)
	if err != nil {
		log.Errorf("error parsing %v timestamp query param %v", key, err)
		return nil, errors.New(fmt.Sprintf("query param: %v should be a valid timestamp", key))
	}
	return &t, nil
}
