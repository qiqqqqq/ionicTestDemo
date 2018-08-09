package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"database/sql"
	"time"
	"CoalSystem/src/model"
)

func Getonecoal(c *gin.Context) {

	/**
	get a coal
	@Param id
	 */

	var err error
	//db, err := gorm.Open("mysql", "hpx:admin@tcp(192.168.1.4:3306)/mydb?charset=utf8")

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	//db.SingularTable(true)

	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}

	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	row := db.QueryRow("select * from coal where id = ? and active = 1;", id)

	name := ""
	active := 0
	date := ""
	row.Scan(&id, &name, &active, &date)

	//err = db.Table("coal").Where("id = '?'",id).Find(&coal).Error
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"name":   name,
		"active": active,
		"date":   date,
	})

}

func Postcoal(c *gin.Context) {

	var err error
	//db, err := gorm.Open("mysql", "hpx:admin@tcp(192.168.1.4:3306)/mydb?charset=utf8")
	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}

	var req model.Coal
	errbind := c.ShouldBindBodyWith(&req, binding.JSON)

	if errbind != nil {
		fmt.Println("binderr")
	}

	name := req.Name
	visibility := req.Visibility

	db.Exec("insert into coal (name,active,date) values (?,?,?);", name, visibility, time.Now())

	//if err := db.Create(coal).Error; err != nil {
	//	fmt.Println("error")
	//}

	c.JSON(http.StatusOK, gin.H{
		"save success ": req.Name,
	})

}

func Deletecoal(c *gin.Context) {

	var err error

	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		fmt.Println("delete strconv err")
	}

	//db, err := gorm.Open("mysql", "hpx:admin@tcp(192.168.1.4:3306)/mydb?charset=utf8")

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		//panic(err)
		fmt.Println("mysql delete connect error")
	}

	db.Exec("update coal set active =0 where id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}

func Getcoals(c *gin.Context) {

	var err error
	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")
	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}

	rows, err := db.Query("select * from coal where active = 1;")

	type Result struct {
		Id     int
		Name   string
		Active int
		Date   string
	}

	id := 0
	name := ""
	active := 0
	date := ""

	res := make([] Result, 0)
	for rows.Next() {

		rows.Scan(&id, &name, &active, &date)
		result := Result{
			Id:     id,
			Name:   name,
			Active: active,
			Date:   date,
		}

		res = append(res, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func Putonecoal(c *gin.Context) {

	var err error
	//db, err := gorm.Open("mysql", "hpx:admin@tcp(192.168.1.4:3306)/mydb?charset=utf8")
	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}

	sid := c.Param("id")

	id, err := strconv.Atoi(sid)

	var req model.Coal

	errbind := c.ShouldBindBodyWith(&req, binding.JSON)

	if errbind != nil {
		fmt.Println("errbind err")
	}

	name := req.Name
	visibility := req.Visibility

	fmt.Println(name)
	fmt.Println(visibility)
	fmt.Println(id)

	db.Exec("update coal set name = ? ,active = ? where id = ?;", name, visibility, id)

	//if err := db.Model(&req).Where(&Coal{Id: id}).Update(Coal{Name: name, Visibility: visibility}).Error; err != nil {
	//	fmt.Println("update err")
	//}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func Coalanalysisreports(c *gin.Context) {

	var err error
	//db, err := gorm.Open("mysql", "hpx:admin@tcp(192.168.1.4:3306)/mydb?charset=utf8")

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}

	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	//rows, err := db.Table("analysis_category").Select("analysis_category.id,analysis_category.name").
	//	Joins("inner join coal_analysis_report on analysis_category.id=coal_analysis_report.analysis_category_id "+
	//	"and coal_analysis_report.coal_id = ?", id).Rows()

	if err != nil {
		fmt.Println("eeeeerr")
	}
	rows, err := db.Query("select id,analysis_category_id from coal_analysis_report where coal_id = ?;", id)

	if err != nil {
		fmt.Println("dbeeeeerr")
	}

	//defer rows.Close()

	type Result struct {
		Id int
		Cid int
	}

	rid := 0
	cid := 0

	res := make([] Result, 0)

	for rows.Next() {
		rows.Scan(&rid,&cid)

		result := Result{
			Id : rid,
			Cid : cid,
		}
		res = append(res, result)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})

}

func Allattribute(c *gin.Context) {

	var err error

	//db, err := gorm.Open("mysql", "qys:admin@tcp(192.168.1.4:3306)/mydb?charset=utf8")

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	//db, err := gorm.Open("mysql", "root:123456@/mydb?charset=utf8")
	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}

	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	//rows , err := db.Query("select id,name from analysis_attribute where id in ( select analysis_attribute_id from " +
	//	" analysis_type_attribute_list,analysis_type where analysis_category_id = ? and analysis_type.id = analysis_type_attribute_list.analysis_type_id );",id);

	rows, err := db.Query("SELECT analysis_type_attribute_list.id,analysis_category.name,analysis_type.name,analysis_attribute.name " +
		"FROM analysis_category INNER JOIN analysis_type ON analysis_category.id = analysis_type.analysis_category_id INNER JOIN " +
		"analysis_type_attribute_list ON analysis_type.id = analysis_type_attribute_list.analysis_type_id INNER JOIN " +
		"analysis_attribute ON analysis_attribute.id = analysis_type_attribute_list.analysis_attribute_id WHERE analysis_category.id = ? " +
		"AND analysis_type.active=1 AND analysis_category.active=1;", id)


	//rows, err := db.Query("select analysis_attribute.name from analysis_attribute inner join "+
	//	"analysis_type_attribute_list on analysis_type_attribute_list.analysis_attribute_id = analysis_attribute.id "+
	//	"and analysis_type_attribute_list.id in (select analysis_type_attribute_list.id from analysis_type_attribute_list "+
	//	"inner join analysis_type on analysis_type.id = analysis_type_attribute_list.analysis_type_id "+
	//	"and analysis_type.analysis_category_id = ?);", id)

	//rows, err := db.Table("analysis_attribute").Select("analysis_attribute.name").
	//	Joins("inner join analysis_type_attribute_list on analysis_attribute.id = analysis_type_attribute_list.analysis_attribute_id ").
	//	Joins("inner join analysis_type on analysis_type.id=analysis_type_attribute_list.analysis_type_id"+
	//	"and analysis_type.analysis_category_id = ?", id).Rows()

	defer rows.Close()

	type Result struct {
		Attribute_list_id int
		Cname string
		Tname string
		Aname string
		Empty string
	}

	listid := 0
	aname := ""
	cname := ""
	tname := ""

	res := make([] Result, 0)
	for rows.Next() {

		rows.Scan(&listid,&cname,&tname,&aname)
		result := Result{
			Attribute_list_id:listid,
			Cname: cname,
			Tname:tname,
			Aname:aname,
			Empty:"",
		}

		res = append(res, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func Allanalysisreport(c *gin.Context) {

	var err error

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		fmt.Println("mysql connect error")
	}

	rows, err := db.Query("select id,name from analysis_category;")

	if err != nil {
		fmt.Println("errrrrr")
	}

	defer rows.Close()

	type Result struct {
		Id int
		Name string
	}

	id := 0
	name := ""
	res := make([] Result, 0)
	for rows.Next() {
		rows.Scan(&id,&name)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println("____")
		result := Result{
			Id:id,
			Name: name,
		}
		res = append(res, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func Getonereport(c *gin.Context) {

	var err error

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		fmt.Println("mysql connect error")
	}

	//cid := c.Param("id")
	//coalid, err := strconv.Atoi(cid)
	//fmt.Println(coalid)
	//
	//pid := c.Param("rid")
	//reportid, err := strconv.Atoi(pid)
	//
	//rows, err := db.Query("select id,detail from report_content where report_content.coal_analysis_report_id = ?"+
	//	" and report_content.analysis_type_attribute_list_id in "+
	//	"(select  analysis_type_attribute_list.id from analysis_attribute,analysis_type_attribute_list,analysis_type"+
	//	" where analysis_attribute.id=analysis_type_attribute_list.analysis_attribute_id "+
	//	"and analysis_type_attribute_list.analysis_type_id in (select id from analysis_type "+
	//	"where analysis_category_id = ?));", reportid, reportid)
	//
	//if err != nil {
	//	fmt.Println("errrrrr")
	//}
	//
	//defer rows.Close()
	//
	//type Result struct {
	//	Id   int
	//	Val string
	//}
	//
	//id := 0
	//val := ""
	//res := make([] Result, 0)
	//for rows.Next() {
	//	rows.Scan(&id, &val)
	//
	//	result := Result{
	//		Id:   id,
	//		Val: val,
	//	}
	//	res = append(res, result)
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"message": res,
	//})


	c_id := c.Param("id")
	r_id :=c.Param("rid")
	type res struct {
		Id int
		Detail string
	}



	var (
		coal_r     res
		coal_rs    []res
	)
	rows, err := db.Query("select content_id,detail from coal_report where coal_id=? and report_id=? and active=1;",c_id,r_id)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&coal_r.Id, &coal_r.Detail)
		coal_rs = append(coal_rs, coal_r)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": coal_rs,

	})
}

func Deleteonereport(c *gin.Context) {

	var err error

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		fmt.Println("mysql connect error")
	}

	//var req model.Coal_analysis_report
	//errbind := c.ShouldBindBodyWith(&req, binding.JSON)

	//if errbind != nil {
	//	fmt.Println("binderrrrr")
	//}

	cid := c.Param("cid")
	rid := c.Param("rid")
	fmt.Println(cid)

	db.Exec("update coal_analysis_report set active = 0 where id = ?;", rid)

	if err != nil {
		fmt.Println("errrrrr")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func Postonereport(c *gin.Context) {

	var err error

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		fmt.Println("mysql connect error")
	}

	var receiver model.Receiver

	errbind := c.ShouldBindBodyWith(&receiver, binding.JSON)

	cid := c.Param("cid")
	coalid, err := strconv.Atoi(cid)
	if err != nil {
		fmt.Println("errcoalid")
	}

	categoryid := 0
	categoryid = receiver.Categoryid

	//fmt.Println(coalid)
	fmt.Println(categoryid)

	ret, err := db.Exec("insert into coal_analysis_report (coal_id,analysis_category_id,active,date)values (?,?,?,?)",
		coalid, categoryid, 1, time.Now())

	ins_id, _ := ret.LastInsertId();
	fmt.Println("currentId")
	fmt.Println(ins_id)

	if errbind != nil {
		fmt.Println("errbinding")
	}

	//fmt.Println(len(req))

	//tx, _ := db.Begin()
	//k := 0

	for i, a := range receiver.Req {

		fmt.Print(i)

		listid := a.Coal_type_attribute_list_id
		detail := a.Detail
		visibility := a.Visibility


		fmt.Print(" listId:")
		fmt.Print(listid)
		fmt.Print(" detail:")

		fmt.Println(detail)

		db.Exec("insert into report_content (coal_analysis_report_id,analysis_type_attribute_list_id,active,detail)values (?,?,?,?);",
			ins_id, listid, visibility, detail)

		if err != nil {
			fmt.Println("asd")
		}
	}

	//if k == len(receiver.Req) {
	//	tx.Commit();
	//} else {
	//	tx.Rollback();
	//}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

}

func Putonereport(c *gin.Context) {

	var err error

	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")

	if err != nil {
		fmt.Println("mysql connect error")
	}


	//var req [] model.Report
	//
	//errbind := c.ShouldBindBodyWith(&req, binding.JSON)
	//
	//if errbind != nil {
	//	fmt.Println("errrbind")
	//}
	//
	//for _, a := range req {
	//
	//	repid := a.Coal_analysis_report_id
	//	listid := a.Coal_type_attribute_list_id
	//	detail := a.Detail
	//	visi := a.Visibility
	//
	//	fmt.Println(repid)
	//	fmt.Println(listid)
	//	fmt.Println(detail)
	//	fmt.Println(visi)
	//
	//	db.Exec("update report_content set detail = ?  where analysis_type_attribute_list_id = ? "+
	//		"and  coal_analysis_report_id = ?", detail, listid, repid)
	//
	//}


	var req model.Report_content
	errbind := c.ShouldBindBodyWith(&req,binding.JSON)
	if errbind != nil {
		fmt.Println("errrrr")
	}
	id := req.Id
	val := req.Val

	db.Exec("update report_content set detail = ? where id = ?",val,id)




	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func AllReports(c *gin.Context)  {

	var err error
	db, err := sql.Open("mysql", "hpx:admin@tcp(192.168.1.7:3306)/mydb?charset=utf8")
	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}

	rows, err := db.Query("select coal_analysis_report.id, coal.id, coal.name, analysis_category.id, analysis_category.name from analysis_category,coal_analysis_report,coal where coal_analysis_report.coal_id = coal.id and coal_analysis_report.analysis_category_id = analysis_category.id and coal_analysis_report.active = 1;")

	type Result struct {
		Id     int
		Coalid int
		Coalname string
		Cid int
		Cname  string
	}

	id := 0
	coalid := 0
	coalname := ""
	cid := 0
	cname := ""

	res := make([] Result, 0)
	for rows.Next() {

		rows.Scan(&id, &coalid, &coalname, &cid,&cname)
		result := Result{
			Id:     id,
			Coalid:coalid,
			Coalname:coalname,
			Cid:cid,
			Cname:cname,
		}

		res = append(res, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})

}


func Qiyingsheng(c *gin.Context)  {

	c.JSON(http.StatusOK, gin.H{
		"message": "qiyingsheng",
	})

}