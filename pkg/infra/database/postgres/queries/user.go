package queries

type userSqlManager struct{}
type userSelectManager struct{}

func User() *userSqlManager {
	return &userSqlManager{}
}

func (userSqlManager) Select() *userSelectManager {
	return &userSelectManager{}
}

func (userSelectManager) ById() string {
	return `
    SELECT 
      u.id as id, 
      u.name as name, 
      u.email as email, 
      u.phone as phone, 
      u.password as password, 
      u.cpf as cpf, 
      u.role as role 
    FROM users u WHERE u.id = $1;
  `
}

func (userSelectManager) All() string {
	return `
    SELECT 
      u.id as id, 
      u.name as name, 
      u.email as email, 
      u.phone as phone, 
      u.password as password, 
      u.cpf as cpf, 
      u.role as role 
    FROM users u;
  `
}

func (userSelectManager) ByEmail() string {
	return `
    SELECT 
      u.id as id, 
      u.name as name, 
      u.email as email, 
      u.phone as phone, 
      u.password as password, 
      u.cpf as cpf, 
      u.role as role 
    FROM users u WHERE u.email = $1;
  `
}

func (userSqlManager) Update() string {
	return `
  UPDATE users SET 
    name=$1, 
    email=$2, 
    phone=$3, 
    cpf=$4, 
    password=$5, 
    role=$6,
    updated_at=CURRENT_TIMESTAMP
  WHERE id=$7;
  `
}

func (userSqlManager) Create() string {
	return `
    INSERT INTO users 
      (id, name, email, cpf, phone, password, role) 
    VALUES 
      ($1, $2, $3, $4, $5, $6, $7);
  `
}

func (userSqlManager) Delete() string {
	return `
  DELETE FROM users WHERE id = $1;
  `
}
