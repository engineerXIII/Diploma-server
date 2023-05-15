package server

import (
	"net/http"
	"strings"

	"github.com/engineerXIII/Diploma-server/pkg/csrf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// _ "github.com/engineerXIII/Diploma-server/docs"
	authHttp "github.com/engineerXIII/Diploma-server/internal/auth/delivery/http"
	authRepository "github.com/engineerXIII/Diploma-server/internal/auth/repository"
	authUseCase "github.com/engineerXIII/Diploma-server/internal/auth/usecase"
	commentsHttp "github.com/engineerXIII/Diploma-server/internal/comments/delivery/http"
	commentsRepository "github.com/engineerXIII/Diploma-server/internal/comments/repository"
	commentsUseCase "github.com/engineerXIII/Diploma-server/internal/comments/usecase"
	jobsHttp "github.com/engineerXIII/Diploma-server/internal/jobs/delivery/http"
	jobsUsecase "github.com/engineerXIII/Diploma-server/internal/jobs/usecase"
	apiMiddlewares "github.com/engineerXIII/Diploma-server/internal/middleware"
	routersHttp "github.com/engineerXIII/Diploma-server/internal/routers/delivery/http"
	routersRepository "github.com/engineerXIII/Diploma-server/internal/routers/repository"
	routersUseCase "github.com/engineerXIII/Diploma-server/internal/routers/usecase"
	sessionRepository "github.com/engineerXIII/Diploma-server/internal/session/repository"
	"github.com/engineerXIII/Diploma-server/internal/session/usecase"
	"github.com/engineerXIII/Diploma-server/pkg/utils"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	//metrics, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	//if err != nil {
	//	s.logger.Errorf("CreateMetrics Error: %s", err)
	//}
	//s.logger.Info(
	//	"Metrics available URL: %s, ServiceName: %s",
	//	s.cfg.Metrics.URL,
	//	s.cfg.Metrics.ServiceName,
	//)

	// Init repositories
	aRepo := authRepository.NewAuthRepository(s.db)
	cRepo := commentsRepository.NewCommentsRepository(s.db)
	rtRepo := routersRepository.NewRouterRepository(s.db)
	sRepo := sessionRepository.NewSessionRepository(s.redisClient, s.cfg)
	//aAWSRepo := authRepository.NewAuthAWSRepository(s.awsClient)
	authRedisRepo := authRepository.NewAuthRedisRepo(s.redisClient)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(s.cfg, aRepo, authRedisRepo, s.logger)
	rtUC := routersUseCase.NewRoutersUseCase(s.cfg, rtRepo, s.logger)
	jobsUC := jobsUsecase.NewJobsUseCase(s.cfg, s.logger)
	commUC := commentsUseCase.NewCommentsUseCase(s.cfg, cRepo, s.logger)
	sessUC := usecase.NewSessionUseCase(sRepo, s.cfg)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, sessUC, s.logger)
	rtHandlers := routersHttp.NewRoutersHandlers(s.cfg, rtUC, s.logger)
	jobHandlers := jobsHttp.NewJobsHandlers(s.cfg, jobsUC, s.logger)
	commHandlers := commentsHttp.NewCommentsHandlers(s.cfg, commUC, s.logger)

	mw := apiMiddlewares.NewMiddlewareManager(sessUC, authUC, s.cfg, []string{"*"}, s.logger)

	e.Use(mw.RequestLoggerMiddleware)

	//docs.SwaggerInfo.Title = "Go example REST API"
	//e.GET("/swagger/*", echoSwagger.WrapHandler)

	if s.cfg.Server.SSL {
		e.Pre(middleware.HTTPSRedirect())
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	//e.Use(mw.MetricsMiddleware(metrics))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("5M"))
	if s.cfg.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	authGroup := v1.Group("/auth")
	rtGroup := v1.Group("/router")
	jobGroup := v1.Group("/job")
	commGroup := v1.Group("/comments")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)
	commentsHttp.MapCommentsRoutes(commGroup, commHandlers, mw)
	routersHttp.MapRoutersRoutes(rtGroup, rtHandlers, mw)
	jobsHttp.MapJobsRoutes(jobGroup, jobHandlers, mw)

	health.GET("", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
