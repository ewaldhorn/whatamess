# Simple Iced Application

Based on <https://blog.logrocket.com/iced-rs-tutorial-rust-frontend-web-app/>

# Structure of an Iced app
An Iced.rs application, similar to ELM, consists of four central concepts:

State
Messages
Update
View
The State is the state of your application. In our case, for example, this is the data we fetch and display from JSONPlaceholder.

Messages are used to trigger flow inside the application. they can be user interaction, timed events, or any other event, 
which might change something within the application.

The Update logic is used to react to these Messages. For example, in our application, there might be a Message for 
navigating to the Detail page. In our Update logic for this message, we will set the route and fetch the data, so we can 
update the application state from List to Detail.

Finally, the View logic describes how to render a certain piece of the application. It displays the State and might 
produce Messages on user interaction.
