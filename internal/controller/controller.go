package controller

import (
	"net/http"

	"github.com/KartoonYoko/alice-skill/internal/config"
	"github.com/gin-gonic/gin"
)

type controller struct {
	r      *gin.Engine
	config *config.Config
}

func (c *controller) webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`
      {
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
      }
    `))
}

func (c *controller) notAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func New(conf *config.Config) *controller {
	r := gin.Default()
	c := &controller{
		r:      r,
		config: conf,
	}

	root := r.Group("/")
	{
		root.POST("/", func(ctx *gin.Context) {
			c.webhook(ctx.Writer, ctx.Request)
		})
		root.GET("/", func(ctx *gin.Context) {
			c.notAllowed(ctx.Writer, ctx.Request)
		})
		root.PUT("/", func(ctx *gin.Context) {
			c.notAllowed(ctx.Writer, ctx.Request)
		})
		root.DELETE("/", func(ctx *gin.Context) {
			c.notAllowed(ctx.Writer, ctx.Request)
		})
	}

	return c
}

func (c *controller) Serve() error {
	var err error
	if err = c.r.Run(c.config.FlagRunAddr); err == http.ErrServerClosed {
		return nil
	}

	return err
}
