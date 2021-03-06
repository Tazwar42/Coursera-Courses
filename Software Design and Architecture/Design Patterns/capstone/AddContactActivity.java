package com.example.sharingapp;

import android.content.Context;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;

/**
 * Add a new contact
 */
public class AddContactActivity extends AppCompatActivity {

    private ContactList contact_list = new ContactList();
    private ContactListController contactListController = new ContactListController(contact_list);
    private Context context;

    private EditText username;
    private EditText email;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_add_contact);

        username = (EditText) findViewById(R.id.username);
        email = (EditText) findViewById(R.id.email);

        context = getApplicationContext();
        //contact_list.loadContacts(context);
        contactListController.loadContacts(context);
    }

    public void saveContact(View view) {

        String username_str = username.getText().toString();
        String email_str = email.getText().toString();

        if (username_str.equals("")) {
            username.setError("Empty field!");
            return;
        }

        if (email_str.equals("")) {
            email.setError("Empty field!");
            return;
        }

        if (!email_str.contains("@")){
            email.setError("Must be an email address!");
            return;
        }

        for (Contact c : contact_list.getContacts()) {
            if (c.getUsername().equals(username_str)) {
                username.setError("Username already taken!");
                return;
            }
        }

        Contact contact = new Contact(username_str, email_str, null);

        // Add Contact
        //AddContactCommand add_contact_command = new AddContactCommand(contact_list, contact, context);
        //add_contact_command.execute();
        ContactController contactController = new ContactController(contact);

        boolean success = contactListController.addContact(contact,context);
        if (!success) {
            return;
        }

        // End AddContactActivity
        finish();
    }
}
