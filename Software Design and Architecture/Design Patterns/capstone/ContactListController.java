package com.example.sharingapp;
import android.content.Context;

import java.util.ArrayList;

public class ContactListController {
    private ContactList contactList;

    public ContactListController(ContactList contactList){
        this.contactList = contactList;
    }

    public void setContacts(ArrayList<Contact> contactList) {
        this.contactList.setContacts(contactList);
    }

    public ArrayList<Contact> getC0ntact() {
        return contactList.getContacts();
    }

    public boolean addContact(Contact contact, Context context){
        AddContactCommand add_contact_command = new AddContactCommand(contactList, contact, context);
        add_contact_command.execute();
        return add_contact_command.isExecuted();
    }

    public boolean deleteContact(Contact contact, Context context) {
        DeleteContactCommand delete_contact_command = new DeleteContactCommand(contactList, contact, context);
        delete_contact_command.execute();
        return delete_contact_command.isExecuted();
    }

    public boolean editContact(Contact contact, Contact updated_contact, Context context){
        EditContactCommand edit_contact_command = new EditContactCommand(contactList, contact, updated_contact, context);
        edit_contact_command.execute();
        return edit_contact_command.isExecuted();
    }

    public Contact getContact(int index) {
        return contactList.getContact(index);
    }

    public int getIndex(Contact contact) {
        return contactList.getIndex(contact);
    }

    public int getSize() {
        return contactList.getSize();
    }

    public void loadContacts(Context context) {
        contactList.loadContacts(context);
    }

//    public ArrayList<Contact> getActiveBorrowers() {
//        return contactList.getActiveBorrowers();
//    }
//
//    public ArrayList<Item> filterItemsByStatus(String status){
//        return item_list.filterItemsByStatus(status);
//    }

    public Contact getContactByUsername(String username) {
        return contactList.getContactByUsername(username);
    }


    public void addObserver(Observer observer) {
        contactList.addObserver(observer);
    }

    public void removeObserver(Observer observer) {
        contactList.removeObserver(observer);
    }

}
