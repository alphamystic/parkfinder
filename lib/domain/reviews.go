package domain



import (
  "fmt"
  "ken/lib/utils"
  ent"ken/lib/entities"

  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// 	id 	review_id 	park_id 	userid 	username 	comment 	created_at 	updated_at
func (dom *Domain) CreateReview(r ent.Review) error {
  var ins *sql.Stmt
  ins,err := dom.Dbs.Prepare("INSERT INTO `park_finder`.`reviews` (review_id,park_id,userid,username,comment,created_at,updated_at) VALUES(?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error preparing to create review: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating review, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(r.ReviewID,r.ParkID,r.UserID,r.UserName,r.Comment,r.CreatedAt,r.UpdatedAt)
  if err != nil {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error executing insert review: %s",err))
    utils.Logerror(e)
    return errors.New("Error creating review while executing.")
  }
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error more than one row affected in creating review: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating review!!!.")
  }
  return nil
}

func (dom *Domain) ListReviews(park_id string) ([]ent.Review,error){
  stmt := "SELECT * FROM `park_finder`.`reviews` WHERE (`park_id` = ? ) ORDER BY updated_at DESC;"
  rows,err := dom.Dbs.Query(stmt,park_id)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error listing reviews: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing reviews.")
  }
  defer rows.Close()
  var ents []ent.Review
  for rows.Next(){
    var r ent.Review
    err = rows.Scan(&r.ID,&r.ReviewID,&r.ParkID,&r.UserID,&r.UserName,&r.Comment,&r.CreatedAt,&r.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESAHN: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all reviews.")
    }
    ents = append(ents,r)
  }
  return ents,nil
}

func (dom *Domain) EditReview(review ent.Review) error {
  upStmt := "UPDATE `park_finder`.`news` SET `handled` = ? WHERE (`newsid` = ?);";
  stmt,err := dom.Dbs.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EMNH: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while marking news as handled.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(true,review.ReviewID)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error editinng review id: %s  %s",review.ReviewID,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while editing review.")
  }
  return nil
}

func (dom *Domain) GetReview(rid string) (*ent.Review,error) {
  var r ent.Review
  row := dom.Dbs.QueryRow("SELECT * FROM `park_finder`.`review` WHERE (`review_id` = ?);",rid)
  err := row.Scan(&r.ID,&r.ReviewID,&r.ParkID,&r.UserID,&r.UserName,&r.Comment,&r.CreatedAt,&r.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error viewing riview id: %s  %s",rid,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing review with the id of %s",rid))
  }
  return &r,nil
}
