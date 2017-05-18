package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// Database connectivity variables
var db *pgx.ConnPool
var db_err error

//Initialise connection to the database
func init() {
	db, db_err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Database: "web_portal",
			User:     "postgres",
			Password: "anil205474",
			Port:     5432,
		},
		MaxConnections: 10,
	})

	if db_err != nil {
		fmt.Println("Can't connect to database")
	}
}

func main() {
	r := gin.Default()

	//*************************Hosting client.html page

	r.GET("/available_Items", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/major2/web/carting/products.html")
		c.Data(200, "text/html", res)
	})
	r.GET("/single_Item", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/major2/web/carting/single_Item.html")
		c.Data(200, "text/html", res)
	})
	r.GET("/main", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/major2/web/carting/index.html")
		c.Data(200, "text/html", res)
	})
	r.GET("/register", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/major2/web/carting/register.html")
		c.Data(200, "text/html", res)
	})
	//**********************fetching Javascript files file
	r.GET("/js/:js_file", func(c *gin.Context) {
		//to ser
		jsFile := c.Param("js_file")

		res, err := ioutil.ReadFile("/home/anil/major2/web/carting/js/" + jsFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, "error while fetching file")
		}
		c.Data(200, "application/javascript", res)

		// c.Data(200, path.Join("applications", "javascript"), res)
	})

	//********************fetching CSS files
	r.GET("/css/:css_file", func(c *gin.Context) {
		//to ser
		cssFile := c.Param("css_file")

		res, err := ioutil.ReadFile("/home/anil/major2/web/carting/css/" + cssFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, "error while fetching file")
		}
		c.Data(200, "text/css", res)

		// c.Data(200, path.Join("applications", "javascript"), res)
	})

	//********************fetching Images
	r.GET("/images/:img_file", func(c *gin.Context) {
		//to ser
		imgFile := c.Param("img_file")
		extension := strings.ToLower(strings.Split(imgFile, ".")[1])

		res, err := ioutil.ReadFile("/home/anil/major2/web/carting/images/" + imgFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, "error while fetching Image")
		}

		if extension == "jpg" {
			c.Data(200, "image/jpg", res)
		} else if extension == "png" {
			c.Data(200, "image/png", res)
		} else if extension == "jpeg" {
			c.Data(200, "image/png", res)
		}

		// c.Data(200, path.Join("applications", "javascript"), res)
	})

	// //********************Registering vendors
	r.POST("/delitem", func(c *gin.Context) {
		var temp temp_item

		c.BindJSON(&temp)

		// fmt.Println(" %#v   ", temp)

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		// fmt.Println(temp.Item_no, temp.Password)
		// var track ID

		// insert into users table
		rows, err := db.Query(` SELECT  item_no, password
		                        from item where item_no = $1 `, temp.Item_no)
		if err != nil {

			fmt.Println("error while item from database", err)
			c.JSON(500, 0)
		}

		defer rows.Close()

		// fmt.Println("after retreivingvalues from database")

		for rows.Next() {
			var t temp_item
			err := rows.Scan(&t.Item_no, &t.Password)
			// fmt.Println(t.Item_no, t.Password)

			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error while retreiving vendors menu")
			}
			if strings.Compare(temp.Password, t.Password) == 0 {
				fmt.Println("exact before deletion", t.Item_no)
				tx.Exec(` Delete   from item where item_no = $1 `, t.Item_no)

				commit_err := tx.Commit()

				if commit_err != nil {
					tx.Rollback()
					c.JSON(500, "ERRor while deletion")
					return
				} else {
					fmt.Println("item deleted successfully")
					c.JSON(222, 1)
				}

			}
		}
		//item deleted successfully
		c.JSON(200, 0)
	})

	//I**************************tem menu updation
	r.POST("/additems", func(c *gin.Context) {
		var val item

		fmt.Println("\n\nRequest Received for Adding items: \n\n ")
		c.BindJSON(&val)
		fmt.Printf("%#v", val)
		fmt.Println("vals:", val.Item_no, val.Name, val.Roll_no, val.Email, val.Mobile, val.Hostel, val.Room,
			val.Itemname, val.Itemtype, val.Sold, val.Price, val.Itemdescription, val.Imageaddress)

		rows, err := db.Query(` SELECT roll_no, password ,block from item where roll_no = $1 `, val.Roll_no)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, "error while retreiving vendors menu")
		}

		defer rows.Close()

		for rows.Next() {
			var t student
			err := rows.Scan(&t.Roll_no, &t.Password)
			// fmt.Println(t.Item_no, t.Password)

			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error while retreiving vendors menu")
			}

			if val.Roll_no == t.Roll_no && strings.Compare(val.Password, t.Password) == 0 && t.Block == false {
				fmt.Println("student verified", t.Roll_no)

				tx, _ := db.Begin() // tx => transaction , _ => error and execute

				defer tx.Rollback() // it will be executed after the completion of local function

				var item_num int

				err := tx.QueryRow(`INSERT INTO item (roll_no,name ,email , mobile , hostel , room , item_name , item_type  , sold , price ,
		item_description, imageaddress,password ) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) returning item_no`, val.Roll_no, val.Name, val.Email, val.Mobile,
					val.Hostel, val.Room, val.Itemname, val.Itemtype, false, val.Price, val.Itemdescription, val.Imageaddress, val.Password).Scan(&item_num)

				if err != nil {
					// c.JSON(500, "error")
					fmt.Println("error while adding \n\n", err)
				}

				commit_err := tx.Commit()

				if commit_err != nil {
					tx.Rollback()
					fmt.Println("error while committing \n\n")
					fmt.Println(commit_err)
					c.JSON(500, "ERR")
					return
				}
				c.JSON(222, 1)
			}
		}
		//item deleted successfully
		c.JSON(500, 0)

		////

	})

	//****************** method for retreiving all available_Items
	r.GET("/getitems", func(c *gin.Context) {

		fmt.Println("\n\nRequest for retreiving items Received : \n\n")

		rows, err := db.Query(` SELECT item_no, name ,email , mobile , hostel , room , item_name , item_type  ,  price ,
		item_description from item `)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, "error while retreiving vendors menu")
		}

		defer rows.Close()

		// var vendors = make(map[string]int)
		items := make([]item, 0)

		for rows.Next() {
			var t item
			err := rows.Scan(&t.Item_no, &t.Name, &t.Email, &t.Mobile, &t.Hostel, &t.Room, &t.Itemname, &t.Itemtype, &t.Price, &t.Itemdescription)
			items = append(items, t)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error while retreiving items")
			}
		}
		fmt.Println("items  sent")
		c.JSON(200, items)

	})

	//****************** method for retreiving selected Items
	r.GET("/getSelectedItems", func(c *gin.Context) {

		fmt.Println("\n\nRequest for retreiving selected items Received : \n\n")
		var temp itemType
		c.BindJSON(&temp)
		rows, err := db.Query(` SELECT item_no, name ,email , mobile , hostel , room , item_name , item_type  ,  price ,
		item_description from item where item_type = $1 `, temp.Type)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, "error while retreiving vendors menu")
		}

		defer rows.Close()

		// var vendors = make(map[string]int)
		items := make([]item, 0)

		for rows.Next() {
			var t item
			err := rows.Scan(&t.Item_no, &t.Name, &t.Email, &t.Mobile, &t.Hostel, &t.Room, &t.Itemname, &t.Itemtype, &t.Price, &t.Itemdescription)
			items = append(items, t)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error while retreiving items")
			}
		}
		fmt.Println("items  sent")
		c.JSON(200, items)

	})

	// ***********************code to verify  admin *********************
	r.POST("/verifyAdmin", func(c *gin.Context) {

		fmt.Println("\n\nRequest to verifyAdmin Received : \n\n")
		var temp admin
		c.BindJSON(&temp)
		rows, err := db.Query(` SELECT employee_no, password from item where item_type = $1 `, temp.Employee_no)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, "error in verifying admin")
		}

		defer rows.Close()

		for rows.Next() {
			var t admin
			err := rows.Scan(&t.Employee_no, &t.Password)
			// fmt.Println(t.Item_no, t.Password)

			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error in verifying admin")
			}
			if t.Employee_no == temp.Employee_no && strings.Compare(temp.Password, t.Password) == 0 {
				fmt.Println("Admin verified", t.Employee_no)
				c.JSON(222, 0)
			}
		}
		//item deleted successfully
		c.JSON(500, 0)
	})
	//***************** function to add student
	r.GET("/addStudent", func(c *gin.Context) {

		fmt.Println("\n\nRequest to add student Received : \n\n")
		var val student
		var rol int
		c.BindJSON(&val)
		fmt.Println("%#v", val)
		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		err := tx.QueryRow(`INSERT INTO student (roll_no,password,block ) values ($1,$2,$3) returning roll_no`,
			val.Roll_no, val.Password, val.Block).Scan(&rol)

		if err != nil {
			// c.JSON(500, "error")
			fmt.Println("error while adding student\n\n", err)
		}

		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			fmt.Println("error while committing \n\n")
			fmt.Println(commit_err)
			c.JSON(500, "ERR")
			return
		}
		c.JSON(222, 1)

	})
	r.POST("/alterStudent", func(c *gin.Context) {

		fmt.Println("\n\nRequest to alter student Received : \n\n")
		var val student

		c.BindJSON(&val)
		fmt.Println("%#v", val)
		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function
		rows, err := db.Query(` SELECT  roll_no,password,block
		                        from students where roll_no = $1 `, val.Roll_no)
		if err != nil {

			fmt.Println("error while item from database", err)
			c.JSON(500, 0)
		}

		defer rows.Close()

		// fmt.Println("after retreivingvalues from database")

		for rows.Next() {
			var t student
			err := rows.Scan(&t.Roll_no, &t.Block)
			// fmt.Println(t.Item_no, t.Password)

			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error while retreiving vendors menu")
			}
			if t.Block == false {
				_, err := tx.Exec(`UPDATE students SET block = true where roll_no= $1`, t.Roll_no)
				if err != nil {
					fmt.Println(err)
					c.JSON(500, "error while updating student")
				}
			} else {
				_, err := tx.Exec(`UPDATE students SET block = false where roll_no= $1`, t.Roll_no)
				if err != nil {
					fmt.Println(err)
					c.JSON(500, "error while updating student")
				}
			}
			c.JSON(222, 0)
		}
	})

	fmt.Println("\n\n\t #####     NITH web_portal server live on :7070     #####")
	r.Run(":7070")
}

//Menu updation
type item struct {
	Item_no         int    `json:"item_no,omitempty"`
	Name            string `json:"name"`
	Roll_no         int    `json:"roll_no"`
	Email           string `json:"email"`
	Mobile          string `json:"mobile"`
	Hostel          string `json:"hostel"`
	Room            string `json:"room"`
	Itemname        string `json:"itemname"`
	Itemtype        string `json:"itemtype"`
	Sold            string `json:"sold"`
	Price           string `json:"price"`
	Itemdescription string `json:"itemdescription"`
	Imageaddress    string `json:"imageaddress,omitempty"`
	Password        string `json:"password"`
}

type temp_item struct {
	Item_no  int    `json:"item_no"`
	Password string `json:"password"`
}

type itemType struct {
	Type string `json:"type"`
}

type admin struct {
	Employee_no int    `json:"employee_no"`
	Password    string `json:"password"`
}

type student struct {
	Roll_no  int    `json:"roll_no"`
	Password string `json:"password"`
	Block    bool   `json:"block"`
}
