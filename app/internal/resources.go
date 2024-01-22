package internal

const (
	MimeTypeContentType     = "Content-Type"
	MimeTypeApplicationZip  = "application/zip"
	MimeTypeApplicationJSON = "application/json"
)

const (
	// SignIn Auth sign_in handler
	SignIn = "/api/auth/sign_in"
	SignUp = "/api/auth/sign_up"

	//RolesURI
	RolesURL = "/api/roles"
	RoleURL  = "/api/roles/:id"

	// UsersURL URI user path
	UsersURL = "/api/users"
	UserURL  = "/api/users/:uuid"

	// CategoriesURL Category  URI user path
	CategoriesURL = "/api/categories"
	CategoryURL   = "/api/categories/:id"
)

const (
	UserTable = "public.user"
)
