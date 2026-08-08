package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"

	"example.com/go-db-test-improvement/db/migrations"
	"example.com/go-db-test-improvement/internal/api"
	memberdomain "example.com/go-db-test-improvement/internal/domain/member"
	memberpostgres "example.com/go-db-test-improvement/internal/infrastructure/member/postgres"
)

// handler implements api.ServerInterface. Only the member context is wired
// up here; the other eight bounded contexts follow the identical
// domain -> infrastructure/postgres -> handler pattern and are omitted from
// main.go for brevity since they are not the focus of this reproduction.
type handler struct {
	members memberdomain.Repository
}

func (h *handler) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *handler) ListMembers(c echo.Context) error {
	members, err := h.members.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	resp := make([]api.Member, 0, len(members))
	for _, m := range members {
		resp = append(resp, api.Member{Id: m.ID.String(), Name: m.Name, Email: m.Email})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *handler) CreateMember(c echo.Context) error {
	var req api.CreateMemberRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	m := &memberdomain.Member{
		ID:        uuid.New(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().UTC(),
	}
	if err := h.members.Create(c.Request().Context(), m); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, api.Member{Id: m.ID.String(), Name: m.Name, Email: m.Email})
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://app:app@localhost:5432/app?sslmode=disable"
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("set goose dialect: %v", err)
	}
	goose.SetBaseFS(migrations.FS)
	if err := goose.Up(db, "."); err != nil {
		log.Fatalf("run migrations: %v", err)
	}

	e := echo.New()
	api.RegisterHandlers(e, &handler{members: memberpostgres.New(db)})

	log.Fatal(e.Start(":8080"))
}
