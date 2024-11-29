package jobaction

type CreateTeamAction string

const (
	CREATE_TEAM_VALIDATE_TEAM                                     CreateTeamAction = "VALIDATE_TEAM"
	CREATE_TEAM_VALIDATE_ORGANIZATION_AND_ORGANIZATION_USER_EXIST CreateTeamAction = "VALIDATE_ORGANIZATION_AND_ORGANIZATION_USER_EXIST"
	CREATE_TEAM_CREATE_TEAM_AND_ADD_INITIAL_TEAM_USER             CreateTeamAction = "CREATE_TEAM_AND_ADD_INITIAL_TEAM_USER"
)

type AddUsersTeamAction string

const (
	ADD_USERS_TEAM_VALIDATE_TEAM                                     AddUsersTeamAction = "VALIDATE_TEAM"
	ADD_USERS_TEAM_VALIDATE_ORGANIZATION_AND_ORGANIZATION_USER_EXIST AddUsersTeamAction = "VALIDATE_ORGANIZATION_AND_ORGANIZATION_USER_EXIST"
	ADD_USERS_TEAM_ADD_TEAM_USER                                     AddUsersTeamAction = "ADD_TEAM_USER"
)
