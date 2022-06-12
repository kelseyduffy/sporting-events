package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kelseyduffy/sporting-events/db"
	"github.com/kelseyduffy/sporting-events/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var teamIDKey = "teamID"

func teams(router chi.Router) {
	router.Get("/teams", getAllTeams)
	router.Post("/teams", createTeam)
	router.Route("/teams/{teamID}", func(router chi.Router) {
		router.Use(TeamContext)
		router.Get("/", getTeam)
		router.Put("/", updateTeam)
		router.Delete("/", deleteTeam)
	})
}
func TeamContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		teamId := chi.URLParam(r, "teamID")
		if teamId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("team ID is required")))
			return
		}
		id, err := strconv.Atoi(teamId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid team ID")))
		}
		ctx := context.WithValue(r.Context(), teamIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createTeam(w http.ResponseWriter, r *http.Request) {
	team := &models.Team{}
	if err := render.Bind(r, team); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddTeam(team); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, team); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getAllTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := dbInstance.GetAllTeams()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, teams); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(teamIDKey).(int)
	team, err := dbInstance.GetTeamById(teamID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &team); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteTeam(w http.ResponseWriter, r *http.Request) {
	teamId := r.Context().Value(teamIDKey).(int)
	err := dbInstance.DeleteTeam(teamId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
func updateTeam(w http.ResponseWriter, r *http.Request) {
	teamId := r.Context().Value(teamIDKey).(int)
	teamData := models.Team{}
	if err := render.Bind(r, &teamData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	team, err := dbInstance.UpdateTeam(teamId, teamData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &team); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
