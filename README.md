# gin-pongo2

## Usage

~~~ go

import p "github.com/easykoo/gin-pongo2"

	p.PrepareTemplates(p.Options{
		Directory:  "public/views/",
		Extensions: []string{".html"},
	})
	
	mData := make(map[string]interface{})
	p.C(c).Pongo2(http.StatusOK, "share", mData)

~~~

##License

This code is under an Apache v2 License.


## Author

* [Steven](https://github.com/easykoo)