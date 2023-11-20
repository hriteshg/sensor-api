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
		log.Errorf("error parsing limit query param %v", err)
		return 0, errors.New("query param: limit should be a float")
	}

	return l, err
}

func GetDate(c *gin.Context, key string, defaultVal *time.Time) (*time.Time, error) {
	dateStr, exists := c.GetQuery(key)
	if dateStr == "" || !exists {
		return defaultVal, nil
	}
	date, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		log.Errorf("error parsing %v date query param %v", key, err)
		return nil, errors.New(fmt.Sprintf("query param: %v should be a valid date in format: DD/MM/YYYY", key))
	}
	return &date, nil
}
