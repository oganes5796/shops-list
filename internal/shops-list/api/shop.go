package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/shops-list/model"
	servShop "github.com/oganes5796/shops-list/internal/shops-list/service/shop"
)

func (im *Implementation) Create(c *gin.Context) {
	var info *model.ShopInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "internal error",
		})
		return
	}

	id, err := im.services.ShopService.Create(c.Request.Context(), info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "api:Create:Service.Create",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (im *Implementation) GetAll(c *gin.Context) {
	lists, err := im.services.ShopService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (im *Implementation) GetByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	shop, err := im.services.ShopService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, servShop.ErrShopNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "shop not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, shop)
}

func (im *Implementation) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var info model.ShopInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err = im.services.ShopService.Update(c.Request.Context(), id, &info)
	if err != nil {
		if errors.Is(err, servShop.ErrShopNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (im *Implementation) Delete(c *gin.Context) {
	idUrl := c.Param("id")
	id, err := strconv.ParseInt(idUrl, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "internal error",
		})
		return
	}

	err = im.services.ShopService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, servShop.ErrShopNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "shop not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
