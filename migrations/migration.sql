CREATE TABLE Users (
    ID SERIAL,
    Username varchar(255) NOT NULL,
    FirstName varchar(255),
    LastName varchar(255),
    Email varchar(255),
    TwilioSID varchar(255),
    PRIMARY KEY (ID)
);
