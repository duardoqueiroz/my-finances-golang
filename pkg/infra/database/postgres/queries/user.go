package queries

type userSqlManager struct{}
type userSelectManager struct{}

func User() *userSqlManager {
	return &userSqlManager{}
}

func (*userSqlManager) Select() *userSelectManager {
	return &userSelectManager{}
}

func (*userSelectManager) ById() string {
	return `
    SELECT 
      u.id as id, 
      u.name as name, 
      u.email as email, 
      u.phone as phone, 
      u.cpf as cpf, 
      u.role as role 
    FROM users u WHERE u.id = $1;
  `
}
