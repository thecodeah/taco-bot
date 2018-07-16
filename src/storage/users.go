package storage

import (
	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2/bson"
)

// User stores information about the user.
type User struct {
	Guild   string `bson:"guild"`
	ID      string `bson:"id"`
	Balance int    `bson:"balance"`
}

// EnsureUser checks if the user exists in the database, if not it
// will create the user in the database.
func (db Database) EnsureUser(userid, guildid string) error {
	inDatabase, err := db.IsUserInDatabase(userid, guildid)
	if err != nil {
		return err
	}

	if !inDatabase {
		err = db.CreateUser(User{
			Guild: guildid,
			ID:    userid,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// EnsureUsers calls EnsureUser on all players in all guilds.
func (db Database) EnsureUsers(session *discordgo.Session) error {
	userGuilds, err := session.UserGuilds(10, "", "")
	if err != nil {
		return err
	}

	for _, guild := range userGuilds {
		members, err := session.GuildMembers(guild.ID, "", 10)
		if err != nil {
			return err
		}

		for _, member := range members {
			err = db.EnsureUser(member.User.ID, guild.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CreateUser creates the user in the database.
func (db Database) CreateUser(user User) (err error) {
	collection := db.session.DB(db.config.Name).C("users")
	err = collection.Insert(&user)
	return
}

// IsUserInDatabase checks if the user is in the database.
func (db Database) IsUserInDatabase(userid, guildid string) (bool, error) {
	collection := db.session.DB(db.config.Name).C("users")

	count, err := collection.Find(bson.M{"guild": guildid, "id": userid}).Count()
	if err != nil {
		return false, err
	}

	if count != 0 {
		return true, nil
	}

	return false, nil
}

// GetUser returns a user struct containing user data.
func (db Database) GetUser(userid, guildid string) (User, error) {
	var (
		result User
		err    error
	)

	collection := db.session.DB(db.config.Name).C("users")
	err = collection.Find(bson.M{"guild": guildid, "id": userid}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// UpdateUser updates a user's account data.
func (db Database) UpdateUser(user User) (err error) {
	collection := db.session.DB(db.config.Name).C("users")
	err = collection.Update(bson.M{"id": user.ID}, user)
	return
}
