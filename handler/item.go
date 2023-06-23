package handler

import (
	"fmt"
	"net/http"

	"ntika/auth"
	"ntika/item"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/golang-jwt/jwt/v5"
)

type handlerTag struct {
	itemService item.Service
	userService auth.Service
}

func NewHandlerItems(itemService item.Service, songService auth.Service) *handlerTag {
	return &handlerTag{itemService, songService}
}

// // ================================================
// QUERY PARAMS OK
// TODO :
// IMPLEMENT FILTER & SORT
func (h *handlerTag) Catalog(c *gin.Context) {
	filter := c.Query("filter")
	sort := c.Query("sort")

	items, err := h.itemService.FindAll(filter, sort)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	// bukus := []song.SongResponse{}
	// for _, b := range songs {
	// 	buku := convertToResponse(b)
	// 	bukus = append(bukus, buku)
	// }
	c.JSON(http.StatusOK, gin.H{
		"data":   items,
		"Filter": filter,
		"Sort":   sort,
	})
}

// func (h *handlerTag) GetTagByID(c *gin.Context) {
// 	idString := c.Param("id")
// 	id, err := strconv.Atoi(idString)

// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": err,
// 		})
// 		return
// 	}

// 	b, err := h.tagService.FindByID(id)
// 	// buku := convertToResponse(b)

// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": b,
// 	})

// }

func (h *handlerTag) Create(c *gin.Context) {
	var item item.ItemInput
	err := c.ShouldBind(&item)
	if err != nil {

		messages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
			messages = append(messages, errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": messages,
		})
		return

	}
	h.itemService.Create(item)

	c.JSON(http.StatusOK, gin.H{
		"msg": item,
	})
}

func (h *handlerTag) Order(c *gin.Context) {

	user_email := Ambil(c)
	user, err := h.userService.FindByEmail(user_email)
	if err != nil {
		println("err")
	}
	order, err := h.itemService.Order(user)

	c.JSON(http.StatusOK, gin.H{
		"msg": order,
	})
}

func (h *handlerTag) ACC(c *gin.Context) {

	var item item.OrderInput
	err := c.ShouldBind(&item)
	if err != nil {

		messages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
			messages = append(messages, errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": messages,
		})
		return

	}
	order, _ := h.itemService.ACC(item)

	c.JSON(http.StatusOK, gin.H{
		"msg": order,
	})
}

// func (h *handlerTag) UpdateTag(c *gin.Context) {
// 	var tag tag.TagInput

// 	err := c.ShouldBind(&tag)
// 	if err != nil {

// 		messages := []string{}

// 		for _, e := range err.(validator.ValidationErrors) {
// 			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
// 			messages = append(messages, errormsg)
// 		}
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": messages,
// 		})
// 		return

// 	}

// 	idString := c.Param("id")
// 	id, err := strconv.Atoi(idString)

// 	if err != nil {

// 		fmt.Println(err)

// 	}

// 	h.tagService.UpdateTag(id, tag)

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": tag,
// 	})
// }

// func (h *handlerTag) DeleteTag(c *gin.Context) {
// 	// var book book.BookInput

// 	idString := c.Param("id")
// 	id, err := strconv.Atoi(idString)

// 	if err != nil {

// 		fmt.Println(err)

// 	}

// 	tag, err := h.tagService.Delete(id)

// 	if err != nil {

// 		fmt.Println(err)

// 	}

// 	msg := fmt.Sprintf("tag %s berhasil didelete", tag.Tag)

// 	c.JSON(http.StatusOK, gin.H{
// 		"msg": msg,
// 	})
// }

// // ============================================================

// func (h *handlerTag) GetAllSongTags(c *gin.Context) {

// 	songtags, err := h.tagService.FindAllSongTags()

// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": err,
// 		})
// 		return
// 	}
// 	// bukus := []song.SongResponse{}
// 	// for _, b := range songs {
// 	// 	buku := convertToResponse(b)
// 	// 	bukus = append(bukus, buku)
// 	// }
// 	c.JSON(http.StatusOK, songtags)
// }

// func (h *handlerTag) AddTag(c *gin.Context) {
// 	var songtag tag.SongTagInput

// 	err := c.ShouldBind(&songtag)
// 	if err != nil {

// 		messages := []string{}

// 		for _, e := range err.(validator.ValidationErrors) {
// 			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
// 			messages = append(messages, errormsg)
// 		}
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": messages,
// 		})
// 		return

// 	}

// 	b, err := h.songService.FindByID(songtag.SongID)

// 	if b.ID == 0 {
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": "lagu tidak ditemukan",
// 		})
// 		return
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	t, err := h.tagService.FindByID(songtag.TagID)

// 	if t.ID == 0 {
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": "tag tidak ditemukan",
// 		})
// 		return
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	f, err := h.tagService.FindSongTag(songtag.TagID, songtag.SongID)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if len(f) > 0 {
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg":         "lagu ini sudah memiliki tag tsb",
// 			"carisongtag": f,
// 		})
// 		return
// 	}

// 	msg := fmt.Sprintf("berhasil menambah tag %s ke lagu %s", t.Tag, b.YtID)

// 	h.tagService.AddTag(songtag, t.Tag, b.YtID)

// 	c.JSON(http.StatusOK, gin.H{
// 		"msg":         msg,
// 		"carilagu":    b,
// 		"caritag":     t,
// 		"carisongtag": f,
// 	})
// }
// func (h *handlerTag) GetSongByTag(c *gin.Context) {
// 	// var songtag tag.SongTagInput
// 	tagid := c.Param("tag")
// 	id, err := strconv.Atoi(tagid)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	f, err := h.tagService.GetSongByTag(id)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	c.JSON(http.StatusOK, f)
// }

// func (h *handlerTag) GetTagsBySong(c *gin.Context) {
// 	// var songtag tag.SongTagInput
// 	songid := c.Param("song")
// 	id, err := strconv.Atoi(songid)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	f, err := h.tagService.GetTagsBySong(id)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	c.JSON(http.StatusOK, f)
// }

// func (h *handlerTag) DeleteSongTag(c *gin.Context) {
// 	// var book book.BookInput

// 	idString := c.Param("id")
// 	id, err := strconv.Atoi(idString)

// 	if err != nil {

// 		fmt.Println(err)

// 	}

// 	tag, err := h.tagService.DeleteSongTag(id)

// 	if err != nil {

// 		fmt.Println(err)

// 	}

// 	msg := fmt.Sprintf("songtag %s dengan Video %s berhasil didelete", tag.Tag, tag.YtID)

// 	c.JSON(http.StatusOK, gin.H{
// 		"msg": msg,
// 	})
// }

// func (h *handlerTag) FilterTag(c *gin.Context) {
// 	var tagss tag.FilterInput

// 	err := c.ShouldBind(&tagss)
// 	if err != nil {

// 		messages := []string{}

// 		for _, e := range err.(validator.ValidationErrors) {
// 			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
// 			messages = append(messages, errormsg)
// 		}
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": messages,
// 		})
// 		return

// 	}

// 	f, err := h.tagService.FilterTag(tagss)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	c.JSON(http.StatusOK, f)
// }

// // func convertToResponseTag(b song.Song) song.SongResponse {

// // 	buku := song.SongResponse{
// // 		ID:   b.ID,
// // 		YtID: b.YtID,
// // 		// Title:       b.Title,
// // 		// Price:       b.Price,
// // 		// Description: b.Description,
// // 		// Rating:      b.Rating,
// // 		// Discount:    b.Discount,
// // 	}
// // 	return buku

// // }
