package db

import (
	"database/sql"

	"github.com/kelseyduffy/sporting-events/models"
)

func (db Database) GetAllTeams() (*models.TeamList, error) {
	list := &models.TeamList{}
	rows, err := db.Conn.Query("SELECT * FROM teams ORDER BY id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.ID, &team.Name, &team.FoundedYear, &team.DissolvedYear, &team.Sport)
		if err != nil {
			return list, err
		}
		list.Teams = append(list.Teams, team)
	}
	return list, nil
}
func (db Database) AddTeam(team *models.Team) error {
	var id int
	query := `INSERT INTO teams (name, founded_year, dissolved_year, sport) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.Conn.QueryRow(query, team.Name, team.FoundedYear, team.DissolvedYear, team.Sport).Scan(&id)
	if err != nil {
		return err
	}
	team.ID = id
	return nil
}
func (db Database) GetTeamById(teamId int) (models.Team, error) {
	team := models.Team{}
	query := `SELECT * FROM teams WHERE id = $1;`
	row := db.Conn.QueryRow(query, teamId)
	switch err := row.Scan(&team.ID, &team.Name, &team.FoundedYear, &team.DissolvedYear, &team.Sport); err {
	case sql.ErrNoRows:
		return team, ErrNoMatch
	default:
		return team, err
	}
}
func (db Database) DeleteTeam(teamId int) error {
	query := `DELETE FROM teams WHERE id = $1;`
	_, err := db.Conn.Exec(query, teamId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
func (db Database) UpdateTeam(teamId int, teamData models.Team) (models.Team, error) {
	team := models.Team{}
	query := `UPDATE teams SET name=$1, founded_year=$2, dissolved_year=$3, sport=$4 WHERE id=$5 RETURNING id, name, founded_year, dissolved_year, sport;`
	err := db.Conn.QueryRow(query, teamData.Name, teamData.FoundedYear, teamData.DissolvedYear, teamData.Sport, teamId).Scan(&team.ID, &team.Name,
		&team.FoundedYear, &team.DissolvedYear, &team.Sport)
	if err != nil {
		if err == sql.ErrNoRows {
			return team, ErrNoMatch
		}
		return team, err
	}
	return team, nil
}
