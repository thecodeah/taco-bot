package storage

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

// Database stores information about the database connection.
type Database struct {
	session *mgo.Session
	users   *mgo.Collection
	guilds  *mgo.Collection
	config  Config
}

// Config stores configuration for the database.
type Config struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

// Connect : Connects to a MongoDB database and returns a database struct.
func Connect(config Config) (*Database, error) {
	var (
		err      error
		database Database
	)

	database.config = config

	database.session, err = mgo.Dial(config.Host + ":" + config.Port)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to dial : %s:%s", config.Host, config.Port)
	}

	if config.Pass != "" {
		database.session.Login(&mgo.Credential{
			Source:   config.Name,
			Username: config.User,
			Password: config.Pass,
		})
		if err != nil {
			return nil, errors.Wrap(err, "Failed to login.")
		}
	}

	database.users, err = database.ensureUsersCollection()
	if err != nil {
		return nil, err
	}

	database.guilds, err = database.ensureGuildsCollection()
	if err != nil {
		return nil, err
	}

	return &database, nil
}
