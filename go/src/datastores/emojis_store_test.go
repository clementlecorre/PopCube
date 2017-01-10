// This file is used to test if user model is working correctly.
// A user is always linked to an emoji
// He has basic channel to join
package datastores

import (
	. "github.com/smartystreets/goconvey/convey"
	. "models"
	"testing"
	u "utils"
)

func TestEmojiStore(t *testing.T) {
	ds := DataStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db
	asi := EmojiStoreImpl{}
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("emojiStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("emojiStoreImpl.Save", "save.transaction.create.already_exist", nil, "Emoji Name: Troll Face")
		emoji := Emoji{
			Name:     "Troll Face",
			Shortcut: ":troll-face:",
			Link:     "emojis/trollface.svg",
		}
		Convey("Given a correct emoji.", func() {
			appError := asi.Save(&emoji, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := asi.Save(&emoji, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&emoji)
	})

	Convey("Testing update function", t, func() {
		dbError := u.NewLocAppError("emojiStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("emojiStoreImpl.Save", "save.transaction.create.already_exist", nil, "Emoji Name: Troll Face")
		emoji := Emoji{
			Name:     "Troll Face",
			Shortcut: ":troll-face:",
			Link:     "emojis/trollface.svg",
		}
		emojiNew := Emoji{
			Name:     "Troll Face2",
			Shortcut: ":troll:",
			Link:     "emojis/trollface2.svg",
		}

		appError := asi.Save(&emoji, ds)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		So(appError, ShouldNotResemble, alreadyExistError)

		Convey("Provided correct Emoji to modify should not return errors", func() {
			appError := asi.Update(&emoji, &emojiNew, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
		})

		Convey("Provided wrong Emoji to modify should result in old_emoji error", func() {
			emoji.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiOld.PreSave", "model.emoji.name.app_error", nil, ""))
				emoji.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiOld.PreSave", "model.emoji.name.app_error", nil, ""))
			})

			emoji.Name = "Correct Name"
			emoji.Link = ""

			Convey("Empty link should result in link error", func() {
				appError = asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiOld.PreSave", "model.emoji.link.app_error", nil, ""))
			})

			emoji.Link = "emojis/trollface.svg"
			emoji.Shortcut = ":this-is-a-tool-long-shortcut:"

			Convey("Too long shortcut or empty shorctcut should return Shortcut error", func() {
				appError := asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiOld.PreSave", "model.emoji.shortcut.app_error", nil, ""))
				emoji.Shortcut = ""
				appError = asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiOld.PreSave", "model.emoji.shortcut.app_error", nil, ""))
			})
		})

		Convey("Provided wrong Emoji to modify should result in newEmoji error", func() {
			emojiNew.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiNew.PreSave", "model.emoji.name.app_error", nil, ""))
				emojiNew.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiNew.PreSave", "model.emoji.name.app_error", nil, ""))
			})

			emojiNew.Name = "Correct Name"
			emojiNew.Link = ""

			Convey("Empty link should result in link error", func() {
				appError = asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiNew.PreSave", "model.emoji.link.app_error", nil, ""))
			})

			emojiNew.Link = "emojis/trollface.svg"
			emojiNew.Shortcut = ":this-is-a-tool-long-shortcut:"

			Convey("Too long shortcut or empty shorctcut should return Shortcut error", func() {
				appError := asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiNew.PreSave", "model.emoji.shortcut.app_error", nil, ""))
				emojiNew.Shortcut = ""
				appError = asi.Update(&emoji, &emojiNew, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Update.emojiNew.PreSave", "model.emoji.shortcut.app_error", nil, ""))
			})
		})
		db.Delete(&emoji)
		db.Delete(&emojiNew)
	})

	Convey("Testing Getters", t, func() {
		emoji0 := Emoji{
			Name:     "Joy Face",
			Shortcut: ":)",
			Link:     "emojis/joyface.svg",
		}
		emoji1 := Emoji{
			Name:     "Troll Face",
			Shortcut: ":troll:",
			Link:     "emojis/trollface.svg",
		}
		emoji2 := Emoji{
			Name:     "nOOb",
			Shortcut: ":noob:",
			Link:     "emojis/sparadra.svg",
		}
		emoji1New := Emoji{
			Name:     "Troll Face NEW",
			Shortcut: ":troll:",
			Link:     "emojis/trollfacenew.svg",
		}
		emoji3 := Emoji{
			Name:     "Face Palm",
			Shortcut: ":facepalm:",
			Link:     "emojis/facepalm.svg",
		}
		asi.Save(&emoji0, ds)
		asi.Save(&emoji1, ds)
		asi.Update(&emoji1, &emoji1New, ds)
		asi.Save(&emoji2, ds)
		asi.Save(&emoji3, ds)
		// Have to be after save so ID are up to date :O
		emojiList := []Emoji{
			emoji0,
			emoji1,
			emoji2,
			emoji3,
		}

		Convey("We have to be able to find all emojis in the db", func() {
			emojis := asi.GetAll(ds)
			So(emojis, ShouldNotResemble, &[]Emoji{})
			So(emojis, ShouldResemble, &emojiList)
		})

		Convey("We have to be able to find an emoji from is name", func() {
			emoji := asi.GetByName(emoji0.Name, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji0)
			emoji = asi.GetByName(emoji2.Name, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji2)
			emoji = asi.GetByName(emoji3.Name, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji3)
			Convey("Should also work from updated value", func() {
				emoji = asi.GetByName(emoji1.Name, ds)
				So(emoji, ShouldNotResemble, &Emoji{})
				So(emoji, ShouldResemble, &emoji1)
			})
		})

		Convey("We have to be able to find an emoji from is link", func() {
			emoji := asi.GetByLink(emoji0.Link, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji0)
			emoji = asi.GetByLink(emoji2.Link, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji2)
			emoji = asi.GetByLink(emoji3.Link, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji3)
			Convey("Should also work from updated value", func() {
				emoji = asi.GetByLink(emoji1.Link, ds)
				So(emoji, ShouldNotResemble, &Emoji{})
				So(emoji, ShouldResemble, &emoji1)
			})
		})

		Convey("We have to be able to find an emoji from its shortcut", func() {
			emoji := asi.GetByShortcut(emoji0.Shortcut, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji0)
			emoji = asi.GetByShortcut(emoji2.Shortcut, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji2)
			emoji = asi.GetByShortcut(emoji3.Shortcut, ds)
			So(emoji, ShouldNotResemble, &Emoji{})
			So(emoji, ShouldResemble, &emoji3)
			Convey("Should also work from updated value", func() {
				emoji = asi.GetByShortcut(emoji1.Shortcut, ds)
				So(emoji, ShouldNotResemble, &Emoji{})
				So(emoji, ShouldResemble, &emoji1)
			})
		})

		Convey("Searching for non existent emoji should return empty", func() {
			emoji := asi.GetByLink("The Mask", ds)
			So(emoji, ShouldResemble, &Emoji{})
			emoji = asi.GetByName("Fantôme", ds)
			So(emoji, ShouldResemble, &Emoji{})
		})

		db.Delete(&emoji0)
		db.Delete(&emoji1)
		db.Delete(&emoji2)
		db.Delete(&emoji3)

		Convey("Searching all in empty table should return empty", func() {
			emojis := asi.GetAll(ds)
			So(emojis, ShouldResemble, &[]Emoji{})
		})
	})

	Convey("Testing delete emoji", t, func() {
		dberror := u.NewLocAppError("emojiStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
		emoji0 := Emoji{
			Name:     "Joy Face",
			Shortcut: ":)",
			Link:     "emojis/joyface.svg",
		}
		emoji1 := Emoji{
			Name:     "Troll Face",
			Shortcut: ":troll:",
			Link:     "emojis/trollface.svg",
		}
		emoji2 := Emoji{
			Name:     "nOOb",
			Shortcut: ":noob:",
			Link:     "emojis/sparadra.svg",
		}
		emoji3 := Emoji{
			Name:     "Face Palm",
			Shortcut: ":facepalm:",
			Link:     "emojis/facepalm.svg",
		}
		asi.Save(&emoji0, ds)
		asi.Save(&emoji1, ds)
		asi.Save(&emoji2, ds)
		asi.Save(&emoji3, ds)
		emoji3Old := emoji3
		emojiList1 := []Emoji{
			emoji0,
			emoji1,
			emoji2,
			emoji3Old,
		}

		Convey("Deleting a known emoji should work", func() {
			appError := asi.Delete(&emoji2, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dberror)
			So(asi.GetByName("God", ds), ShouldResemble, &Emoji{})
		})

		Convey("Trying to delete from non conform emoji should return specific emoji error and should not delete emojis.", func() {
			emoji3.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Delete(&emoji3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Delete.emoji.PreSave", "model.emoji.name.app_error", nil, ""))
				So(asi.GetAll(ds), ShouldResemble, &emojiList1)
				emoji3.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Delete(&emoji3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Delete.emoji.PreSave", "model.emoji.name.app_error", nil, ""))
				So(asi.GetAll(ds), ShouldResemble, &emojiList1)
			})

			emoji3.Name = "nOOb"
			emoji3.Link = ""

			Convey("Empty link should result in link error", func() {
				appError := asi.Delete(&emoji3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, u.NewLocAppError("emojiStoreImpl.Delete.emoji.PreSave", "model.emoji.link.app_error", nil, ""))
				So(asi.GetAll(ds), ShouldResemble, &emojiList1)
			})
		})

		db.Delete(&emoji0)
		db.Delete(&emoji1)
		db.Delete(&emoji2)
		db.Delete(&emoji3)
	})
}
