package com.example.sharingapp;

import android.content.Context;
import android.content.Intent;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;

/**
 * Editing a pre-existing contact consists of deleting the old contact and adding a new contact with the old
 * contact's id.
 * Note: You will not be able contacts which are "active" borrowers
 */
public class EditContactActivity extends AppCompatActivity {

    private ContactList contact_list = new ContactList();
    private ContactListController  contactListController = new ContactListController(contact_list);
    private ContactController contactController;
    private Contact contact;
    private EditText email;
    private EditText username;
    private Context context;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_edit_contact);

        context = getApplicationContext();
        contact_list.loadContacts(context);

        Intent intent = getIntent();
        int pos = intent.getIntExtra("position", 0);

        contact = contact_list.getContact(pos);

        username = (EditText) findViewById(R.id.username);
        email = (EditText) findViewById(R.id.email);

        username.setText(contact.getUsername());
        email.setText(contact.getEmail());

        contactListController.addObserver(this);
        contactListController.loadContacts(context);


    }

    public void saveContact(View view) {

        String email_str = email.getText().toString();

        if (email_str.equals("")) {
            email.setError("Empty field!");
            return;
        }

        if (!email_str.contains("@")){
            email.setError("Must be an email address!");
            return;
        }

        String username_str = username.getText().toString();

        // Check that username is unique AND username is changed (Note: if username was not changed
        // then this should be fine, because it was already unique.)
        if (!contact_list.isUsernameAvailable(username_str) && !(contact.getUsername().equals(username_str))) {
            username.setError("Username already taken!");
            return;
        }

        String id = contact.getId(); // Reuse the contact id
        Contact updated_contact = new Contact(username_str, email_str, id);

        // Edit contact
//        EditContactCommand edit_contact_command = new EditContactCommand(contact_list, contact, updated_contact, context);
//        edit_contact_command.execute();

        ContactController updated_contact_controller = new ContactController(updated_contact);

        boolean success = contactListController.editContact(contact,updated_contact,context);
        if (!success){
            return;
        }

        contactListController.removeObserver(this);

        // End EditContactActivity
        finish();
    }

    public void deleteContact(View view) {

        // Delete contact
        //DeleteContactCommand delete_contact_command = new DeleteContactCommand(contact_list, contact, context);
        //delete_contact_command.execute();

        boolean success = contactListController.deleteContact(contact,context);
        if (!success){
            return;
        }

        contactListController.removeObserver(this);
        // End EditContactActivity
        finish();
    }
}
