package db

import (
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/concourse/atc"
)

//counterfeiter:generate . WorkerArtifact

// TODO-L Can this be consolidated with atc/runtime/types.go -> Artifact OR Alternatively, there shouldn't be a volume reference here
type WorkerArtifact interface {
	ID() int
	Name() string
	BuildID() int
	CreatedAt() time.Time
	Volume(teamID int) (CreatedVolume, bool, error)
}

type artifact struct {
	conn DbConn

	id        int
	name      string
	buildID   int
	createdAt time.Time
}

func (a *artifact) ID() int              { return a.id }
func (a *artifact) Name() string         { return a.name }
func (a *artifact) BuildID() int         { return a.buildID }
func (a *artifact) CreatedAt() time.Time { return a.createdAt }

func (a *artifact) Volume(teamID int) (CreatedVolume, bool, error) {
	where := map[string]any{
		"v.team_id":            teamID,
		"v.worker_artifact_id": a.id,
	}

	_, created, err := getVolume(a.conn, where)
	if err != nil {
		return nil, false, err
	}

	if created == nil {
		return nil, false, nil
	}

	return created, true, nil
}

func saveWorkerArtifact(tx Tx, conn DbConn, atcArtifact atc.WorkerArtifact) (WorkerArtifact, error) {

	var artifactID int

	values := map[string]any{
		"name": atcArtifact.Name,
	}

	if atcArtifact.BuildID != 0 {
		values["build_id"] = atcArtifact.BuildID
	}

	err := psql.Insert("worker_artifacts").
		SetMap(values).
		Suffix("RETURNING id").
		RunWith(tx).
		QueryRow().
		Scan(&artifactID)

	if err != nil {
		return nil, err
	}

	artifact, found, err := getWorkerArtifact(tx, conn, artifactID)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("not found")
	}

	return artifact, nil
}

func getWorkerArtifact(tx Tx, conn DbConn, id int) (WorkerArtifact, bool, error) {
	var (
		createdAtTime sql.NullTime
		buildID       sql.NullInt64
	)

	artifact := &artifact{conn: conn}

	err := psql.Select("id", "created_at", "name", "build_id").
		From("worker_artifacts").
		Where(sq.Eq{
			"id": id,
		}).
		RunWith(tx).
		QueryRow().
		Scan(&artifact.id, &createdAtTime, &artifact.name, &buildID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	artifact.createdAt = createdAtTime.Time
	artifact.buildID = int(buildID.Int64)

	return artifact, true, nil
}
