package subjs

const (
	// SubjectElasticUserFetch topic for fetching user
	SubjectElasticUserFetch = "elastic.v1.user.fetch"
	// SubjectElasticUserPut topic for create / update user
	SubjectElasticUserPut = "elastic.v1.user.put"
	// SubjectElasticPasswordReset topic for reset user's password
	SubjectElasticPasswordReset = "elastic.v1.user.password_reset"

	// SubjectAccessGrant topic for granting access
	SubjectAccessGrant = "access.v1.command.grant"
)
