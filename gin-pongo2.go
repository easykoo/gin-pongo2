package gin_pongo2

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/flosch/pongo2.v3"
)

type Options struct {
	Directory  string
	Extensions []string
}

type Context gin.Context

func C(c *gin.Context) *Context {
	return (*Context)(c)
}

func prepareOptions(opt Options) Options {
	// Defaults
	if len(opt.Directory) == 0 {
		opt.Directory = "templates"
	}
	if len(opt.Extensions) == 0 {
		opt.Extensions = []string{".tmpl", ".html"}
	}

	return opt
}

var tmplMap map[string]*pongo2.Template

func compile(opt Options) map[string]*pongo2.Template {
	tmplMap = make(map[string]*pongo2.Template)

	dirPath := filepath.Dir(opt.Directory)
	//	Log.Println("dir:", dirPath)
	fileInfos, _ := ioutil.ReadDir(dirPath)

	for _, fileInfo := range fileInfos {
		for _, s := range opt.Extensions {
			if isMatched, _ := regexp.MatchString(".*"+s+"$", fileInfo.Name()); isMatched {
				t, err := pongo2.FromFile(path.Join(opt.Directory, fileInfo.Name()))
				if err != nil {
					log.Fatalf("\"%s\": %v", fileInfo.Name(), err)
				}
				//				Log.Println("xxx:", strings.Replace(fileInfo.Name(), s, "", -1))
				tmplMap[strings.Replace(fileInfo.Name(), s, "", -1)] = t
			} else {
				//				Log.Println("yyy:", strings.Replace(fileInfo.Name(), s, "", -1))
			}
		}
	}

	return tmplMap
}

func PrepareTemplates(option Options) {
	compile(prepareOptions(option))
}

func (c *Context) Pongo2(code int, templateName string, dataMap map[string]interface{}) {
	template, exist := tmplMap[templateName]
	if !exist {
		errMsg := "template " + templateName + " not found"
		log.Println(errMsg)
		http.Error(c.Writer, errMsg, http.StatusInternalServerError)
	}
	err := template.ExecuteWriter(dataMap, c.Writer)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

}
