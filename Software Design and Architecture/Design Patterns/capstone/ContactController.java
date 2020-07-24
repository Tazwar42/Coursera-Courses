package com.example.sharingapp;
import android.graphics.Bitmap;

public class ContactController {
    private Contact contact;

    public ContactController(Contact contact){
        this.contact = contact;
    }

    public String getId(){
        return contact.getId();
    }

    public void setId() {
        contact.setId();
    }

    public void setEmail(String email) {
        contact.setEmail(email);
    }

    public void getEmail(){
        contact.getEmail();
    }

    public void setUsername(String username){
        contact.setUsername(username);
    }

    public void getUsername(){
        contact.getUsername();
    }


    public Contact getContact() { return this.contact; }

    public void addObserver(Observer observer) {
        contact.addObserver(observer);
    }

    public void removeObserver(Observer observer) {
        contact.removeObserver(observer);
    }

}
