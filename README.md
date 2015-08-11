# gin-pongo2

## Getting Started

how to use:

#+BEGIN_SRC go

import p "github.com/easykoo/gin-pongo2"

	p.PrepareTemplates(p.Options{
		Directory:  "public/views/",
		Extensions: []string{".html"},
	})
	
	mData := make(map[string]interface{})
	p.C(c).Pongo2(http.StatusOK, "share", mData)

#+END_SRC

##License

This code is under an Apache v2 License.


## Author

* [Steven](https://github.com/easykoo)