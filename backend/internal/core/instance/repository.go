package instance

import (
	"database/sql"
	"fmt"
)

type InstanceRepository struct {
	db *sql.DB
}

func NewInstanceRepository(db *sql.DB) *InstanceRepository {
	return &InstanceRepository{db}
}

func (repository InstanceRepository) Create(instance Instance) (uint64, error) {
	query := `
		INSERT INTO instances (name, ram, disk, vcpus, osvariant, console, location, extraargs)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var instanceID uint64
	err := repository.db.QueryRow(
		query,
		instance.Name,
		instance.RAM,
		instance.Disk,
		instance.Vcpus,
		instance.OsVariant,
		instance.Console,
		instance.Location,
		instance.ExtraArgs,
	).Scan(&instanceID)
	if err != nil {
		return 0, err
	}
	return instanceID, nil
}

func (repository InstanceRepository) GetAll(param string) ([]Instance, error) {
	param = fmt.Sprintf("%%%s%%", param)

	result, err := repository.db.Query(
		`SELECT id, name, ram, disk, vcpus, osvariant, console, location, extraargs FROM instances WHERE name ILIKE $1`,
		param,
	)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var instances []Instance

	for result.Next() {
		var instance Instance

		if err = result.Scan(
			&instance.ID,
			&instance.Name,
			&instance.RAM,
			&instance.Disk,
			&instance.Vcpus,
			&instance.OsVariant,
			&instance.Console,
			&instance.Location,
			&instance.ExtraArgs,
		); err != nil {
			return nil, err
		}

		instances = append(instances, instance)
	}

	return instances, nil
}

func (repository InstanceRepository) GetByID(ID uint64) (Instance, error) {
	result, err := repository.db.Query(
		`SELECT id, name, ram, disk, vcpus, osvariant, console, location, extraargs FROM instances WHERE id = $1`,
		ID,
	)
	if err != nil {
		return Instance{}, err
	}
	defer result.Close()

	var instance Instance
	if result.Next() {
		if err = result.Scan(
			&instance.ID,
			&instance.Name,
			&instance.RAM,
			&instance.Disk,
			&instance.Vcpus,
			&instance.OsVariant,
			&instance.Console,
			&instance.Location,
			&instance.ExtraArgs,
		); err != nil {
			return Instance{}, err
		}
	}
	return instance, nil
}

func (repository InstanceRepository) Update(ID uint64, instance Instance) error {
	statement, err := repository.db.Prepare(
		`UPDATE instances SET 
		name = $1,
		ram = $2,
		disk = $3,
		vcpus = $4,
		osvariant = $5,
		console = $6,
		location = $7,
		extraargs = $8
		WHERE id = $9`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		instance.Name,
		instance.RAM,
		instance.Disk,
		instance.Vcpus,
		instance.OsVariant,
		instance.Console,
		instance.Location,
		instance.ExtraArgs,
		ID); err != nil {
		return err
	}

	return nil
}

func (repository InstanceRepository) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM instances WHERE id = $1",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}
