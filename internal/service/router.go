package service

import (
	"github.com/dl-nft-books/nonce-auth-svc/internal/config"
	"github.com/dl-nft-books/nonce-auth-svc/internal/data/pg"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/handlers"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/helpers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxDB(pg.NewMasterQ(cfg.DB())),
			helpers.CtxServiceConfig(cfg.ServiceConfig()),
			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
			helpers.CtxNetworkConnector(*cfg.NetworkConnector()),
		),
	)

	r.Route("/integrations/nonce-auth-svc", func(r chi.Router) {
		r.Post("/nonce", handlers.GetNonce)
		r.Post("/register", handlers.Register)
		r.Get("/refresh-token", handlers.RefreshToken)
		r.Get("/created-at", handlers.CreatedAt)

		r.Get("/validate", handlers.Validate)

		r.Route("/login", func(r chi.Router) {
			//r.Post("/", handlers.Login)
			r.Post("/admin", handlers.AdminLogin)
		})

		r.Route("/users", func(r chi.Router) {
			r.Post("/", handlers.CreateUser)
			r.Get("/", handlers.GetListUsers)
			r.Route("/{address}", func(r chi.Router) {
				r.Patch("/", handlers.UpdateUserByAddress)
				r.Get("/", handlers.GetUserByAddress)
				r.Delete("/", handlers.DeleteUserByAddress)
			})
		})
	})

	return r
}
