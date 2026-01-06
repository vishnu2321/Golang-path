package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const RANGE_FOR_ID = 1000

type Contact struct {
	Id      int
	Name    string
	PhoneNo string
	Address string
}

type ContactsSlice []Contact

func (contacts *ContactsSlice) CreateContact() error {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("To Create New Contact. Enter below details")
	fmt.Print("Enter the Name: ")
	nameStr, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Print("Enter the Phone Number: ")
	phnoStr, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Print("Enter the Address: ")
	AddStr, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	//generate id
	id := rand.Intn(RANGE_FOR_ID) + 1
	newContact := Contact{
		Id:      id,
		Name:    cleanString(nameStr),
		PhoneNo: cleanString(phnoStr),
		Address: cleanString(AddStr),
	}
	*contacts = append(*contacts, newContact)
	fmt.Println("Contact created successfully.")
	return nil
}

func (contacts *ContactsSlice) UpdateContact(id int, newContact Contact) error {
	found := false
	for i, contact := range *contacts {
		if contact.Id == id {
			if newContact.Name != "" {
				(*contacts)[i].Name = cleanString(newContact.Name)
			}
			if newContact.PhoneNo != "" {
				(*contacts)[i].PhoneNo = cleanString(newContact.PhoneNo)
			}
			if newContact.Address != "" {
				(*contacts)[i].Address = cleanString(newContact.Address)
			}
			fmt.Printf("Update successful for contact with ID:%d\n", id)
			return nil
		}
	}
	if !found {
		return errors.New("Contact with given id not found")
	}
	return nil
}

func (contacts *ContactsSlice) DeleteContact(id int) error {
	oldContactInd := -1
	for i, contact := range *contacts {
		if contact.Id == id {
			oldContactInd = i
			break
		}
	}
	if oldContactInd == -1 {
		return errors.New("Contact with given id not found")
	}
	*contacts = append((*contacts)[:oldContactInd], (*contacts)[oldContactInd+1:]...)
	fmt.Printf("Delete successful for contact with ID:%d \n", id)
	return nil
}

func (contacts *ContactsSlice) ShowContacts() {
	if len(*contacts) == 0 {
		fmt.Println("No Contacts to show.")
	}
	for i, contact := range *contacts {
		fmt.Printf("\nContact #%d\n", i+1)
		fmt.Printf(" ID:%d\n", contact.Id)
		fmt.Printf(" Name:%s\n", contact.Name)
		fmt.Printf(" PhoneNo:%s\n", contact.PhoneNo)
		fmt.Printf(" Address:%s\n", contact.Address)
	}
}

func cleanString(numStr string) string {
	cleanedString := strings.ReplaceAll(numStr, "\r\n", "")
	return cleanedString
}

func ProcesOp(contacts *ContactsSlice, opNo string) error {
	r := bufio.NewReader(os.Stdin)
	switch {
	case opNo == "1":
		return contacts.CreateContact()
	case opNo == "2":
		fmt.Println("Enter the id to update: ")
		idStr, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(cleanString(idStr))
		if err != nil {
			return err
		}
		fmt.Print("Enter the New Name(enter to skip): ")
		nameStr, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Print("Enter the New Phone Number(enter to skip): ")
		phnoStr, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Print("Enter the New Address(enter to skip): ")
		AddStr, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		newContact := Contact{
			Name:    cleanString(nameStr),
			PhoneNo: cleanString(phnoStr),
			Address: cleanString(AddStr),
		}
		return contacts.UpdateContact(id, newContact)
	case opNo == "3":
		fmt.Println("Enter the id to delete: ")
		idStr, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(cleanString(idStr))
		if err != nil {
			return err
		}
		return contacts.DeleteContact(id)
	case opNo == "4":
		contacts.ShowContacts()
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("      				     	CONTACT CRUD											")
	fmt.Println("-----------------------------------------------------------------------------------")

	contacts := ContactsSlice{
		Contact{
			Id:      1,
			Name:    "Temp",
			PhoneNo: "1235",
			Address: "Temp add",
		},
	}
	// indexMap := map[int]Contact{}

	for {
		fmt.Print("\nOptions(0 to exit):")
		fmt.Print("\n 1. Create Contact \n 2. Update Contact \n 3. Delete Contact\n 4. Show All Contacts\n")
		fmt.Print("Select a option:")
		userOptStr, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		opt := cleanString(userOptStr)
		if _, err := strconv.Atoi(opt); err != nil {
			fmt.Println("Invalid option")
			continue
		}

		if opt == "0" {
			break
		}

		err = ProcesOp(&contacts, opt)
		if err != nil {
			fmt.Println(err)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
}
