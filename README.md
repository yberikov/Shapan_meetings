This Go package includes functionalities to match users based on certain criteria and schedule events in Google Calendar.

Prerequisites

Before running the application, ensure you have the following prerequisites installed:

Go 1.16 or higher
Git
Installation


To install the package, clone the repository to your local machine:

```
git clone https://github.com/your-username/match-scheduler.git
```

Usage


Set Up Google Calendar API Credentials:

Obtain credentials for the Google Calendar API from the Google Cloud Console and save them in a file named credentials.json in the project root directory.


Build and Run:

Navigate to the project directory.

Build and run the application using the following command:


```
go run .
```
Endpoints:

The application exposes the following endpoint:

/searchSpeaking: This endpoint searches for potential matches among users. It accepts POST requests with user data and returns matching pairs.
