package router

import (
	"time"

	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Router struct {
}

func DefaultStructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		param := gin.LogFormatterParams{
			TimeStamp:    time.Now(),
			Latency:      time.Since(start),
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			StatusCode:   c.Writer.Status(),
			ErrorMessage: c.Errors.ByType(gin.ErrorTypePrivate).String(),
			BodySize:     c.Writer.Size(),
			Path:         c.Request.URL.Path,
		}

		var logEvent *logrus.Entry
		if c.Writer.Status() >= 500 {
			logEvent = logrus.WithFields(logrus.Fields{
				"client_id":   param.ClientIP,
				"method":      param.Method,
				"status_code": param.StatusCode,
				"body_size":   param.BodySize,
				"path":        param.Path,
				"latency":     param.Latency.String(),
				"query":       c.Request.URL.RawQuery,
			})
			logEvent.Error(param.ErrorMessage)
		} else {
			logEvent = logrus.WithFields(logrus.Fields{
				"client_id":   param.ClientIP,
				"method":      param.Method,
				"status_code": param.StatusCode,
				"body_size":   param.BodySize,
				"path":        param.Path,
				"latency":     param.Latency.String(),
			})
			logEvent.Info(param.ErrorMessage)
		}
	}
}

func (r *Router) IniteRoutes(rep *repository.Repository, service *services.Service, authMiddleware gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(DefaultStructuredLogger()) // Добавление пользовательского middleware для логирования

	authRouter := new(AuthRouter)
	authRouter.RegisterRoutes(router, rep, service)

	userRouter := new(UserRouter)
	userRouter.RegisterRoutes(router, rep, authMiddleware)

	workoutRouter := new(WorkoutRouter)
	workoutRouter.RegisterRoutes(router, rep, authMiddleware)

	goalRouter := new(GoalRouter)
	goalRouter.RegisterRoutes(router, rep, authMiddleware)

	achievemetRouter := new(AchievementRouter)
	achievemetRouter.RegisterRoutes(router, rep, authMiddleware)

	return router
}
