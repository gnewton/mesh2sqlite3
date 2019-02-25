package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strings"

	"github.com/gnewton/gomesh2016"
	_ "github.com/mattn/go-sqlite3"
)

//const DESC_FILE = "/home/gnewton/tmp/mesh2sqlite3/desc2016.xml"
const DESC_FILE = "/home/gnewton/tmp/mesh2sqlite3/desc2019.gz"

var topLevel = []string{
	"Anatomy", "A",
	"Organisms", "B",
	"Diseases", "C",
	"Chemicals and Drugs", "D",
	"Analytical, Diagnostic and Therapeutic Techniques and Equipment", "E",
	"Psychiatry and Psychology", "F",
	"Phenomena and Processes", "G",
	"Disciplines and Occupations", "H",
	"Anthropology, Education, Sociology and Social Phenomena", "I",
	"Technology, Industry, Agriculture", "J",
	"Humanities", "K",
	"Information Science", "L",
	"Named Groups", "M",
	"Health Care", "N",
	"Publication Characteristics", "V",
	"Geographicals", "Z",
}

func main() {
	fmt.Println(len(os.Args), os.Args)
	loadData()

}

type MeshTree struct {
	ID             int64
	DescriptorUI   string `sql:"size:16"`
	DescriptorName string
	Tree           string
	Year           int16
	Depth          int
	T0             *string `sql:"size:1"`
	T1             *string `sql:"size:3"`
	T2             *string `sql:"size:3"`
	T3             *string `sql:"size:3"`
	T4             *string `sql:"size:3"`
	T5             *string `sql:"size:3"`
	T6             *string `sql:"size:3"`
	T7             *string `sql:"size:3"`
	T8             *string `sql:"size:3"`
	T9             *string `sql:"size:3"`
	T10            *string `sql:"size:3"`
	T11            *string `sql:"size:3"`
	T12            *string `sql:"size:3"`
	T13            *string `sql:"size:3"`
}

func loadData() {
	db, err := dbInit("mesh2019_sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}
	var count int64 = 0
	addToplevel(db, &count)

	///////////////////////////
	descChan, dfile, err := gomesh2016.DescriptorChannelFromFile(DESC_FILE)
	defer dfile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	depth := 0
	dname := ""
	ddes := ""
	dtree := ""

	for d := range descChan {
		//fmt.Println(d)
		//fmt.Println("[", d.DescriptorUI, d.DescriptorName, "]")
		if d.DescriptorUI == "D005260" || d.DescriptorUI == "D008297" {
			fmt.Printf("%+v\n", d)
			//continue
		}
		if d.DescriptorUI == "D048531" {
			fmt.Println(d)
		}
		if d.TreeNumberList != nil {
			for _, tree := range d.TreeNumberList.TreeNumber {
				count = count + 1
				s, p := split(tree)
				if d.DescriptorUI == "D048531" {
					fmt.Println("+++", s, p)
				}
				mt := new(MeshTree)
				mt.ID = count
				mt.Year = d.DateCreated.Year.Text
				mt.DescriptorUI = d.DescriptorUI
				mt.DescriptorName = d.DescriptorName
				mt.Tree = tree
				mt.T0 = &s
				l := len(p)
				mt.Depth = l
				switch l {

				case 13:
					mt.T13 = &p[12]
					fmt.Println(tree)
					fallthrough

				case 12:
					mt.T12 = &p[11]
					fmt.Println(tree)
					fallthrough

				case 11:
					mt.T11 = &p[10]
					fallthrough

				case 10:
					mt.T10 = &p[9]
					fallthrough

				case 9:
					mt.T9 = &p[8]
					fallthrough
				case 8:
					mt.T8 = &p[7]
					fallthrough

				case 7:
					mt.T7 = &p[6]
					fallthrough

				case 6:
					mt.T6 = &p[5]
					fallthrough

				case 5:
					mt.T5 = &p[4]
					fallthrough

				case 4:
					mt.T4 = &p[3]
					fallthrough

				case 3:
					mt.T3 = &p[2]
					fallthrough
				case 2:
					mt.T2 = &p[1]
					fallthrough
				case 1:
					mt.T1 = &p[0]

				}
				db.Create(mt)

				if len(p) > depth {
					depth = len(p)
					dname = d.DescriptorName
					ddes = d.DescriptorUI
					dtree = tree
				}

			}
		}
	}
	fmt.Println(depth, dname, ddes, dtree)
	makeIndexes(db)
	db.Close()

}

func addToplevel(db *gorm.DB, count *int64) {

	for i := 0; i < len(topLevel); i = i + 2 {
		*count = *count + 1
		fmt.Println(topLevel[i], topLevel[i+1])
		name, id := topLevel[i], topLevel[i+1]
		mt := new(MeshTree)
		mt.ID = *count
		mt.DescriptorUI = id
		mt.DescriptorName = name
		mt.Tree = id
		mt.T0 = &id
		db.Create(mt)
	}
}

func split(t string) (string, []string) {
	parts := strings.Split(t, ".")
	start := string([]rune(t)[0])
	first := parts[0]
	parts[0] = string([]rune(first)[1:len(first)])
	return start, parts
}
