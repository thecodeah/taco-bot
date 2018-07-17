package storage

import (
	"github.com/bwmarrin/discordgo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User stores information about the user.
type User struct {
	GuildID string `bson:"guildid"`
	UserID  string `bson:"userid"`
	Balance int    `bson:"balance"`
}

func (db Database) ensureUsersCollection() (collection *mgo.Collection, err error) {
	collection = db.session.DB(db.config.Name).C("users")

	err = db.session.DB(db.config.Name).C("users").EnsureIndex(mgo.Index{
		Key:    []string{"guildid", "userid"},
		Unique: true,
	})
	return
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
			GuildID: guildid,
			UserID:  userid,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// EnsureUsers calls EnsureUser on all players in all guilds.
func (db Database) EnsureUsers(session *discordgo.Session) error {
	userGuilds, err := session.UserGuilds(100, "", "")
	if err != nil {
		return err
	}

	for _, guild := range userGuilds {
		members, err := session.GuildMembers(guild.ID, "", 1000)
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
	err = db.users.Insert(&user)
	return
}

// IsUserInDatabase checks if the user is in the database.
func (db Database) IsUserInDatabase(userid, guildid string) (bool, error) {
	count, err := db.users.Find(bson.M{"guildid": guildid, "userid": userid}).Count()
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

	err = db.users.Find(bson.M{"guildid": guildid, "userid": userid}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// UpdateUser updates a user's account data.
func (db Database) UpdateUser(user User) (err error) {
	err = db.users.Update(bson.M{"guildid": user.GuildID, "userid": user.UserID}, user)
	return
}
