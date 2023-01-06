package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tianmarillio/channela-backend/config"
	"github.com/tianmarillio/channela-backend/src/api/youtubeapi"
	"github.com/tianmarillio/channela-backend/src/models"
)

type channelController struct{}

var ChannelController channelController

func (ctr channelController) Create(c *gin.Context) {
	// get user from auth
	var reqUser, _ = c.Get("user")
	user := reqUser.(*models.User)

	// extract body & validate
	var body struct {
		YoutubeChannelId string `json:"youtubeChannelId"`
		Description      string `json:"description"`
	}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// find channel on youtube api
	foundChannel := youtubeapi.FindChannel(body.YoutubeChannelId)

	// create db row
	channel := models.Channel{
		ID:               uuid.NewString(),
		CreatedBy:        user.ID,
		YoutubeChannelId: body.YoutubeChannelId,
		Description:      body.Description,
		Title:            foundChannel.Title,
		CustomUrl:        foundChannel.CustomUrl,
	}
	result := config.DB.Create(&channel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	// return json
	c.JSON(http.StatusCreated, gin.H{
		"channelId": channel.ID,
	})
}

func (ctr channelController) FindAll(c *gin.Context) {
	// query db
	var channels []*models.Channel
	result := config.DB.Find(&channels)
	// TODO: join user
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	// return json
	c.JSON(http.StatusOK, gin.H{
		"channels": channels,
	})
}

func (ctr channelController) FindById(c *gin.Context) {
	// extract params & validate
	channelId := c.Param("channelId")

	// query db
	var channel *models.Channel
	result := config.DB.Where("id = ?", channelId).Find(&channel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// return json
	c.JSON(http.StatusOK, gin.H{
		"channel": channel,
	})
}

func (ctr channelController) SearchYoutubeChannel(c *gin.Context) {
	// extract query parameter q
	keyword, _ := c.GetQuery("q")

	// call youtube api search channel
	channels := youtubeapi.SearchChannel(keyword)

	// return channels search result
	c.JSON(http.StatusOK, gin.H{
		"channels": channels,
	})
}

func (ctr channelController) Update(c *gin.Context) {
	// get user from auth
	var reqUser, _ = c.Get("user")
	user := reqUser.(*models.User)

	// extract params & validate
	channelId := c.Param("channelId")

	// extract body & validate
	var body struct {
		Description string `json:"description"`
	}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// update db row
	channel := models.Channel{
		UpdatedBy:   user.ID,
		Description: body.Description,
	}
	result := config.DB.
		Where("id = ?", channelId).
		Updates(&channel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// return json
	c.JSON(http.StatusAccepted, gin.H{
		"channelId": channelId,
	})
}

func (ctr channelController) Delete(c *gin.Context) {
	// extract params & validate
	channelId := c.Param("channelId")
	fmt.Println("channelId>", channelId)

	// delete db row
	var channel *models.Channel
	result := config.DB.
		Where("id = ?", channelId).
		Delete(&channel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// return json
	c.JSON(http.StatusAccepted, gin.H{
		"channelId": channelId,
	})
}
