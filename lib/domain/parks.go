package domain

import (
  "fmt"
  "ken/lib/utils"
  ent"ken/lib/entities"

  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// Persists a park to a DB
func (dom *Domain) CreatePark(p ent.Park) error {
  var ins *sql.Stmt
  ins,err := dom.Dbs.Prepare("INSERT INTO `park_finder`.`parks` (park_id,name,location,description,created_at,updated_at) VALUES(?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error creating park: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating park, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(p.ParkID,p.Name,p.Location,p.Description,p.CreatedAt,p.UpdatedAt)
  if err != nil {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error executing insert review: %s",err))
    utils.Logerror(e)
    return errors.New("Error creating user while executing.")
  }
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error creating park, more than one row affected: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating park.")
  }
  return nil
}

func (dom *Domain) UpdatePark(p ent.Park) error {
  upStmt := "UPDATE `park_finder`.`parks` SET `name` = ? AND location = ? AND description =? WHERE (`park_id` = ?);";
  stmt,err := dom.Dbs.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error preparing to update park: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while updating park")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(&p.ID,&p.ParkID,&p.Name,&p.Location,&p.Description,&p.CreatedAt,&p.UpdatedAt)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error updating park with id: %s. %s",&p.ParkID,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while updating park.")
  }
  return nil
}

func (dom *Domain) ViewPark(pid string) (*ent.Park,error){
  var p ent.Park
  row := dom.Dbs.QueryRow("SELECT * FROM `park_finder`.`parks` WHERE (`park_id` = ?);",pid)
  err := row.Scan(&p.ID,&p.ParkID,&p.Name,&p.Location,&p.Description,&p.CreatedAt,&p.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error viewing park with id: %s  %s",pid,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing park with the id of %s",pid))
  }
  return &p,nil
}

// Returns a list of parks filtered by location or just return all
func  (dom *Domain) ListParksByLocation(location string) ([]ent.Park,error){
  stmt := "SELECT * FROM `park_finder`.`parks` WHERE (`location` = ?) ORDER BY updated_at DESC;"
  rows,err := dom.Dbs.Query(stmt,location)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error listing parks: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing parks.")
  }
  defer rows.Close()
  var parks []ent.Park
  for rows.Next(){
    var p ent.Park
    err = rows.Scan(&p.ID,&p.ParkID,&p.Name,&p.Location,&p.Description,&p.CreatedAt,&p.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning for list parks: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all parks.")
    }
    parks = append(parks,p)
  }
  return parks,nil
}

func  (dom *Domain) ListParks() ([]ent.Park,error){
  stmt := "SELECT * FROM `park_finder`.`parks` ORDER BY updated_at DESC;"
  rows,err := dom.Dbs.Query(stmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error listing parks: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing parks.")
  }
  defer rows.Close()
  var parks []ent.Park
  for rows.Next(){
    var p ent.Park
    err = rows.Scan(&p.ID,&p.ParkID,&p.Name,&p.Location,&p.Description,&p.CreatedAt,&p.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning for list parks: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all parks.")
    }
    parks = append(parks,p)
  }
  return parks,nil
}



// Returns a list f locationparks where the name is likely the one supplied
func  (dom *Domain) SearchParks(park_name string) ([]ent.Park,error){
  return nil,nil
}
