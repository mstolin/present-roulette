package db

import (
	"fmt"

	"github.com/mstolin/present-roulette/utils/models"
)

func (dbHandler DatabaseHandler) GetItemByUser(user models.User, itemId int) (models.Item, error) {
	assoc := dbHandler.DB.Model(&user).Where("id = ?", itemId).Association("Items")

	var item models.Item
	if count := assoc.Count(); count > 0 {
		if err := assoc.Find(&item); err != nil {
			return item, err
		} else {
			return item, nil
		}
	} else {
		return item, fmt.Errorf("no item found with id %d", itemId)
	}
}

func (dbHandler DatabaseHandler) AddItemsToUser(user models.User, items []models.Item) error {
	if err := dbHandler.DB.Model(&user).Association("Items").Append(&items); err != nil {
		return err
	} else {
		return nil
	}
}

func (dbHandler DatabaseHandler) UpdateItemByUser(user models.User, itemId int, update models.ItemUpdate) (models.Item, error) {
	item, err := dbHandler.GetItemByUser(user, itemId)
	if err != nil {
		return item, err
	}

	// TODO This is dirty, there has to be a better way
	clone := item
	clone.Name = update.Name
	clone.VendorID = update.VendorID
	clone.Price = update.Price
	clone.HasBeenBaught = update.HasBeenBaught

	if err := dbHandler.DB.Model(&item).Updates(clone).Error; err != nil {
		return item, err
	} else {
		return item, nil
	}
}

func (dbHandler DatabaseHandler) DeleteItem(user models.User, itemId int) (models.Item, error) {
	item, err := dbHandler.GetItemByUser(user, itemId)
	if err != nil {
		return item, err
	}

	if err := dbHandler.DB.Model(&user).Association("Items").Delete(item); err != nil {
		return item, err
	} else {
		return item, nil
	}
}
