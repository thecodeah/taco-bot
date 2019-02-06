package storage

import (
	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Guild : stores information about the guild.
type Guild struct {
	GuildID string `bson:"guildid"`
	Balance int    `bson:"balance"`
}

func (db Database) ensureGuildsCollection() (collection *mgo.Collection, err error) {
	collection = db.session.DB(db.config.Name).C("guilds")

	err = db.session.DB(db.config.Name).C("guilds").EnsureIndex(mgo.Index{
		Key:    []string{"guildid"},
		Unique: true,
	})
	return
}

// EnsureGuild checks if the guild exists in the database, if not it
// will create the guild in the database.
func (db Database) EnsureGuild(guildid string) (guild Guild, err error) {
	inDatabase, err := db.IsGuildInDatabase(guildid)
	if err != nil {
		return
	}

	if !inDatabase {
		guild = Guild{
			GuildID: guildid,
			Balance: 0,
		}

		err = db.CreateGuild(guild)
		if err != nil {
			return
		}
	}

	return guild, nil
}

// IsGuildInDatabase checks if the guild is in the database.
func (db Database) IsGuildInDatabase(guildid string) (bool, error) {
	count, err := db.guilds.Find(bson.M{"guildid": guildid}).Count()
	if err != nil {
		return false, err
	}

	if count != 0 {
		return true, nil
	}

	return false, nil
}

// CreateGuild creates the guild in the database.
func (db Database) CreateGuild(guildInfo Guild) (err error) {
	err = db.guilds.Insert(&guildInfo)
	return
}

// GetGuild returns a guild struct containing guild data.
func (db Database) GetGuild(guildid string) (guild Guild, err error) {
	err = db.guilds.Find(bson.M{"guildid": guildid}).One(&guild)
	if err != nil {
		// Ensure the guild if something goes wrong.
		guild, err = db.EnsureGuild(guildid)
		if err != nil {
			return guild, err
		}
	}

	return guild, nil
}

// IncreaseGuildBalance increases the balance of the guild by the specified
// amount.
func (db Database) IncreaseGuildBalance(guild Guild, amount int) (err error) {
	err = db.guilds.Update(bson.M{"guildid": guild.GuildID}, bson.M{"$inc": bson.M{"balance": amount}})
	return
}
