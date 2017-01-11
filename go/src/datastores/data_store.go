/*Package datastores implements the basics databases communication functions used by PopCube chat api.

Stores

The following is a list of stores described:
	Avatar: Contain all informations for avatar management
	Channel: Contain all informations for channel management
	Emojis: Contain all informations for emojis management
	Organisation: Contain all informations for organisation management
	Parameter: Contain all informations for parmeters management
	Role: Contain all informations for roles management
	User: Contain all informations for users management
*/
// Created by Titouan FREVILLE <titouanfreville@gmail.com>
//
// Inspired by mattermost project
package datastores

import (
	// Importing sql driver. They are used by gorm package and used by default from blank.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"models"
	u "utils"
)

// dbStore Struct to manage Db knowledge
type dbStore struct {
	Db  *gorm.DB
	Err error
}

// type StoreResult struct {
// 	Data interface{}
// 	Err  *models.AppError
// }

// type StoreChannel chan StoreResult

// func Must(sc StoreChannel) interface{} {
// 	r := <-sc
// 	if r.Err != nil {
// 		l4g.Close()
// 		time.Sleep(time.Second)
// 		panic(r.Err)
// 	}

// 	return r.Data
// }

// InitConnection init Database connection && database models
func (ds *dbStore) InitConnection(user string, dbname string, password string) {
	connectionChain := user + ":" + password + "@(database:3306)/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionChain)
	// db.AutoMigrate( &models.Channel{}, &models.Emoji{}, &models.Folder{},
	// 	&models.Member{}, &models.Message{}, &models.Organisation{}, ,
	// 	&models.Role{}, &models.User{})
	db.AutoMigrate(&models.Avatar{}, &models.Emoji{}, &models.Organisation{}, &models.Parameter{}, &models.Role{})
	ds.Db = db
	ds.Err = err
}

// CloseConnection close database connection
func (ds *dbStore) CloseConnection() {
	db := *ds.Db
	defer db.Close()
	ds.Db = &gorm.DB{}
}

// Store interface the Stores and usefull DB functions
type Store interface {
	Organisation() OrganisationStore
	Avatar() AvatarStore
	Emoji() EmojiStore
	InitConnection()
	CloseConnection()
}

/*AvatarStore interface the avatar communication */
type AvatarStore interface {
	Save(avatar *models.Avatar, ds dbStore) *u.AppError
	Update(avatar *models.Avatar, newAvatar *models.Avatar, ds dbStore) *u.AppError
	GetByName(avatarName string, ds dbStore) *models.Avatar
	GetByLink(avatarLink string, ds dbStore) *models.Avatar
	GetAll(ds dbStore) *models.Avatar
	Delete(avatar *models.Avatar, ds dbStore) *u.AppError
}

/*ChannelStore interface the channel communication*/
type ChannelStore interface {
	Save(channel *models.Channel, ds dbStore) *u.AppError
	Update(channel *models.Channel, newChannel *models.Channel, ds dbStore) *u.AppError
	GetByName(channelName string, ds dbStore) *models.Channel
	GetByType(channelType string, ds dbStore) *models.Channel
	GetPublic(ds dbStore) *[]models.Channel
	GetPrivate(ds dbStore) *[]models.Channel
	GetAll(ds dbStore) *[]models.Channel
	Delete(channel *models.Channel, ds dbStore) *u.AppError
}

/*EmojiStore interface the emoji communication*/
type EmojiStore interface {
	Save(emoji *models.Emoji, ds dbStore) *u.AppError
	Update(emoji *models.Emoji, newEmoji *models.Emoji, ds dbStore) *u.AppError
	GetByName(emojiName string, ds dbStore) *models.Emoji
	GetByShortcut(emojiShortcut string, ds dbStore) *models.Emoji
	GetByLink(emojiLink string, ds dbStore) *models.Emoji
	GetAll(ds dbStore) *models.Emoji
	Delete(emoji *models.Emoji, ds dbStore) *u.AppError
}

/*FolderStore interface communication with message table*/
type FolderStore interface {
	Save(message *models.Folder, ds dbStore) *u.AppError
	Update(message *models.Folder, newFolder *models.Folder, ds dbStore) *u.AppError
	GetByName(messageName string, ds dbStore) *[]models.Folder
	GetByType(messageType string, ds dbStore) *[]models.Folder
	GetByLink(messageLink string, ds dbStore) *[]models.Folder
	GetByMessage(message *models.Message, ds dbStore) *[]models.Folder
	GetAll(ds dbStore) *[]models.Folder
	Delete(message *models.Folder, ds dbStore) *u.AppError
}

/*MemberStore interface communication with member table*/
type MemberStore interface {
	Save(emoji *models.Emoji, ds dbStore) *u.AppError
	Update(emoji *models.Emoji, newEmoji *models.Emoji, ds dbStore) *u.AppError
	GetByUser()
	GetByChannel()
	GetByRole()
	GetAll(ds dbStore) *models.Emoji
	Delete(emoji *models.Emoji, ds dbStore) *u.AppError
}

/*MessageStore interface communication with message table*/
type MessageStore interface {
	Save(message *models.Message, ds dbStore) *u.AppError
	Update(message *models.Message, newMessage *models.Message, ds dbStore) *u.AppError
	GetByDate(messageDate int, ds dbStore) *[]models.Message
	GetByCreator(creator *models.User, ds dbStore) *[]models.Message
	GetByChannel(channel *models.Channel, ds dbStore) *[]models.Message
	GetAll(ds dbStore) *[]models.Message
	Delete(message *models.Message, ds dbStore) *u.AppError
}

/*OrganisationStore interface the organisation communication
Organisation is unique in the database. So they are no use of providing an user to get.
Delete is useless as we will down the docker stack in case an organisation leace.
*/
type OrganisationStore interface {
	Save(organisation *models.Organisation, ds dbStore) *u.AppError
	Update(organisation *models.Organisation, newOrganisation *models.Organisation, ds dbStore) *u.AppError
	Get(ds dbStore) *models.Organisation
}

/*ParameterStore interface the parameter communication*/
type ParameterStore interface {
	Save(parameter *models.Parameter, ds dbStore) *u.AppError
	Update(parameter *models.Parameter, newParameter *models.Parameter, ds dbStore) *u.AppError
	GetAll(ds dbStore) *models.Parameter
}

/*RoleStore interface the role communication*/
type RoleStore interface {
	Save(role *models.Role, ds dbStore) *u.AppError
	Update(role *models.Role, newRole *models.Role, ds dbStore) *u.AppError
	GetByName(roleName string, ds dbStore) *models.Role
	GetByRights(roleRights *models.Role, ds dbStore) *[]models.Role
	GetAll(ds dbStore) *[]models.Role
	Delete(role *models.Role, ds dbStore) *u.AppError
}

/*UserStore interface the user communication*/
type UserStore interface {
	Save(user *models.User, ds dbStore) *u.AppError
	Update(user *models.User, newUser *models.User, ds dbStore) *u.AppError
	GetByUserName(userName string, ds dbStore) *models.User
	GetByEmail(userEmail string, ds dbStore) *models.User
	GetByDate(userDate int, ds dbStore) *[]models.User
	GetDeleted(ds dbStore) *[]models.User
	GetByNickName(nickName string, ds dbStore) *models.User
	GetByFirstName(firstName string, ds dbStore) *[]models.User
	GetByLastName(lastName string, ds dbStore) *[]models.User
	GetByRole(role *models.Role, ds dbStore) *[]models.User
	GetByChannel(channel *models.Channel, ds dbStore) *[]models.User
	GetAll(ds dbStore) *[]models.User
	Delete(user *models.User, ds dbStore) *u.AppError
}
