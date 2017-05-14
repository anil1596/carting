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
			User:     "anil",
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

	r.GET("/uploaditem", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/major2/web/desi_carting/uploaditem.html")
		c.Data(200, "text/html", res)
	})
	//**********************fetching Javascript files file
	r.GET("/js/:js_file", func(c *gin.Context) {
		//to ser
		jsFile := c.Param("js_file")

		res, err := ioutil.ReadFile("/home/anil/major2/web/js/" + jsFile)
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

		res, err := ioutil.ReadFile("/home/anil/major2/web/css/" + cssFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, "error while fetching file")
		}
		c.Data(200, "text/css", res)

		// c.Data(200, path.Join("applications", "javascript"), res)
	})

	//********************fetching Images
	r.GET("/img/:img_file", func(c *gin.Context) {
		//to ser
		imgFile := c.Param("img_file")
		extension := strings.ToLower(strings.Split(imgFile, ".")[1])

		res, err := ioutil.ReadFile("/home/anil/major2/web/images/" + imgFile)
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
	// r.POST("/registervendor", func(c *gin.Context) {
	// 	var ven vendor

	// 	c.BindJSON(&ven)

	// 	fmt.Println("\n\nRequest Received  for vendor registration: \n\n ")

	// 	tx, _ := db.Begin() // tx => transaction , _ => error and execute

	// 	defer tx.Rollback() // it will be executed after the completion of local function
	// 	fmt.Println(ven.Owner, ven.Name, ven.Email, ven.Mobile, ven.Address, ven.Image, ven.Description, ven.Offer, ven.Password)
	// 	// var track ID
	// 	var num VID
	// 	// insert into users table
	// 	err := tx.QueryRow(`
	//     INSERT INTO vendors (owner, vendorname, email ,mobile ,address  ,imageaddress ,description,offer, password ) values ($1, $2, $3, $4, $5, $6, $7,$8,$9) returning vendorid
	//       `, ven.Owner, ven.Name, ven.Email, ven.Mobile, ven.Address, ven.Image, ven.Description, ven.Offer, ven.Password).Scan(&num.Vendorid)
	// 	fmt.Println(err)
	// 	commit_err := tx.Commit()

	// 	if commit_err != nil {
	// 		tx.Rollback()
	// 		c.JSON(500, "ERR")
	// 		return
	// 	}
	// 	fmt.Println("Vendor registered and his ID:", num)
	// 	c.JSON(200, num)

	// })

	//I**************************tem menu updation
	r.POST("/additems", func(c *gin.Context) {
		var val item

		fmt.Println("\n\nRequest Received for Adding items: \n\n ")
		c.BindJSON(&val)

		// fmt.Printf("%#v", menu)
		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		fmt.Println("vals:", val.Item_no, val.Name, val.Email, val.Mobile, val.Hostel, val.Room,
			val.Itemname, val.Itemtype, val.Sold, val.Price, val.Itemdescription, val.Imageaddress)

		// temp := "anil"
		// _, err := tx.Exec(`INSERT INTO item (name ,email , mobile , hostel , room , item_name , item_type  , sold , price ,
		// item_description, imageaddress  ) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) `, temp, temp, temp, temp, temp, temp, temp, false, temp, temp, temp)

		var item_num int

		err := tx.QueryRow(`INSERT INTO item (name ,email , mobile , hostel , room , item_name , item_type  , sold , price ,
		item_description, imageaddress  ) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning item_no`, val.Name, val.Email, val.Mobile,
			val.Hostel, val.Room, val.Itemname, val.Itemtype, false, val.Price, val.Itemdescription, val.Imageaddress).Scan(&item_num)

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
		c.JSON(200, 1)

	})

	//*************************customer registration
	// r.GET("/registercustomer", func(c *gin.Context) {
	// 	var cus customer
	// 	c.BindJSON(&cus)

	// 	fmt.Println("\n\nRequest Received : \n\n")

	// 	tx, _ := db.Begin() // tx => transaction , _ => error and execute

	// 	defer tx.Rollback() // it will be executed after the completion of local function

	// 	// var track CSID

	// 	// insert into users table
	// 	tx.QueryRow(`
	//     INSERT INTO customers ( customer_id,first_name, last_name, email , mobile , hostel , room ,   bid_amount text NOT NULL,
	//     item_id int NOT NULL references item(item_no) ) values ($1, $2, $3, $4, $5) returning customer_id
	//       `, cus.Name, cus.Email, cus.Mobile, cus.Address, cus.Password).Scan(&track.Customerid)

	// 	commit_err := tx.Commit()

	// 	if commit_err != nil {
	// 		tx.Rollback()
	// 		c.JSON(500, "ERR")
	// 		return
	// 	}
	// 	// fmt.Println("cutomer registered and his ID:", track.Customerid)
	// 	c.JSON(200, track)

	// })

	// //*****************************Serving vendors and their id's
	// r.GET("/getvendors", func(c *gin.Context) {
	// 	// c.BindJSON(&cus)

	// 	fmt.Println("\n\nRequest Received : \n\n")

	// 	rows, err := db.Query(` SELECT  vendorid, vendorname from vendors `)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(500, "error while retreiving vendors data")
	// 	}

	// 	defer rows.Close()

	// 	// var vendors = make(map[string]int)
	// 	ven := make([]VendorsToSend, 0)

	// 	for rows.Next() {
	// 		var t VendorsToSend
	// 		err := rows.Scan(&t.Vendorid, &t.Vendorname)
	// 		ven = append(ven, t)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			c.JSON(500, "error while retreiving vendors data")
	// 		}
	// 	}
	// 	c.JSON(200, ven)
	// 	fmt.Println("Vendors names are sent")
	// })

	// //****************** method to serve request for MENU of particular vendor
	// r.POST("/getvendorsmenu", func(c *gin.Context) {
	// 	var id VID
	// 	c.BindJSON(&id)

	// 	fmt.Println("\n\nRequest for retreiving vendors menu Received : \n\n")

	// 	rows, err := db.Query(` SELECT  item_no, item_name, item_type, item_nature, price, item_description, imageaddress, discount
	// 	                        from itemmenu where vendor_id = $1 `, id.Vendorid)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(500, "error while retreiving vendors menu")
	// 	}

	// 	defer rows.Close()

	// 	// var vendors = make(map[string]int)
	// 	items := make([]item, 0)

	// 	for rows.Next() {
	// 		var t item
	// 		err := rows.Scan(&t.Itemno, &t.Name, &t.IType, &t.Nature, &t.Price, &t.Description, &t.Image, &t.Discount)
	// 		items = append(items, t)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			c.JSON(500, "error while retreiving vendors menu")
	// 		}
	// 	}
	// 	c.JSON(200, items)
	// 	fmt.Println("Vendors Menu  sent")
	// })

	fmt.Println("\n\n\t #####     Foodies server live on :7070     #####")
	r.Run(":7070")
}

// vendor holds the incoming requests for a vendor registration.
type customer struct {
	Customer_id int    `json:"customer_id,omitempty"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Hostel      string `json:"hostel"`
	Room        string `json:"room"`
	Bid_amount  string `json:"bid_amount"`
	Item_id     int    `json:"item_id"`
}

type VID struct {
	Vendorid int `json:"vendorid"`
}

// type customer struct {
// 	Customerid int      `json:"customer_id,omitempty"`
// 	Name       string   `json:"customer_name"`
// 	Email      string   `json:"emailid"`
// 	Mobile     []string `json:"mobile"`
// 	Address    string   `json:"address"`
// 	Password   string   `json:"password"`
// }

// type CSID struct {
// 	Customerid int `json:"customerid,omitempty"`
// }

// type MENU struct {
// 	ITEMS []item `json:"items"`
// }

//Menu updation
type item struct {
	Item_no         int    `json:"item_no,omitempty"`
	Name            string `json:"name"`
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
}

// //Menu updation
// type item struct {
// 	Vendorid    int     `json:"vendor_id"`
// 	Itemno      int     `json:"item_no,omitempty"`
// 	Name        string  `json:"item_name"`
// 	IType       string  `json:"item_type"`
// 	Nature      bool    `json:"item_nature"`
// 	Description string  `json:"item_description"`
// 	Price       string  `json:"price"`
// 	Image       string  `json:"imageaddress,omitempty"`
// 	Discount    float64 `json:"discount,omitempty"`
// }

// type VendorsToSend struct {
// 	Vendorid   int    `json:"vendor_id"`
// 	Vendorname string `json:"vendorname"`
// }
