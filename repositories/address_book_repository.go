package repositories

import (
	"github.com/mahfuzon/address_book_app/models"
	"gorm.io/gorm"
)

type AddressBookRepository interface {
	Create(contact models.Contact) (models.Contact, error)
	Delete(contact models.Contact) error
	FindByName(name string) (models.Contact, error)
	FindByPhoneNumberOrName(phoneNumber, name string) (models.Contact, error)
	FindBySlug(slug string) (models.Contact, error)
}

type addressBookRepository struct {
	db *gorm.DB
}

func NewAddressBookRepository(db *gorm.DB) AddressBookRepository {
	return &addressBookRepository{db: db}
}

func (addressBookRepository *addressBookRepository) Create(contact models.Contact) (models.Contact, error) {
	err := addressBookRepository.db.Create(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (addressBookRepository *addressBookRepository) Delete(contact models.Contact) error {
	err := addressBookRepository.db.Delete(&contact).Error
	if err != nil {
		return err
	}

	return nil
}

func (addressBookRepository *addressBookRepository) FindByName(name string) (models.Contact, error) {
	contact := models.Contact{}
	err := addressBookRepository.db.Where("name like ?", "%"+name+"%").First(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (addressBookRepository *addressBookRepository) FindBySlug(slug string) (models.Contact, error) {
	contact := models.Contact{}
	err := addressBookRepository.db.Where("slug = ?", slug).First(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (addressBookRepository *addressBookRepository) FindByPhoneNumberOrName(phoneNumber, name string) (models.Contact, error) {
	contact := models.Contact{}
	err := addressBookRepository.db.Where("name = ?", name).Or("phone_number = ?", phoneNumber).First(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}
