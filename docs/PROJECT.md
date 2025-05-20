# Natural stupidity Project

The People project aims to create an internal platform for Doutor Finanças that allows to create a calendar, which shows who is on vacations. It should allow to create users by hierarchy, where the rank is:

## Base Features
This project aims to deliver the following high level features:

### User Profiles and Data
- Each user should be able to modify and view it's own profile
- All users should be able to view basic profile information from one another
- The user visible profile should have a picture, basic information such as Name, Phone, Email, Department, Position
- The user should have more information for Need To Know, such as:
    - Personal Information (optional):
        - Phone
        - Email
        - Address
        - DOB
    - Funny Information (optional; Optionally visible):
        - Have animals? Which type&name&age
        - Have kids? How many?
        - Personal Interests or Hobbies
- An Administrator may moderate or modify this information
### User Types
- There are several user types
    - Regular user
    - Power user (higher permissions)
    - Administrator (all permissions)
    - Technical Administrator (cannot see user data, but can manage all internal aspects of the application, such as configurations, etc)
- User types should be managed in a general user management list, visible for administrators, for instance, in the user profiles
### User Groups and Roles
- There may be several user roles. They should be createable in the application. This roles can be of the following types:
    - Specialist (VI)
    - Team Leader (V)
    - Head (IV)
    - Sub Director (III)
    - Director (II)
    - Board (I)
    - Other (X)
- The roman numerals represent the hierarquic scale (less is higher). The X represents a role off the hierachy, used for Jobs other than Roles. This aims to be a simplification
- We should be able to define Teams of users, and who is their team leader (admin functionality)
- We should be able to define collections of teams, named directions, and who is their director and/or sub director (admin functionality)
- We should be be able to define a Board, and which directions report to which member of the board, and also, which member of the board reports to which other member of the board. (admin functionality)
- There may be a member that does not report to anyone. Usually it should be either the Chairman of the Board or the CEO (admin functionality)
- Users should authenticate with password. We should allow user creation.
- The default 1st user is Admin and the password is 123456. It will eventually be changed.
### Calendar, Days off, Sick days
Each user should be able to have vacation days
- Administrators can configure a default number of days for all collaborators
- Administrators can change the number of days a member may take. This change only affects a certain year (by default the current one)
- Each member can take vacation days. It can only be:
    - Week days
    - Non holidays (depending on the country calendar)
- The calendar should mark the birthday of the user.
- The user should be able to view it's time off
- The team leader should be able to view their teams time off (based on the user teams feature)
- The director should be able to view it's team leads vacations days and their teams vacations
- The C-levels the same. It should work with nested reporting logic, a user views the vacations of those who report to them, and those who report to this one.
- Indirect reporting should also appear, but it a "lighter" way
- The filters to view teams, directions, etc. Should be by each teams (or all), each direction (or all), etc.


## Technical Details and decisions
- We are using golang for backend
- We are using bootstrap + htmx for frontend
- We have figmas that can be consulted for implementation examples, but not yet for this application
- We are using mysql for this application database
- We already have a set of API implemented. We may provide the details for it's usage in a near future
- This should be done one step at a time, tested, verified and afterwards we can do the next one.

# Database configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=natural_stupidity

# Application configuration
APP_PORT=8080
GIN_MODE=debug
LOG_LEVEL=debug
JWT_SECRET=your_jwt_secret

# Air live reload
AIR_WD=/app
AIR_PORT=8080
